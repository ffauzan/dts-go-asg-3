-- name: CreateUser :one
INSERT INTO "user" (username, email, password)
VALUES ( $1, $2, $3 )
RETURNING *;

-- name: GetUserByID :one
SELECT * FROM "user"
WHERE id = $1;

-- name: GetUserByUsername :one
SELECT * FROM "user"
WHERE username = $1;

-- name: UpdateUser :one
UPDATE "user" SET email = $1, password = $2, updated_at = CURRENT_TIMESTAMP
WHERE id = $3
RETURNING *;

-- name: CreateProfile :one
INSERT INTO "profile" (user_id, first_name, last_name)
VALUES ( $1, $2, $3 )
RETURNING *;

-- name: GetProfileByUserID :one
SELECT * FROM "profile"
WHERE user_id = $1;

-- name: UpdateProfile :one
UPDATE "profile" SET first_name = $1, last_name = $2, updated_at = CURRENT_TIMESTAMP
WHERE user_id = $3
RETURNING *;
