-- name: RegisterTest :one
INSERT INTO tests (name, run_config, labels, matrix, cron_schedule, next_run_at)
VALUES (
  sqlc.arg('name'),
  sqlc.arg('run_config'),
  sqlc.arg('labels'),
  sqlc.arg('matrix'),
  sqlc.narg('cron_schedule'),
  sqlc.arg('next_run_at')
)
RETURNING *;

-- name: GetTest :one
SELECT *
FROM tests
WHERE tests.id = sqlc.arg('id');

-- name: ListTests :many
SELECT *
FROM tests
ORDER BY tests.name ASC;

-- name: UpdateTest :exec
UPDATE tests
SET
  name = sqlc.arg('name')::varchar,
  run_config = sqlc.arg('run_config')::jsonb,
  labels = sqlc.arg('labels')::jsonb,
  matrix = sqlc.arg('matrix')::jsonb,
  cron_schedule = sqlc.narg('cron_schedule')::varchar,
  next_run_at = sqlc.narg('next_run_at')::timestamptz,
  updated_at = CURRENT_TIMESTAMP
WHERE id = sqlc.arg('id')::uuid;

-- name: ListTestsToSchedule :many
SELECT tests.*
FROM tests
LEFT JOIN runs
ON runs.test_id = tests.id AND runs.result = 'unknown' AND runs.started_at IS NULL
WHERE tests.next_run_at < CURRENT_TIMESTAMP AND runs.id IS NULL
FOR UPDATE OF tests SKIP LOCKED;

-- name: QueryTests :many
SELECT *
FROM tests
WHERE
  (sqlc.narg('ids')::uuid[] IS NULL OR tests.id = ANY (sqlc.narg('ids')::uuid[])) AND
  (sqlc.narg('test_suite_ids')::uuid[] IS NULL OR tests.id = ANY (
    SELECT tests.id
    FROM test_suites
    JOIN tests
    ON tests.labels @> test_suites.labels
    WHERE test_suites.id = ANY (sqlc.narg('test_suite_ids')::uuid[])
    )) AND
  (sqlc.narg('labels')::jsonb IS NULL OR tests.labels @> sqlc.narg('labels')::jsonb)
ORDER BY tests.name ASC;
