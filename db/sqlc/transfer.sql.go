// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.13.0
// source: transfer.sql

package db

import (
	"context"
	"database/sql"
	"time"
)

const createTransfer = `-- name: CreateTransfer :execresult
INSERT INTO transfers (
    from_account_id,
    to_account_id,
    amount,
    created_at
) VALUES (
    ?, ?, ?, ?
)
`

type CreateTransferParams struct {
	FromAccountID int64     `json:"from_account_id"`
	ToAccountID   int64     `json:"to_account_id"`
	Amount        int64     `json:"amount"`
	CreatedAt     time.Time `json:"created_at"`
}

func (q *Queries) CreateTransfer(ctx context.Context, arg CreateTransferParams) (sql.Result, error) {
	return q.db.ExecContext(ctx, createTransfer,
		arg.FromAccountID,
		arg.ToAccountID,
		arg.Amount,
		arg.CreatedAt,
	)
}

const getTransfer = `-- name: GetTransfer :one
SELECT id, from_account_id, to_account_id, amount, created_at FROM transfers
WHERE id = ? LIMIT 1
`

func (q *Queries) GetTransfer(ctx context.Context, id int64) (Transfer, error) {
	row := q.db.QueryRowContext(ctx, getTransfer, id)
	var i Transfer
	err := row.Scan(
		&i.ID,
		&i.FromAccountID,
		&i.ToAccountID,
		&i.Amount,
		&i.CreatedAt,
	)
	return i, err
}

const listTransfer = `-- name: ListTransfer :many
SELECT id, from_account_id, to_account_id, amount, created_at FROM transfers
WHERE from_account_id = ? OR
      to_account_id = ?
ORDER BY id LIMIT ?,?
`

type ListTransferParams struct {
	FromAccountID int64 `json:"from_account_id"`
	ToAccountID   int64 `json:"to_account_id"`
	Offset        int32 `json:"offset"`
	Limit         int32 `json:"limit"`
}

func (q *Queries) ListTransfer(ctx context.Context, arg ListTransferParams) ([]Transfer, error) {
	rows, err := q.db.QueryContext(ctx, listTransfer,
		arg.FromAccountID,
		arg.ToAccountID,
		arg.Offset,
		arg.Limit,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Transfer{}
	for rows.Next() {
		var i Transfer
		if err := rows.Scan(
			&i.ID,
			&i.FromAccountID,
			&i.ToAccountID,
			&i.Amount,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
