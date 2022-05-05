// Code generated by pggen. DO NOT EDIT.

package db

import (
	"context"
	"fmt"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgtype"
	"github.com/jackc/pgx/v4"
)

const issueAccessTokenSQL = `INSERT INTO access_tokens (name, token_hash, scopes, expires_at)
VALUES ($1, $2, $3, $4)
RETURNING id, name, scopes, issued_at, expires_at;`

type IssueAccessTokenParams struct {
	Name      string
	TokenHash string
	Scopes    []AccessTokenScope
	ExpiresAt pgtype.Timestamptz
}

type IssueAccessTokenRow struct {
	ID        string             `json:"id"`
	Name      string             `json:"name"`
	Scopes    []AccessTokenScope `json:"scopes"`
	IssuedAt  pgtype.Timestamptz `json:"issued_at"`
	ExpiresAt pgtype.Timestamptz `json:"expires_at"`
}

// IssueAccessToken implements Querier.IssueAccessToken.
func (q *DBQuerier) IssueAccessToken(ctx context.Context, params IssueAccessTokenParams) (IssueAccessTokenRow, error) {
	ctx = context.WithValue(ctx, "pggen_query_name", "IssueAccessToken")
	row := q.conn.QueryRow(ctx, issueAccessTokenSQL, params.Name, params.TokenHash, q.types.newAccessTokenScopeArrayInit(params.Scopes), params.ExpiresAt)
	var item IssueAccessTokenRow
	scopesArray := q.types.newAccessTokenScopeArray()
	if err := row.Scan(&item.ID, &item.Name, scopesArray, &item.IssuedAt, &item.ExpiresAt); err != nil {
		return item, fmt.Errorf("query IssueAccessToken: %w", err)
	}
	if err := scopesArray.AssignTo(&item.Scopes); err != nil {
		return item, fmt.Errorf("assign IssueAccessToken row: %w", err)
	}
	return item, nil
}

// IssueAccessTokenBatch implements Querier.IssueAccessTokenBatch.
func (q *DBQuerier) IssueAccessTokenBatch(batch genericBatch, params IssueAccessTokenParams) {
	batch.Queue(issueAccessTokenSQL, params.Name, params.TokenHash, q.types.newAccessTokenScopeArrayInit(params.Scopes), params.ExpiresAt)
}

// IssueAccessTokenScan implements Querier.IssueAccessTokenScan.
func (q *DBQuerier) IssueAccessTokenScan(results pgx.BatchResults) (IssueAccessTokenRow, error) {
	row := results.QueryRow()
	var item IssueAccessTokenRow
	scopesArray := q.types.newAccessTokenScopeArray()
	if err := row.Scan(&item.ID, &item.Name, scopesArray, &item.IssuedAt, &item.ExpiresAt); err != nil {
		return item, fmt.Errorf("scan IssueAccessTokenBatch row: %w", err)
	}
	if err := scopesArray.AssignTo(&item.Scopes); err != nil {
		return item, fmt.Errorf("assign IssueAccessToken row: %w", err)
	}
	return item, nil
}

const validateAccessTokenSQL = `SELECT EXISTS(
  SELECT TRUE
  FROM access_tokens
  WHERE
    token_hash = $1 AND
    (expires_at IS NULL OR expires_at > CURRENT_TIMESTAMP) AND
    scopes @> $2);`

// ValidateAccessToken implements Querier.ValidateAccessToken.
func (q *DBQuerier) ValidateAccessToken(ctx context.Context, tokenHash string, scopes []AccessTokenScope) (bool, error) {
	ctx = context.WithValue(ctx, "pggen_query_name", "ValidateAccessToken")
	row := q.conn.QueryRow(ctx, validateAccessTokenSQL, tokenHash, q.types.newAccessTokenScopeArrayInit(scopes))
	var item bool
	if err := row.Scan(&item); err != nil {
		return item, fmt.Errorf("query ValidateAccessToken: %w", err)
	}
	return item, nil
}

// ValidateAccessTokenBatch implements Querier.ValidateAccessTokenBatch.
func (q *DBQuerier) ValidateAccessTokenBatch(batch genericBatch, tokenHash string, scopes []AccessTokenScope) {
	batch.Queue(validateAccessTokenSQL, tokenHash, q.types.newAccessTokenScopeArrayInit(scopes))
}

// ValidateAccessTokenScan implements Querier.ValidateAccessTokenScan.
func (q *DBQuerier) ValidateAccessTokenScan(results pgx.BatchResults) (bool, error) {
	row := results.QueryRow()
	var item bool
	if err := row.Scan(&item); err != nil {
		return item, fmt.Errorf("scan ValidateAccessTokenBatch row: %w", err)
	}
	return item, nil
}

const getAccessTokenSQL = `SELECT id, name, scopes, issued_at, expires_at, revoked_at
FROM access_tokens
WHERE id = $1::uuid;`

type GetAccessTokenRow struct {
	ID        string             `json:"id"`
	Name      string             `json:"name"`
	Scopes    []AccessTokenScope `json:"scopes"`
	IssuedAt  pgtype.Timestamptz `json:"issued_at"`
	ExpiresAt pgtype.Timestamptz `json:"expires_at"`
	RevokedAt pgtype.Timestamptz `json:"revoked_at"`
}

// GetAccessToken implements Querier.GetAccessToken.
func (q *DBQuerier) GetAccessToken(ctx context.Context, id string) (GetAccessTokenRow, error) {
	ctx = context.WithValue(ctx, "pggen_query_name", "GetAccessToken")
	row := q.conn.QueryRow(ctx, getAccessTokenSQL, id)
	var item GetAccessTokenRow
	scopesArray := q.types.newAccessTokenScopeArray()
	if err := row.Scan(&item.ID, &item.Name, scopesArray, &item.IssuedAt, &item.ExpiresAt, &item.RevokedAt); err != nil {
		return item, fmt.Errorf("query GetAccessToken: %w", err)
	}
	if err := scopesArray.AssignTo(&item.Scopes); err != nil {
		return item, fmt.Errorf("assign GetAccessToken row: %w", err)
	}
	return item, nil
}

// GetAccessTokenBatch implements Querier.GetAccessTokenBatch.
func (q *DBQuerier) GetAccessTokenBatch(batch genericBatch, id string) {
	batch.Queue(getAccessTokenSQL, id)
}

// GetAccessTokenScan implements Querier.GetAccessTokenScan.
func (q *DBQuerier) GetAccessTokenScan(results pgx.BatchResults) (GetAccessTokenRow, error) {
	row := results.QueryRow()
	var item GetAccessTokenRow
	scopesArray := q.types.newAccessTokenScopeArray()
	if err := row.Scan(&item.ID, &item.Name, scopesArray, &item.IssuedAt, &item.ExpiresAt, &item.RevokedAt); err != nil {
		return item, fmt.Errorf("scan GetAccessTokenBatch row: %w", err)
	}
	if err := scopesArray.AssignTo(&item.Scopes); err != nil {
		return item, fmt.Errorf("assign GetAccessToken row: %w", err)
	}
	return item, nil
}

const listAccessTokensSQL = `SELECT id, name, scopes, issued_at, expires_at, revoked_at
FROM access_tokens
WHERE
  CASE WHEN $1
   THEN TRUE
   ELSE expires_at IS NULL OR expires_at > CURRENT_TIMESTAMP
  END AND
  CASE WHEN $2
   THEN TRUE
   ELSE revoked_at IS NULL OR revoked_at > CURRENT_TIMESTAMP
  END;`

type ListAccessTokensRow struct {
	ID        string             `json:"id"`
	Name      string             `json:"name"`
	Scopes    []AccessTokenScope `json:"scopes"`
	IssuedAt  pgtype.Timestamptz `json:"issued_at"`
	ExpiresAt pgtype.Timestamptz `json:"expires_at"`
	RevokedAt pgtype.Timestamptz `json:"revoked_at"`
}

// ListAccessTokens implements Querier.ListAccessTokens.
func (q *DBQuerier) ListAccessTokens(ctx context.Context, includeExpired bool, includeRevoked bool) ([]ListAccessTokensRow, error) {
	ctx = context.WithValue(ctx, "pggen_query_name", "ListAccessTokens")
	rows, err := q.conn.Query(ctx, listAccessTokensSQL, includeExpired, includeRevoked)
	if err != nil {
		return nil, fmt.Errorf("query ListAccessTokens: %w", err)
	}
	defer rows.Close()
	items := []ListAccessTokensRow{}
	scopesArray := q.types.newAccessTokenScopeArray()
	for rows.Next() {
		var item ListAccessTokensRow
		if err := rows.Scan(&item.ID, &item.Name, scopesArray, &item.IssuedAt, &item.ExpiresAt, &item.RevokedAt); err != nil {
			return nil, fmt.Errorf("scan ListAccessTokens row: %w", err)
		}
		if err := scopesArray.AssignTo(&item.Scopes); err != nil {
			return nil, fmt.Errorf("assign ListAccessTokens row: %w", err)
		}
		items = append(items, item)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("close ListAccessTokens rows: %w", err)
	}
	return items, err
}

// ListAccessTokensBatch implements Querier.ListAccessTokensBatch.
func (q *DBQuerier) ListAccessTokensBatch(batch genericBatch, includeExpired bool, includeRevoked bool) {
	batch.Queue(listAccessTokensSQL, includeExpired, includeRevoked)
}

// ListAccessTokensScan implements Querier.ListAccessTokensScan.
func (q *DBQuerier) ListAccessTokensScan(results pgx.BatchResults) ([]ListAccessTokensRow, error) {
	rows, err := results.Query()
	if err != nil {
		return nil, fmt.Errorf("query ListAccessTokensBatch: %w", err)
	}
	defer rows.Close()
	items := []ListAccessTokensRow{}
	scopesArray := q.types.newAccessTokenScopeArray()
	for rows.Next() {
		var item ListAccessTokensRow
		if err := rows.Scan(&item.ID, &item.Name, scopesArray, &item.IssuedAt, &item.ExpiresAt, &item.RevokedAt); err != nil {
			return nil, fmt.Errorf("scan ListAccessTokensBatch row: %w", err)
		}
		if err := scopesArray.AssignTo(&item.Scopes); err != nil {
			return nil, fmt.Errorf("assign ListAccessTokens row: %w", err)
		}
		items = append(items, item)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("close ListAccessTokensBatch rows: %w", err)
	}
	return items, err
}

const revokeAccessTokenSQL = `UPDATE access_tokens
SET revoked_at = CURRENT_TIMESTAMP
WHERE id = $1::uuid;`

// RevokeAccessToken implements Querier.RevokeAccessToken.
func (q *DBQuerier) RevokeAccessToken(ctx context.Context, id string) (pgconn.CommandTag, error) {
	ctx = context.WithValue(ctx, "pggen_query_name", "RevokeAccessToken")
	cmdTag, err := q.conn.Exec(ctx, revokeAccessTokenSQL, id)
	if err != nil {
		return cmdTag, fmt.Errorf("exec query RevokeAccessToken: %w", err)
	}
	return cmdTag, err
}

// RevokeAccessTokenBatch implements Querier.RevokeAccessTokenBatch.
func (q *DBQuerier) RevokeAccessTokenBatch(batch genericBatch, id string) {
	batch.Queue(revokeAccessTokenSQL, id)
}

// RevokeAccessTokenScan implements Querier.RevokeAccessTokenScan.
func (q *DBQuerier) RevokeAccessTokenScan(results pgx.BatchResults) (pgconn.CommandTag, error) {
	cmdTag, err := results.Exec()
	if err != nil {
		return cmdTag, fmt.Errorf("exec RevokeAccessTokenBatch: %w", err)
	}
	return cmdTag, err
}
