// Code generated by pggen. DO NOT EDIT.

package db

import (
	"context"
	"fmt"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgtype"
	"github.com/jackc/pgx/v4"
)

const defineTestSuiteSQL = `INSERT INTO test_suites (name, labels)
VALUES ($1::varchar, $2::jsonb)
RETURNING *;`

type DefineTestSuiteRow struct {
	ID         string             `json:"id"`
	Name       string             `json:"name"`
	Labels     pgtype.JSONB       `json:"labels"`
	CreatedAt  pgtype.Timestamptz `json:"created_at"`
	UpdatedAt  pgtype.Timestamptz `json:"updated_at"`
	ArchivedAt pgtype.Timestamptz `json:"archived_at"`
}

// DefineTestSuite implements Querier.DefineTestSuite.
func (q *DBQuerier) DefineTestSuite(ctx context.Context, name string, labels pgtype.JSONB) (DefineTestSuiteRow, error) {
	ctx = context.WithValue(ctx, "pggen_query_name", "DefineTestSuite")
	row := q.conn.QueryRow(ctx, defineTestSuiteSQL, name, labels)
	var item DefineTestSuiteRow
	if err := row.Scan(&item.ID, &item.Name, &item.Labels, &item.CreatedAt, &item.UpdatedAt, &item.ArchivedAt); err != nil {
		return item, fmt.Errorf("query DefineTestSuite: %w", err)
	}
	return item, nil
}

// DefineTestSuiteBatch implements Querier.DefineTestSuiteBatch.
func (q *DBQuerier) DefineTestSuiteBatch(batch genericBatch, name string, labels pgtype.JSONB) {
	batch.Queue(defineTestSuiteSQL, name, labels)
}

// DefineTestSuiteScan implements Querier.DefineTestSuiteScan.
func (q *DBQuerier) DefineTestSuiteScan(results pgx.BatchResults) (DefineTestSuiteRow, error) {
	row := results.QueryRow()
	var item DefineTestSuiteRow
	if err := row.Scan(&item.ID, &item.Name, &item.Labels, &item.CreatedAt, &item.UpdatedAt, &item.ArchivedAt); err != nil {
		return item, fmt.Errorf("scan DefineTestSuiteBatch row: %w", err)
	}
	return item, nil
}

const updateTestSuiteSQL = `UPDATE test_suites
SET
  name = $1::varchar,
  labels = $2::jsonb,
  updated_at = CURRENT_TIMESTAMP
WHERE id = $3;`

type UpdateTestSuiteParams struct {
	Name   string
	Labels pgtype.JSONB
	ID     string
}

// UpdateTestSuite implements Querier.UpdateTestSuite.
func (q *DBQuerier) UpdateTestSuite(ctx context.Context, params UpdateTestSuiteParams) (pgconn.CommandTag, error) {
	ctx = context.WithValue(ctx, "pggen_query_name", "UpdateTestSuite")
	cmdTag, err := q.conn.Exec(ctx, updateTestSuiteSQL, params.Name, params.Labels, params.ID)
	if err != nil {
		return cmdTag, fmt.Errorf("exec query UpdateTestSuite: %w", err)
	}
	return cmdTag, err
}

// UpdateTestSuiteBatch implements Querier.UpdateTestSuiteBatch.
func (q *DBQuerier) UpdateTestSuiteBatch(batch genericBatch, params UpdateTestSuiteParams) {
	batch.Queue(updateTestSuiteSQL, params.Name, params.Labels, params.ID)
}

// UpdateTestSuiteScan implements Querier.UpdateTestSuiteScan.
func (q *DBQuerier) UpdateTestSuiteScan(results pgx.BatchResults) (pgconn.CommandTag, error) {
	cmdTag, err := results.Exec()
	if err != nil {
		return cmdTag, fmt.Errorf("exec UpdateTestSuiteBatch: %w", err)
	}
	return cmdTag, err
}

const getTestSuiteSQL = `SELECT *
FROM test_suites
WHERE id = $1::uuid;`

type GetTestSuiteRow struct {
	ID         string             `json:"id"`
	Name       string             `json:"name"`
	Labels     pgtype.JSONB       `json:"labels"`
	CreatedAt  pgtype.Timestamptz `json:"created_at"`
	UpdatedAt  pgtype.Timestamptz `json:"updated_at"`
	ArchivedAt pgtype.Timestamptz `json:"archived_at"`
}

// GetTestSuite implements Querier.GetTestSuite.
func (q *DBQuerier) GetTestSuite(ctx context.Context, id string) (GetTestSuiteRow, error) {
	ctx = context.WithValue(ctx, "pggen_query_name", "GetTestSuite")
	row := q.conn.QueryRow(ctx, getTestSuiteSQL, id)
	var item GetTestSuiteRow
	if err := row.Scan(&item.ID, &item.Name, &item.Labels, &item.CreatedAt, &item.UpdatedAt, &item.ArchivedAt); err != nil {
		return item, fmt.Errorf("query GetTestSuite: %w", err)
	}
	return item, nil
}

// GetTestSuiteBatch implements Querier.GetTestSuiteBatch.
func (q *DBQuerier) GetTestSuiteBatch(batch genericBatch, id string) {
	batch.Queue(getTestSuiteSQL, id)
}

// GetTestSuiteScan implements Querier.GetTestSuiteScan.
func (q *DBQuerier) GetTestSuiteScan(results pgx.BatchResults) (GetTestSuiteRow, error) {
	row := results.QueryRow()
	var item GetTestSuiteRow
	if err := row.Scan(&item.ID, &item.Name, &item.Labels, &item.CreatedAt, &item.UpdatedAt, &item.ArchivedAt); err != nil {
		return item, fmt.Errorf("scan GetTestSuiteBatch row: %w", err)
	}
	return item, nil
}

const listTestSuitesSQL = `SELECT *
FROM test_suites
WHERE labels @> $1::jsonb
ORDER BY name ASC;`

type ListTestSuitesRow struct {
	ID         string             `json:"id"`
	Name       string             `json:"name"`
	Labels     pgtype.JSONB       `json:"labels"`
	CreatedAt  pgtype.Timestamptz `json:"created_at"`
	UpdatedAt  pgtype.Timestamptz `json:"updated_at"`
	ArchivedAt pgtype.Timestamptz `json:"archived_at"`
}

// ListTestSuites implements Querier.ListTestSuites.
func (q *DBQuerier) ListTestSuites(ctx context.Context, labels pgtype.JSONB) ([]ListTestSuitesRow, error) {
	ctx = context.WithValue(ctx, "pggen_query_name", "ListTestSuites")
	rows, err := q.conn.Query(ctx, listTestSuitesSQL, labels)
	if err != nil {
		return nil, fmt.Errorf("query ListTestSuites: %w", err)
	}
	defer rows.Close()
	items := []ListTestSuitesRow{}
	for rows.Next() {
		var item ListTestSuitesRow
		if err := rows.Scan(&item.ID, &item.Name, &item.Labels, &item.CreatedAt, &item.UpdatedAt, &item.ArchivedAt); err != nil {
			return nil, fmt.Errorf("scan ListTestSuites row: %w", err)
		}
		items = append(items, item)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("close ListTestSuites rows: %w", err)
	}
	return items, err
}

// ListTestSuitesBatch implements Querier.ListTestSuitesBatch.
func (q *DBQuerier) ListTestSuitesBatch(batch genericBatch, labels pgtype.JSONB) {
	batch.Queue(listTestSuitesSQL, labels)
}

// ListTestSuitesScan implements Querier.ListTestSuitesScan.
func (q *DBQuerier) ListTestSuitesScan(results pgx.BatchResults) ([]ListTestSuitesRow, error) {
	rows, err := results.Query()
	if err != nil {
		return nil, fmt.Errorf("query ListTestSuitesBatch: %w", err)
	}
	defer rows.Close()
	items := []ListTestSuitesRow{}
	for rows.Next() {
		var item ListTestSuitesRow
		if err := rows.Scan(&item.ID, &item.Name, &item.Labels, &item.CreatedAt, &item.UpdatedAt, &item.ArchivedAt); err != nil {
			return nil, fmt.Errorf("scan ListTestSuitesBatch row: %w", err)
		}
		items = append(items, item)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("close ListTestSuitesBatch rows: %w", err)
	}
	return items, err
}

const archiveTestSuiteSQL = `UPDATE test_suites
SET archived_at = CURRENT_TIMESTAMP
WHERE id = $1;`

// ArchiveTestSuite implements Querier.ArchiveTestSuite.
func (q *DBQuerier) ArchiveTestSuite(ctx context.Context, id string) (pgconn.CommandTag, error) {
	ctx = context.WithValue(ctx, "pggen_query_name", "ArchiveTestSuite")
	cmdTag, err := q.conn.Exec(ctx, archiveTestSuiteSQL, id)
	if err != nil {
		return cmdTag, fmt.Errorf("exec query ArchiveTestSuite: %w", err)
	}
	return cmdTag, err
}

// ArchiveTestSuiteBatch implements Querier.ArchiveTestSuiteBatch.
func (q *DBQuerier) ArchiveTestSuiteBatch(batch genericBatch, id string) {
	batch.Queue(archiveTestSuiteSQL, id)
}

// ArchiveTestSuiteScan implements Querier.ArchiveTestSuiteScan.
func (q *DBQuerier) ArchiveTestSuiteScan(results pgx.BatchResults) (pgconn.CommandTag, error) {
	cmdTag, err := results.Exec()
	if err != nil {
		return cmdTag, fmt.Errorf("exec ArchiveTestSuiteBatch: %w", err)
	}
	return cmdTag, err
}
