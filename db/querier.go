// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.14.0

package db

import (
	"context"
	"database/sql"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgtype"
)

type Querier interface {
	AppendLogsToRun(ctx context.Context, db DBTX, arg AppendLogsToRunParams) error
	ArchiveTest(ctx context.Context, db DBTX, id uuid.UUID) error
	ArchiveTestSuite(ctx context.Context, db DBTX, id uuid.UUID) error
	AssignRun(ctx context.Context, db DBTX, arg AssignRunParams) (AssignRunRow, error)
	// TODO re: ::text[] https://github.com/kyleconroy/sqlc/issues/1256
	AuthAccessToken(ctx context.Context, db DBTX, tokenHash string) (AuthAccessTokenRow, error)
	CreateTestRunConfig(ctx context.Context, db DBTX, arg CreateTestRunConfigParams) (TestRunConfig, error)
	DefineTestSuite(ctx context.Context, db DBTX, arg DefineTestSuiteParams) (TestSuite, error)
	// TODO re: ::text[] https://github.com/kyleconroy/sqlc/issues/1256
	GetAccessToken(ctx context.Context, db DBTX, id uuid.UUID) (GetAccessTokenRow, error)
	GetRun(ctx context.Context, db DBTX, id uuid.UUID) (GetRunRow, error)
	GetRunner(ctx context.Context, db DBTX, id uuid.UUID) (Runner, error)
	GetTest(ctx context.Context, db DBTX, id uuid.UUID) (GetTestRow, error)
	GetTestSuite(ctx context.Context, db DBTX, id uuid.UUID) (TestSuite, error)
	// TODO re: ::text[] https://github.com/kyleconroy/sqlc/issues/1256
	IssueAccessToken(ctx context.Context, db DBTX, arg IssueAccessTokenParams) (IssueAccessTokenRow, error)
	// TODO re: ::text[] https://github.com/kyleconroy/sqlc/issues/1256
	ListAccessTokens(ctx context.Context, db DBTX, arg ListAccessTokensParams) ([]ListAccessTokensRow, error)
	ListRunners(ctx context.Context, db DBTX, heartbeatSince sql.NullTime) ([]Runner, error)
	ListRuns(ctx context.Context, db DBTX, arg ListRunsParams) ([]ListRunsRow, error)
	ListTestSuites(ctx context.Context, db DBTX, labels pgtype.JSONB) ([]TestSuite, error)
	ListTests(ctx context.Context, db DBTX, labels pgtype.JSONB) ([]ListTestsRow, error)
	ListTestsIDsMatchingLabelKeys(ctx context.Context, db DBTX, arg ListTestsIDsMatchingLabelKeysParams) ([]ListTestsIDsMatchingLabelKeysRow, error)
	ListTestsToSchedule(ctx context.Context, db DBTX) ([]ListTestsToScheduleRow, error)
	QueryRunners(ctx context.Context, db DBTX, arg QueryRunnersParams) ([]Runner, error)
	QueryRuns(ctx context.Context, db DBTX, arg QueryRunsParams) ([]QueryRunsRow, error)
	QueryTestSuites(ctx context.Context, db DBTX, arg QueryTestSuitesParams) ([]TestSuite, error)
	QueryTests(ctx context.Context, db DBTX, arg QueryTestsParams) ([]QueryTestsRow, error)
	RegisterRunner(ctx context.Context, db DBTX, arg RegisterRunnerParams) (Runner, error)
	RegisterTest(ctx context.Context, db DBTX, arg RegisterTestParams) (RegisterTestRow, error)
	ResetOrphanedRuns(ctx context.Context, db DBTX, before time.Time) error
	RevokeAccessToken(ctx context.Context, db DBTX, id uuid.UUID) error
	RunSummaryForRunner(ctx context.Context, db DBTX, arg RunSummaryForRunnerParams) ([]RunSummaryForRunnerRow, error)
	RunSummaryForTest(ctx context.Context, db DBTX, arg RunSummaryForTestParams) ([]RunSummaryForTestRow, error)
	ScheduleRun(ctx context.Context, db DBTX, arg ScheduleRunParams) (ScheduleRunRow, error)
	UIListRecentRuns(ctx context.Context, db DBTX, limit int32) ([]UIListRecentRunsRow, error)
	UIListTests(ctx context.Context, db DBTX) ([]Test, error)
	UIRunnerSummary(ctx context.Context, db DBTX, arg UIRunnerSummaryParams) ([]UIRunnerSummaryRow, error)
	UIRunsSummary(ctx context.Context, db DBTX, arg UIRunsSummaryParams) ([]UIRunsSummaryRow, error)
	UITestResults(ctx context.Context, db DBTX) ([]UITestResultsRow, error)
	UITestsByLabels(ctx context.Context, db DBTX) ([]UITestsByLabelsRow, error)
	UpdateResultData(ctx context.Context, db DBTX, arg UpdateResultDataParams) error
	UpdateRun(ctx context.Context, db DBTX, arg UpdateRunParams) error
	UpdateRunnerHeartbeat(ctx context.Context, db DBTX, id uuid.UUID) error
	UpdateTest(ctx context.Context, db DBTX, arg UpdateTestParams) error
	UpdateTestSuite(ctx context.Context, db DBTX, arg UpdateTestSuiteParams) error
}

var _ Querier = (*Queries)(nil)
