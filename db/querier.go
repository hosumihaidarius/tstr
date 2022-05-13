// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.13.0

package db

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
	"github.com/jackc/pgtype"
)

type Querier interface {
	ArchiveTest(ctx context.Context, db DBTX, id uuid.UUID) error
	ArchiveTestSuite(ctx context.Context, db DBTX, id uuid.UUID) error
	AssignRun(ctx context.Context, db DBTX, arg AssignRunParams) (AssignRunRow, error)
	AuthAccessToken(ctx context.Context, db DBTX, tokenHash string) (AuthAccessTokenRow, error)
	CreateTestRunConfig(ctx context.Context, db DBTX, arg CreateTestRunConfigParams) (TestRunConfig, error)
	DefineTestSuite(ctx context.Context, db DBTX, arg DefineTestSuiteParams) (TestSuite, error)
	GetAccessToken(ctx context.Context, db DBTX, id uuid.UUID) (GetAccessTokenRow, error)
	GetRun(ctx context.Context, db DBTX, id uuid.UUID) (GetRunRow, error)
	GetRunner(ctx context.Context, db DBTX, id uuid.UUID) (Runner, error)
	GetTest(ctx context.Context, db DBTX, id uuid.UUID) (GetTestRow, error)
	GetTestSuite(ctx context.Context, db DBTX, id uuid.UUID) (TestSuite, error)
	// TODO re: ::text[] https://github.com/kyleconroy/sqlc/issues/1256
	IssueAccessToken(ctx context.Context, db DBTX, arg IssueAccessTokenParams) (IssueAccessTokenRow, error)
	ListAccessTokens(ctx context.Context, db DBTX, arg ListAccessTokensParams) ([]ListAccessTokensRow, error)
	ListRunners(ctx context.Context, db DBTX, heartbeatSince sql.NullTime) ([]Runner, error)
	ListRuns(ctx context.Context, db DBTX, arg ListRunsParams) ([]ListRunsRow, error)
	ListTestSuites(ctx context.Context, db DBTX, labels pgtype.JSONB) ([]TestSuite, error)
	ListTests(ctx context.Context, db DBTX, labels pgtype.JSONB) ([]ListTestsRow, error)
	ListTestsIDsMatchingLabelKeys(ctx context.Context, db DBTX, arg ListTestsIDsMatchingLabelKeysParams) ([]ListTestsIDsMatchingLabelKeysRow, error)
	ListTestsToSchedule(ctx context.Context, db DBTX) ([]ListTestsToScheduleRow, error)
	RegisterRunner(ctx context.Context, db DBTX, arg RegisterRunnerParams) (Runner, error)
	RegisterTest(ctx context.Context, db DBTX, arg RegisterTestParams) (RegisterTestRow, error)
	RevokeAccessToken(ctx context.Context, db DBTX, id uuid.UUID) error
	ScheduleRun(ctx context.Context, db DBTX, arg ScheduleRunParams) (ScheduleRunRow, error)
	UpdateRun(ctx context.Context, db DBTX, arg UpdateRunParams) error
	UpdateTest(ctx context.Context, db DBTX, arg UpdateTestParams) error
	UpdateTestSuite(ctx context.Context, db DBTX, arg UpdateTestSuiteParams) error
}

var _ Querier = (*Queries)(nil)
