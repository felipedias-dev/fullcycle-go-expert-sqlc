// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.18.0
// source: query.sql

package db

import (
	"context"
)

const getCategory = `-- name: GetCategory :one
SELECT id, name, description FROM categories WHERE id = ?
`

func (q *Queries) GetCategory(ctx context.Context, id string) (Category, error) {
	row := q.db.QueryRowContext(ctx, getCategory, id)
	var i Category
	err := row.Scan(&i.ID, &i.Name, &i.Description)
	return i, err
}

const listCategories = `-- name: ListCategories :many
SELECT id, name, description FROM categories
`

func (q *Queries) ListCategories(ctx context.Context) ([]Category, error) {
	rows, err := q.db.QueryContext(ctx, listCategories)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Category
	for rows.Next() {
		var i Category
		if err := rows.Scan(&i.ID, &i.Name, &i.Description); err != nil {
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
