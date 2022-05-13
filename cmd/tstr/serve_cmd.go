package main

import (
	"context"
	"crypto/sha512"
	"database/sql"
	"encoding/hex"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	grpc_validator "github.com/grpc-ecosystem/go-grpc-middleware/validator"
	"github.com/jackc/pgx/v4/pgxpool"
	grpczerolog "github.com/jwreagor/grpc-zerolog"
	"github.com/nanzhong/tstr/api/admin/v1"
	"github.com/nanzhong/tstr/api/control/v1"
	"github.com/nanzhong/tstr/api/runner/v1"
	"github.com/nanzhong/tstr/db"
	"github.com/nanzhong/tstr/grpc/auth"
	"github.com/nanzhong/tstr/grpc/server"
	"github.com/nanzhong/tstr/webui"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/hlog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
)

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "serve the gRPC API and web UI.",
	Args:  cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		ctx := context.Background()

		// TODO Use console writer for now for easy development/debugging, perhaps remove and rely on humanlog in the future.
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
		grpclog.SetLoggerV2(grpczerolog.New(log.Logger.With().Str("component", "client-grpc").Logger()))

		pgxPool, err := pgxpool.Connect(context.Background(), viper.GetString("serve.pg-dsn"))
		if err != nil {
			log.Fatal().Err(err).Msg("failed to connect to pg")
		}
		defer pgxPool.Close()

		if viper.GetString("serve.bootstrap-token") != "" {
			tokenHashBytes := sha512.Sum512([]byte(viper.GetString("serve.bootstrap-token")))
			tokenHash := hex.EncodeToString(tokenHashBytes[:])

			var textScopes []string
			for _, s := range []db.AccessTokenScope{db.AccessTokenScopeAdmin, db.AccessTokenScopeControlRw, db.AccessTokenScopeRunner} {
				textScopes = append(textScopes, string(s))
			}
			_, err := db.New().IssueAccessToken(ctx, pgxPool, db.IssueAccessTokenParams{
				Name:      "bootstrap-token",
				TokenHash: tokenHash,
				Scopes:    textScopes,
				ExpiresAt: sql.NullTime{Valid: true, Time: time.Now().Add(24 * time.Hour)},
			})
			if err != nil {
				log.Fatal().
					Err(err).
					Msg("failed to issue bootstrap token")
			}
		}

		// TODO: consider using cmux to serve http and grpc on the same port?
		grpcListener, err := net.Listen("tcp", viper.GetString("serve.api-addr"))
		if err != nil {
			log.Fatal().
				Err(err).
				Str("api-addr", viper.GetString("serve.api-addr")).
				Msg("failed to listen on api addr")
		}

		webListener, err := net.Listen("tcp", viper.GetString("serve.web-addr"))
		if err != nil {
			log.Fatal().
				Err(err).
				Str("web-addr", viper.GetString("serve.web-addr")).
				Msg("failed to listen on web addr")
		}

		grpcServer := grpc.NewServer(
			grpc.ChainUnaryInterceptor(
				grpc_validator.UnaryServerInterceptor(),
				auth.UnaryServerInterceptor(pgxPool),
			),
			grpc.ChainStreamInterceptor(
				grpc_validator.StreamServerInterceptor(),
				auth.StreamServerInterceptor(pgxPool),
			),
		)

		controlServer := server.NewControlServer(pgxPool)
		control.RegisterControlServiceServer(grpcServer, controlServer)

		adminServer := server.NewAdminServer(pgxPool)
		admin.RegisterAdminServiceServer(grpcServer, adminServer)

		runnerServer := server.NewRunnerServer(pgxPool)
		runner.RegisterRunnerServiceServer(grpcServer, runnerServer)

		webui := webui.NewWebUI()

		httpServer := http.Server{
			Handler: hlog.NewHandler(log.Logger)(webui.Handler()),
		}
		ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
		defer stop()

		go func() {
			<-ctx.Done()

			log.Info().Msg("shutting down")

			// Give one minute for running requests to complete
			shutdownCtx, cancel := context.WithTimeout(context.Background(), 1*time.Minute)
			defer cancel()

			var eg errgroup.Group
			eg.Go(func() error {
				log.Info().Msg("attempting to shutdown grpc server")
				grpcServer.GracefulStop()
				return nil
			})
			eg.Go(func() error {
				log.Info().Msg("attempting to shutdown http server")
				return httpServer.Shutdown(shutdownCtx)
			})
			err := eg.Wait()
			if err != nil {
				log.Error().Err(err).Msg("failed to gracefully shutdown")
			}
		}()

		log.Info().Msg("tstr starting")
		var eg errgroup.Group
		eg.Go(func() error {
			log.Info().
				Str("api-addr", viper.GetString("serve.api-addr")).
				Msg("starting grpc server")
			return grpcServer.Serve(grpcListener)
		})
		eg.Go(func() error {
			log.Info().
				Str("web-addr", viper.GetString("serve.web-addr")).
				Msg("starting http server")
			return httpServer.Serve(webListener)
		})
		err = eg.Wait()
		log.Info().Msg("tstr shutdown")
	},
}

func init() {
	serveCmd.Flags().String("api-addr", "0.0.0.0:9000", "The address to serve the gRPC API on.")
	viper.BindPFlag("serve.api-addr", serveCmd.Flags().Lookup("api-addr"))

	serveCmd.Flags().String("web-addr", "0.0.0.0:9090", "The address to serve the web UI on.")
	viper.BindPFlag("serve.web-addr", serveCmd.Flags().Lookup("web-addr"))

	serveCmd.Flags().String("pg-dsn", "", "The PostgreSQL DSN to use.")
	viper.BindPFlag("serve.pg-dsn", serveCmd.Flags().Lookup("pg-dsn"))

	serveCmd.Flags().String("bootstrap-token", "", "Bootstrap with provided access token (note that this token will have admin scope valid for 24h).")
	viper.BindPFlag("serve.bootstrap-token", serveCmd.Flags().Lookup("bootstrap-token"))

	rootCmd.AddCommand(serveCmd)
}
