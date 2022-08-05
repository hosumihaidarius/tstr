package auth

import (
	"context"
	"crypto/sha512"
	"encoding/hex"
	"errors"
	"strings"

	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	commonv1 "github.com/nanzhong/tstr/api/common/v1"
	"github.com/nanzhong/tstr/db"
	"github.com/nanzhong/tstr/grpc/types"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

const mdAuthKey = "authorization"

// TODO We can probably be a bit smarter/less verbose here and instead of direct
// string matches on the full method, build up the set of allowable tokens given
// a full method and a list of regexes (or something like that)
var scopeAuthorizations = map[string][]commonv1.AccessToken_Scope{
	"/tstr.control.v1.ControlService/RegisterTest":     {commonv1.AccessToken_CONTROL_RW},
	"/tstr.control.v1.ControlService/UpdateTest":       {commonv1.AccessToken_CONTROL_RW},
	"/tstr.control.v1.ControlService/GetTest":          {commonv1.AccessToken_CONTROL_RW, commonv1.AccessToken_CONTROL_R},
	"/tstr.control.v1.ControlService/ListTests":        {commonv1.AccessToken_CONTROL_RW, commonv1.AccessToken_CONTROL_R},
	"/tstr.control.v1.ControlService/ArchiveTest":      {commonv1.AccessToken_CONTROL_RW},
	"/tstr.control.v1.ControlService/DefineTestSuite":  {commonv1.AccessToken_CONTROL_RW},
	"/tstr.control.v1.ControlService/UpdateSuite":      {commonv1.AccessToken_CONTROL_RW},
	"/tstr.control.v1.ControlService/GetTestSuite":     {commonv1.AccessToken_CONTROL_RW, commonv1.AccessToken_CONTROL_R},
	"/tstr.control.v1.ControlService/ListTestSuites":   {commonv1.AccessToken_CONTROL_RW, commonv1.AccessToken_CONTROL_R},
	"/tstr.control.v1.ControlService/ArchiveTestSuite": {commonv1.AccessToken_CONTROL_RW},
	"/tstr.control.v1.ControlService/GetRun":           {commonv1.AccessToken_CONTROL_RW, commonv1.AccessToken_CONTROL_R},
	"/tstr.control.v1.ControlService/ListRuns":         {commonv1.AccessToken_CONTROL_RW, commonv1.AccessToken_CONTROL_R},
	"/tstr.control.v1.ControlService/ScheduleRun":      {commonv1.AccessToken_CONTROL_RW},

	"/tstr.admin.v1.AdminService/IssueAccessToken":  {commonv1.AccessToken_ADMIN},
	"/tstr.admin.v1.AdminService/GetAccessToken":    {commonv1.AccessToken_ADMIN},
	"/tstr.admin.v1.AdminService/ListAccessTokens":  {commonv1.AccessToken_ADMIN},
	"/tstr.admin.v1.AdminService/RevokeAccessToken": {commonv1.AccessToken_ADMIN},

	"/tstr.runner.v1.RunnerService/RegisterRunner": {commonv1.AccessToken_RUNNER},
	"/tstr.runner.v1.RunnerService/NextRun":        {commonv1.AccessToken_RUNNER},
	"/tstr.runner.v1.RunnerService/SubmitRun":      {commonv1.AccessToken_RUNNER},

	"/tstr.data.v1.DataService/GetTest":         {commonv1.AccessToken_DATA},
	"/tstr.data.v1.DataService/QueryTests":      {commonv1.AccessToken_DATA},
	"/tstr.data.v1.DataService/GetTestSuite":    {commonv1.AccessToken_DATA},
	"/tstr.data.v1.DataService/QueryTestSuites": {commonv1.AccessToken_DATA},
	"/tstr.data.v1.DataService/GetRun":          {commonv1.AccessToken_DATA},
	"/tstr.data.v1.DataService/QueryRuns":       {commonv1.AccessToken_DATA},
	"/tstr.data.v1.DataService/SummarizeRuns":   {commonv1.AccessToken_DATA},
	"/tstr.data.v1.DataService/GetRunner":       {commonv1.AccessToken_DATA},
	"/tstr.data.v1.DataService/QueryRunners":    {commonv1.AccessToken_DATA},
}

// TODO We shouldn't reach out to the db each time to auth, especially when
// these results are easy to cache/invalidate.

func UnaryServerInterceptor(pgxPool *pgxpool.Pool) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		md, exists := metadata.FromIncomingContext(ctx)
		if !exists {
			return nil, status.Error(codes.Unauthenticated, "failed to authenticate request: missing access token")
		}

		_, tokenHash, err := tokenFromMD(md)
		if err != nil {
			return nil, status.Error(codes.Unauthenticated, "failed to authenticate request: invalid access token")
		}

		validScopes := scopeAuthorizations[info.FullMethod]
		auth, err := db.New().AuthAccessToken(ctx, pgxPool, tokenHash)
		if err != nil {
			if errors.Is(err, pgx.ErrNoRows) {
				return nil, status.Error(codes.Unauthenticated, "failed to authenticate request: invalid access token")
			}

			log.Error().Err(err).Msg("failed to validate access token")
			return nil, status.Error(codes.Internal, "failed to authenticate request")
		}

		for _, vs := range types.FromAccessTokenScopes(validScopes) {
			for _, s := range auth.Scopes {
				if string(vs) == s {
					return handler(ctx, req)
				}
			}
		}
		return nil, status.Error(codes.PermissionDenied, "failed to authenticate request: invalid access token scopes")
	}
}

func StreamServerInterceptor(pgxPool *pgxpool.Pool) grpc.StreamServerInterceptor {
	return func(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
		md, exists := metadata.FromIncomingContext(ss.Context())
		if !exists {
			return status.Error(codes.Unauthenticated, "failed to authenticate request: missing access token")
		}

		_, tokenHash, err := tokenFromMD(md)
		if err != nil {
			return status.Error(codes.Unauthenticated, "failed to authenticate request: invalid access token")
		}

		validScopes := scopeAuthorizations[info.FullMethod]
		auth, err := db.New().AuthAccessToken(ss.Context(), pgxPool, tokenHash)
		if err != nil {
			if errors.Is(err, pgx.ErrNoRows) {
				return status.Error(codes.Unauthenticated, "failed to authenticate request: invalid access token")
			}

			log.Error().Err(err).Msg("failed to validate access token")
			return status.Error(codes.Internal, "failed to authenticate request")
		}

		for _, vs := range types.FromAccessTokenScopes(validScopes) {
			for _, s := range auth.Scopes {
				if string(vs) == s {
					return handler(srv, ss)
				}
			}
		}
		return status.Error(codes.PermissionDenied, "failed to authenticate request: invalid access token scopes")
	}
}

func UnaryClientInterceptor(accessToken string) grpc.UnaryClientInterceptor {
	return func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
		ctx = metadata.AppendToOutgoingContext(ctx, mdAuthKey, "bearer "+accessToken)
		return invoker(ctx, method, req, reply, cc, opts...)
	}
}

func StreamClientInterceptor(accessToken string) grpc.StreamClientInterceptor {
	return func(ctx context.Context, desc *grpc.StreamDesc, cc *grpc.ClientConn, method string, streamer grpc.Streamer, opts ...grpc.CallOption) (grpc.ClientStream, error) {
		ctx = metadata.AppendToOutgoingContext(ctx, mdAuthKey, "bearer "+accessToken)
		return streamer(ctx, desc, cc, method, opts...)
	}
}

func tokenFromMD(md metadata.MD) (string, string, error) {
	vals := md.Get(mdAuthKey)
	if vals == nil || len(vals) != 1 {
		return "", "", errors.New("invalid access token")
	}
	parts := strings.Split(strings.ToLower(vals[0]), " ")
	if len(parts) != 2 || parts[0] != "bearer" {
		return "", "", errors.New("invalid access token")
	}

	token := parts[1]
	tokenHashBytes := sha512.Sum512([]byte(token))
	tokenHash := hex.EncodeToString(tokenHashBytes[:])

	return token, tokenHash, nil
}
