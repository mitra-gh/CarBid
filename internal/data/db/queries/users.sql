-- name: CreateUser :one
INSERT INTO users (
    username,
    email,
    hashed_password,
    full_name
) VALUES (
    @username::varchar,
    @email::varchar,
    @hashed_password::varchar,
    @full_name::varchar
) RETURNING *;

-- name: GetUser :one
SELECT * FROM users 
WHERE id = @user_id::int LIMIT 1;

-- name: GetUserByUsername :one
SELECT * FROM users 
WHERE username = @username::varchar LIMIT 1;

-- name: GetUserByEmail :one
SELECT * FROM users 
WHERE email = @email::varchar LIMIT 1;

-- name: ListUsers :many
SELECT * FROM users
ORDER BY username
LIMIT @page_limit::int
OFFSET @page_offset::int;

-- name: UpdateUser :one
UPDATE users
SET
    username = COALESCE(@new_username::varchar, username),
    email = COALESCE(@new_email::varchar, email),
    hashed_password = COALESCE(@new_password::varchar, hashed_password),
    full_name = COALESCE(@new_full_name::varchar, full_name),
    is_active = COALESCE(@is_active::boolean, is_active),
    is_admin = COALESCE(@is_admin::boolean, is_admin)
WHERE id = @user_id::int
RETURNING *;

-- name: DeleteUser :exec
DELETE FROM users
WHERE id = @user_id::int;

-- name: CountUsers :one
SELECT COUNT(*) FROM users;

-- name: SearchUsers :many
SELECT * FROM users
WHERE 
    username LIKE '%' || @search_term::varchar || '%' OR
    email LIKE '%' || @search_term::varchar || '%' OR
    full_name LIKE '%' || @search_term::varchar || '%'
ORDER BY username
LIMIT @result_limit::int;

-- name: UpdateUserLastLogin :exec
UPDATE users
SET last_login_at = NOW()
WHERE id = @user_id::int;