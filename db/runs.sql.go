// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0
// source: runs.sql

package db

import (
	"context"
	"database/sql"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgtype"
)

const appendLogsToRun = `-- name: AppendLogsToRun :exec
UPDATE runs
SET logs = COALESCE(logs, '[]'::jsonb) || $1
WHERE id = $2
`

type AppendLogsToRunParams struct {
	Logs pgtype.JSONB
	ID   uuid.UUID
}

func (q *Queries) AppendLogsToRun(ctx context.Context, db DBTX, arg AppendLogsToRunParams) error {
	_, err := db.Exec(ctx, appendLogsToRun, arg.Logs, arg.ID)
	return err
}

const assignRun = `-- name: AssignRun :one
UPDATE runs
SET runner_id = $1::uuid
WHERE runs.id = (
  SELECT matching_runs.id
  FROM runs AS matching_runs
  WHERE
    matching_runs.id = ANY ($2::uuid[]) AND
    matching_runs.runner_id IS NULL
  LIMIT 1
  FOR UPDATE SKIP LOCKED
)
RETURNING id, test_id, test_run_config, test_matrix_id, labels, runner_id, result, logs, result_data, scheduled_at, started_at, finished_at
`

type AssignRunParams struct {
	RunnerID uuid.UUID
	RunIDs   []uuid.UUID
}

func (q *Queries) AssignRun(ctx context.Context, db DBTX, arg AssignRunParams) (Run, error) {
	row := db.QueryRow(ctx, assignRun, arg.RunnerID, arg.RunIDs)
	var i Run
	err := row.Scan(
		&i.ID,
		&i.TestID,
		&i.TestRunConfig,
		&i.TestMatrixID,
		&i.Labels,
		&i.RunnerID,
		&i.Result,
		&i.Logs,
		&i.ResultData,
		&i.ScheduledAt,
		&i.StartedAt,
		&i.FinishedAt,
	)
	return i, err
}

const deleteRunsForTest = `-- name: DeleteRunsForTest :exec
DELETE FROM runs
WHERE test_id = $1
`

func (q *Queries) DeleteRunsForTest(ctx context.Context, db DBTX, testID uuid.UUID) error {
	_, err := db.Exec(ctx, deleteRunsForTest, testID)
	return err
}

const getRun = `-- name: GetRun :one
SELECT runs.id, runs.test_id, runs.test_run_config, runs.test_matrix_id, runs.labels, runs.runner_id, runs.result, runs.logs, runs.result_data, runs.scheduled_at, runs.started_at, runs.finished_at
FROM runs
JOIN tests
ON runs.test_id = tests.id AND tests.namespace = $1
WHERE runs.id = $2
`

type GetRunParams struct {
	Namespace string
	ID        uuid.UUID
}

func (q *Queries) GetRun(ctx context.Context, db DBTX, arg GetRunParams) (Run, error) {
	row := db.QueryRow(ctx, getRun, arg.Namespace, arg.ID)
	var i Run
	err := row.Scan(
		&i.ID,
		&i.TestID,
		&i.TestRunConfig,
		&i.TestMatrixID,
		&i.Labels,
		&i.RunnerID,
		&i.Result,
		&i.Logs,
		&i.ResultData,
		&i.ScheduledAt,
		&i.StartedAt,
		&i.FinishedAt,
	)
	return i, err
}

const listPendingRuns = `-- name: ListPendingRuns :many
SELECT runs.id, runs.test_id, runs.test_run_config, runs.test_matrix_id, runs.labels, runs.runner_id, runs.result, runs.logs, runs.result_data, runs.scheduled_at, runs.started_at, runs.finished_at, tests.namespace
FROM runs
JOIN tests
ON runs.test_id = tests.id
WHERE runner_id IS NULL
`

type ListPendingRunsRow struct {
	ID            uuid.UUID
	TestID        uuid.UUID
	TestRunConfig pgtype.JSONB
	TestMatrixID  uuid.NullUUID
	Labels        pgtype.JSONB
	RunnerID      uuid.NullUUID
	Result        NullRunResult
	Logs          pgtype.JSONB
	ResultData    pgtype.JSONB
	ScheduledAt   sql.NullTime
	StartedAt     sql.NullTime
	FinishedAt    sql.NullTime
	Namespace     string
}

func (q *Queries) ListPendingRuns(ctx context.Context, db DBTX) ([]ListPendingRunsRow, error) {
	rows, err := db.Query(ctx, listPendingRuns)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []ListPendingRunsRow
	for rows.Next() {
		var i ListPendingRunsRow
		if err := rows.Scan(
			&i.ID,
			&i.TestID,
			&i.TestRunConfig,
			&i.TestMatrixID,
			&i.Labels,
			&i.RunnerID,
			&i.Result,
			&i.Logs,
			&i.ResultData,
			&i.ScheduledAt,
			&i.StartedAt,
			&i.FinishedAt,
			&i.Namespace,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listRuns = `-- name: ListRuns :many
SELECT runs.id, runs.test_id, runs.test_run_config, runs.test_matrix_id, runs.labels, runs.runner_id, runs.result, runs.logs, runs.result_data, runs.scheduled_at, runs.started_at, runs.finished_at
FROM runs
JOIN tests
ON runs.test_id = tests.id AND tests.namespace = $1
`

func (q *Queries) ListRuns(ctx context.Context, db DBTX, namespace string) ([]Run, error) {
	rows, err := db.Query(ctx, listRuns, namespace)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Run
	for rows.Next() {
		var i Run
		if err := rows.Scan(
			&i.ID,
			&i.TestID,
			&i.TestRunConfig,
			&i.TestMatrixID,
			&i.Labels,
			&i.RunnerID,
			&i.Result,
			&i.Logs,
			&i.ResultData,
			&i.ScheduledAt,
			&i.StartedAt,
			&i.FinishedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const queryRuns = `-- name: QueryRuns :many
SELECT runs.id, runs.test_id, runs.test_run_config, runs.test_matrix_id, runs.labels, runs.runner_id, runs.result, runs.logs, runs.result_data, runs.scheduled_at, runs.started_at, runs.finished_at
FROM runs
JOIN tests
ON runs.test_id = tests.id AND tests.namespace = $1
WHERE
  ($2::uuid[] IS NULL OR runs.id = ANY ($2::uuid[])) AND
  ($3::uuid[] IS NULL OR runs.test_id = ANY ($3::uuid[])) AND
  ($4::uuid[] IS NULL OR runner_id = ANY ($4::uuid[])) AND
  ($5::text[] IS NULL OR result::text = ANY ($5::text[])) AND
  ($6::timestamptz IS NULL OR scheduled_at < $6::timestamptz) AND
  ($7::timestamptz IS NULL OR scheduled_at > $7::timestamptz) AND
  ($8::timestamptz IS NULL OR started_at < $8::timestamptz) AND
  ($9::timestamptz IS NULL OR started_at > $9::timestamptz) AND
  ($10::timestamptz IS NULL OR finished_at < $10::timestamptz) AND
  ($11::timestamptz IS NULL OR finished_at > $11::timestamptz)
ORDER BY scheduled_at DESC
`

type QueryRunsParams struct {
	Namespace       string
	Ids             []uuid.UUID
	TestIds         []uuid.UUID
	RunnerIds       []uuid.UUID
	Results         []string
	ScheduledBefore sql.NullTime
	ScheduledAfter  sql.NullTime
	StartedBefore   sql.NullTime
	StartedAfter    sql.NullTime
	FinishedBefore  sql.NullTime
	FinishedAfter   sql.NullTime
}

func (q *Queries) QueryRuns(ctx context.Context, db DBTX, arg QueryRunsParams) ([]Run, error) {
	rows, err := db.Query(ctx, queryRuns,
		arg.Namespace,
		arg.Ids,
		arg.TestIds,
		arg.RunnerIds,
		arg.Results,
		arg.ScheduledBefore,
		arg.ScheduledAfter,
		arg.StartedBefore,
		arg.StartedAfter,
		arg.FinishedBefore,
		arg.FinishedAfter,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Run
	for rows.Next() {
		var i Run
		if err := rows.Scan(
			&i.ID,
			&i.TestID,
			&i.TestRunConfig,
			&i.TestMatrixID,
			&i.Labels,
			&i.RunnerID,
			&i.Result,
			&i.Logs,
			&i.ResultData,
			&i.ScheduledAt,
			&i.StartedAt,
			&i.FinishedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const resetOrphanedRuns = `-- name: ResetOrphanedRuns :exec
UPDATE runs
SET runner_id = NULL
WHERE
  result = 'unknown' AND
  started_at IS NULL AND
  scheduled_at < $1::timestamptz
`

func (q *Queries) ResetOrphanedRuns(ctx context.Context, db DBTX, before time.Time) error {
	_, err := db.Exec(ctx, resetOrphanedRuns, before)
	return err
}

const runSummariesForRunner = `-- name: RunSummariesForRunner :many
SELECT runs.id, tests.namespace AS test_namespace, tests.id AS test_id, tests.name AS test_name, runs.test_run_config, runs.labels, runs.runner_id, runs.result, runs.scheduled_at, runs.started_at, runs.finished_at, runs.result_data
FROM runs
JOIN tests
ON runs.test_id = tests.id AND tests.namespace = $1
WHERE runs.runner_id = $2 AND runs.scheduled_at > $3
ORDER by runs.scheduled_at desc
`

type RunSummariesForRunnerParams struct {
	Namespace      string
	RunnerID       uuid.NullUUID
	ScheduledAfter sql.NullTime
}

type RunSummariesForRunnerRow struct {
	ID            uuid.UUID
	TestNamespace string
	TestID        uuid.UUID
	TestName      string
	TestRunConfig pgtype.JSONB
	Labels        pgtype.JSONB
	RunnerID      uuid.NullUUID
	Result        NullRunResult
	ScheduledAt   sql.NullTime
	StartedAt     sql.NullTime
	FinishedAt    sql.NullTime
	ResultData    pgtype.JSONB
}

func (q *Queries) RunSummariesForRunner(ctx context.Context, db DBTX, arg RunSummariesForRunnerParams) ([]RunSummariesForRunnerRow, error) {
	rows, err := db.Query(ctx, runSummariesForRunner, arg.Namespace, arg.RunnerID, arg.ScheduledAfter)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []RunSummariesForRunnerRow
	for rows.Next() {
		var i RunSummariesForRunnerRow
		if err := rows.Scan(
			&i.ID,
			&i.TestNamespace,
			&i.TestID,
			&i.TestName,
			&i.TestRunConfig,
			&i.Labels,
			&i.RunnerID,
			&i.Result,
			&i.ScheduledAt,
			&i.StartedAt,
			&i.FinishedAt,
			&i.ResultData,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const runSummariesForTest = `-- name: RunSummariesForTest :many
SELECT runs.id, tests.namespace AS test_namespace, tests.id AS test_id, tests.name AS test_name, runs.test_run_config, runs.labels, runs.runner_id, runs.result, runs.scheduled_at, runs.started_at, runs.finished_at, runs.result_data
FROM runs
JOIN tests
ON runs.test_id = tests.id AND tests.namespace = $1
WHERE runs.test_id = $2 AND runs.scheduled_at > $3
ORDER by runs.scheduled_at desc
`

type RunSummariesForTestParams struct {
	Namespace      string
	TestID         uuid.UUID
	ScheduledAfter sql.NullTime
}

type RunSummariesForTestRow struct {
	ID            uuid.UUID
	TestNamespace string
	TestID        uuid.UUID
	TestName      string
	TestRunConfig pgtype.JSONB
	Labels        pgtype.JSONB
	RunnerID      uuid.NullUUID
	Result        NullRunResult
	ScheduledAt   sql.NullTime
	StartedAt     sql.NullTime
	FinishedAt    sql.NullTime
	ResultData    pgtype.JSONB
}

func (q *Queries) RunSummariesForTest(ctx context.Context, db DBTX, arg RunSummariesForTestParams) ([]RunSummariesForTestRow, error) {
	rows, err := db.Query(ctx, runSummariesForTest, arg.Namespace, arg.TestID, arg.ScheduledAfter)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []RunSummariesForTestRow
	for rows.Next() {
		var i RunSummariesForTestRow
		if err := rows.Scan(
			&i.ID,
			&i.TestNamespace,
			&i.TestID,
			&i.TestName,
			&i.TestRunConfig,
			&i.Labels,
			&i.RunnerID,
			&i.Result,
			&i.ScheduledAt,
			&i.StartedAt,
			&i.FinishedAt,
			&i.ResultData,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const scheduleRun = `-- name: ScheduleRun :one
INSERT INTO runs (test_id, test_run_config, labels, test_matrix_id)
SELECT tests.id, tests.run_config, $1, $2
FROM tests
WHERE tests.id = $3 AND tests.namespace = $4
RETURNING id, test_id, test_run_config, test_matrix_id, labels, runner_id, result, logs, result_data, scheduled_at, started_at, finished_at
`

type ScheduleRunParams struct {
	Labels       pgtype.JSONB
	TestMatrixID uuid.NullUUID
	TestID       uuid.UUID
	Namespace    string
}

func (q *Queries) ScheduleRun(ctx context.Context, db DBTX, arg ScheduleRunParams) (Run, error) {
	row := db.QueryRow(ctx, scheduleRun,
		arg.Labels,
		arg.TestMatrixID,
		arg.TestID,
		arg.Namespace,
	)
	var i Run
	err := row.Scan(
		&i.ID,
		&i.TestID,
		&i.TestRunConfig,
		&i.TestMatrixID,
		&i.Labels,
		&i.RunnerID,
		&i.Result,
		&i.Logs,
		&i.ResultData,
		&i.ScheduledAt,
		&i.StartedAt,
		&i.FinishedAt,
	)
	return i, err
}

const summarizeRunsBreakdownResult = `-- name: SummarizeRunsBreakdownResult :many
WITH intervals AS (
  SELECT generate_series(
    date_trunc($1, $2::timestamptz) + make_interval(secs => $5),
    date_trunc($1, $3::timestamptz),
    make_interval(secs => $5)
  ) as start
)
SELECT 
  intervals.start::timestamptz,
  COUNT(runs.id) FILTER (WHERE result = 'pass') as pass,
  COUNT(runs.id) FILTER (WHERE result = 'fail') as fail,
  COUNT(runs.id) FILTER (WHERE result = 'error') as error,
  COUNT(runs.id) FILTER (WHERE result = 'unknown') as unknown
FROM intervals
LEFT JOIN runs
ON
  intervals.start = date_trunc($1, runs.scheduled_at) AND
  runs.scheduled_at > $2 AND
  runs.scheduled_at < $3
JOIN tests
ON runs.test_id = tests.id AND tests.namespace = $4
GROUP BY intervals.start
ORDER BY intervals.start ASC
`

type SummarizeRunsBreakdownResultParams struct {
	Precision string
	StartTime sql.NullTime
	EndTime   sql.NullTime
	Namespace string
	Interval  float64
}

type SummarizeRunsBreakdownResultRow struct {
	IntervalsStart time.Time
	Pass           int64
	Fail           int64
	Error          int64
	Unknown        int64
}

func (q *Queries) SummarizeRunsBreakdownResult(ctx context.Context, db DBTX, arg SummarizeRunsBreakdownResultParams) ([]SummarizeRunsBreakdownResultRow, error) {
	rows, err := db.Query(ctx, summarizeRunsBreakdownResult,
		arg.Precision,
		arg.StartTime,
		arg.EndTime,
		arg.Namespace,
		arg.Interval,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []SummarizeRunsBreakdownResultRow
	for rows.Next() {
		var i SummarizeRunsBreakdownResultRow
		if err := rows.Scan(
			&i.IntervalsStart,
			&i.Pass,
			&i.Fail,
			&i.Error,
			&i.Unknown,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const summarizeRunsBreakdownTest = `-- name: SummarizeRunsBreakdownTest :many
WITH intervals AS (
  SELECT generate_series(
    date_trunc($1, $2::timestamptz) + make_interval(secs => $5),
    date_trunc($1, $3::timestamptz),
    make_interval(secs => $5)
  ) as start
)
SELECT 
  intervals.start::timestamptz,
  tests.id,
  tests.name,
  COUNT(tests.id) FILTER (WHERE runs.result = 'pass') as pass,
  COUNT(tests.id) FILTER (WHERE runs.result = 'fail') as fail,
  COUNT(tests.id) FILTER (WHERE runs.result = 'error') as error,
  COUNT(tests.id) FILTER (WHERE runs.result = 'unknown') as unknown
FROM intervals
LEFT JOIN runs
ON
  intervals.start = date_trunc($1, runs.scheduled_at) AND
  runs.scheduled_at > $2 AND
  runs.scheduled_at < $3
JOIN tests
ON runs.test_id = tests.id AND tests.namespace = $4
GROUP BY intervals.start, tests.id
ORDER BY intervals.start ASC
`

type SummarizeRunsBreakdownTestParams struct {
	Precision string
	StartTime sql.NullTime
	EndTime   sql.NullTime
	Namespace string
	Interval  float64
}

type SummarizeRunsBreakdownTestRow struct {
	IntervalsStart time.Time
	ID             uuid.UUID
	Name           string
	Pass           int64
	Fail           int64
	Error          int64
	Unknown        int64
}

func (q *Queries) SummarizeRunsBreakdownTest(ctx context.Context, db DBTX, arg SummarizeRunsBreakdownTestParams) ([]SummarizeRunsBreakdownTestRow, error) {
	rows, err := db.Query(ctx, summarizeRunsBreakdownTest,
		arg.Precision,
		arg.StartTime,
		arg.EndTime,
		arg.Namespace,
		arg.Interval,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []SummarizeRunsBreakdownTestRow
	for rows.Next() {
		var i SummarizeRunsBreakdownTestRow
		if err := rows.Scan(
			&i.IntervalsStart,
			&i.ID,
			&i.Name,
			&i.Pass,
			&i.Fail,
			&i.Error,
			&i.Unknown,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const timeoutRuns = `-- name: TimeoutRuns :exec
UPDATE runs
SET
  result = 'error',
  finished_at = CURRENT_TIMESTAMP,
  logs = COALESCE(logs, '[]'::jsonb) || $1
WHERE
  result = 'unknown' AND
  runner_id IS NOT NULL AND
  CURRENT_TIMESTAMP > started_at + make_interval(secs => COALESCE(test_run_config['timeout_seconds']::int, $2::int))
`

type TimeoutRunsParams struct {
	TimeoutLog     pgtype.JSONB
	DefaultTimeout int32
}

func (q *Queries) TimeoutRuns(ctx context.Context, db DBTX, arg TimeoutRunsParams) error {
	_, err := db.Exec(ctx, timeoutRuns, arg.TimeoutLog, arg.DefaultTimeout)
	return err
}

const updateResultData = `-- name: UpdateResultData :exec
UPDATE runs
SET result_data = result_data || $1::jsonb 
WHERE id = $2
`

type UpdateResultDataParams struct {
	ResultData pgtype.JSONB
	ID         uuid.UUID
}

func (q *Queries) UpdateResultData(ctx context.Context, db DBTX, arg UpdateResultDataParams) error {
	_, err := db.Exec(ctx, updateResultData, arg.ResultData, arg.ID)
	return err
}

const updateRun = `-- name: UpdateRun :exec
UPDATE runs
SET
  result = $1,
  started_at = $2::timestamptz,
  finished_at = $3::timestamptz
WHERE id = $4
`

type UpdateRunParams struct {
	Result     NullRunResult
	StartedAt  sql.NullTime
	FinishedAt sql.NullTime
	ID         uuid.UUID
}

func (q *Queries) UpdateRun(ctx context.Context, db DBTX, arg UpdateRunParams) error {
	_, err := db.Exec(ctx, updateRun,
		arg.Result,
		arg.StartedAt,
		arg.FinishedAt,
		arg.ID,
	)
	return err
}
