// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0
// source: users.sql

package sqlc

import (
	"context"
)

const countUsers = `-- name: CountUsers :one
SELECT COUNT(*) FROM users
`

func (q *Queries) CountUsers(ctx context.Context) (int64, error) {
	row := q.db.QueryRow(ctx, countUsers)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const createUser = `-- name: CreateUser :one
INSERT INTO users (
    username,
    email,
    hashed_password,
    full_name
) VALUES (
    $1::varchar,
    $2::varchar,
    $3::varchar,
    $4::varchar
) RETURNING id, username, first_name, last_name, mobile_number, email, password, enabled, created_at, updated_at
`

type CreateUserParams struct {
	Username       string `json:"username"`
	Email          string `json:"email"`
	HashedPassword string `json:"hashed_password"`
	FullName       string `json:"full_name"`
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (User, error) {
	row := q.db.QueryRow(ctx, createUser,
		arg.Username,
		arg.Email,
		arg.HashedPassword,
		arg.FullName,
	)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.FirstName,
		&i.LastName,
		&i.MobileNumber,
		&i.Email,
		&i.Password,
		&i.Enabled,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const deleteUser = `-- name: DeleteUser :exec
DELETE FROM users
WHERE id = $1::int
`

func (q *Queries) DeleteUser(ctx context.Context, userID int32) error {
	_, err := q.db.Exec(ctx, deleteUser, userID)
	return err
}

const getUser = `-- name: GetUser :one
SELECT id, username, first_name, last_name, mobile_number, email, password, enabled, created_at, updated_at FROM users 
WHERE id = $1::int LIMIT 1
`

func (q *Queries) GetUser(ctx context.Context, userID int32) (User, error) {
	row := q.db.QueryRow(ctx, getUser, userID)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.FirstName,
		&i.LastName,
		&i.MobileNumber,
		&i.Email,
		&i.Password,
		&i.Enabled,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getUserByEmail = `-- name: GetUserByEmail :one
SELECT id, username, first_name, last_name, mobile_number, email, password, enabled, created_at, updated_at FROM users 
WHERE email = $1::varchar LIMIT 1
`

func (q *Queries) GetUserByEmail(ctx context.Context, email string) (User, error) {
	row := q.db.QueryRow(ctx, getUserByEmail, email)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.FirstName,
		&i.LastName,
		&i.MobileNumber,
		&i.Email,
		&i.Password,
		&i.Enabled,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getUserByUsername = `-- name: GetUserByUsername :one
SELECT id, username, first_name, last_name, mobile_number, email, password, enabled, created_at, updated_at FROM users 
WHERE username = $1::varchar LIMIT 1
`

func (q *Queries) GetUserByUsername(ctx context.Context, username string) (User, error) {
	row := q.db.QueryRow(ctx, getUserByUsername, username)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.FirstName,
		&i.LastName,
		&i.MobileNumber,
		&i.Email,
		&i.Password,
		&i.Enabled,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const listUsers = `-- name: ListUsers :many
SELECT id, username, first_name, last_name, mobile_number, email, password, enabled, created_at, updated_at FROM users
ORDER BY username
LIMIT $2::int
OFFSET $1::int
`

type ListUsersParams struct {
	PageOffset int32 `json:"page_offset"`
	PageLimit  int32 `json:"page_limit"`
}

func (q *Queries) ListUsers(ctx context.Context, arg ListUsersParams) ([]User, error) {
	rows, err := q.db.Query(ctx, listUsers, arg.PageOffset, arg.PageLimit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []User
	for rows.Next() {
		var i User
		if err := rows.Scan(
			&i.ID,
			&i.Username,
			&i.FirstName,
			&i.LastName,
			&i.MobileNumber,
			&i.Email,
			&i.Password,
			&i.Enabled,
			&i.CreatedAt,
			&i.UpdatedAt,
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

const searchUsers = `-- name: SearchUsers :many
SELECT id, username, first_name, last_name, mobile_number, email, password, enabled, created_at, updated_at FROM users
WHERE 
    username LIKE '%' || $1::varchar || '%' OR
    email LIKE '%' || $1::varchar || '%' OR
    full_name LIKE '%' || $1::varchar || '%'
ORDER BY username
LIMIT $2::int
`

type SearchUsersParams struct {
	SearchTerm  string `json:"search_term"`
	ResultLimit int32  `json:"result_limit"`
}

func (q *Queries) SearchUsers(ctx context.Context, arg SearchUsersParams) ([]User, error) {
	rows, err := q.db.Query(ctx, searchUsers, arg.SearchTerm, arg.ResultLimit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []User
	for rows.Next() {
		var i User
		if err := rows.Scan(
			&i.ID,
			&i.Username,
			&i.FirstName,
			&i.LastName,
			&i.MobileNumber,
			&i.Email,
			&i.Password,
			&i.Enabled,
			&i.CreatedAt,
			&i.UpdatedAt,
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

const updateUser = `-- name: UpdateUser :one
UPDATE users
SET
    username = COALESCE($1::varchar, username),
    email = COALESCE($2::varchar, email),
    hashed_password = COALESCE($3::varchar, hashed_password),
    full_name = COALESCE($4::varchar, full_name),
    is_active = COALESCE($5::boolean, is_active),
    is_admin = COALESCE($6::boolean, is_admin)
WHERE id = $7::int
RETURNING id, username, first_name, last_name, mobile_number, email, password, enabled, created_at, updated_at
`

type UpdateUserParams struct {
	NewUsername string `json:"new_username"`
	NewEmail    string `json:"new_email"`
	NewPassword string `json:"new_password"`
	NewFullName string `json:"new_full_name"`
	IsActive    bool   `json:"is_active"`
	IsAdmin     bool   `json:"is_admin"`
	UserID      int32  `json:"user_id"`
}

func (q *Queries) UpdateUser(ctx context.Context, arg UpdateUserParams) (User, error) {
	row := q.db.QueryRow(ctx, updateUser,
		arg.NewUsername,
		arg.NewEmail,
		arg.NewPassword,
		arg.NewFullName,
		arg.IsActive,
		arg.IsAdmin,
		arg.UserID,
	)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.FirstName,
		&i.LastName,
		&i.MobileNumber,
		&i.Email,
		&i.Password,
		&i.Enabled,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const updateUserLastLogin = `-- name: UpdateUserLastLogin :exec
UPDATE users
SET last_login_at = NOW()
WHERE id = $1::int
`

func (q *Queries) UpdateUserLastLogin(ctx context.Context, userID int32) error {
	_, err := q.db.Exec(ctx, updateUserLastLogin, userID)
	return err
}
