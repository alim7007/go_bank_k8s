-- name: CreateUser :one
INSERT INTO users (
  username,
  hashed_password,
  full_name,
  email
) VALUES (
  $1, $2, $3, $4
) RETURNING *;

-- name: GetUser :one
SELECT * FROM users
WHERE username = $1 LIMIT 1;

-- name: UpdateUser :one
Update users Set
hashed_password = Coalesce(sqlc.narg(hashed_password), hashed_password),
password_changed_at = COALESCE(sqlc.narg(password_changed_at), password_changed_at),
full_name = Coalesce(sqlc.narg(full_name), full_name),
email = Coalesce(sqlc.narg(email), email)
Where
  username = sqlc.arg(username)
RETURNING *;

-- ##############################################################################
-- -- name: UpdateUser :one
-- Update users Set
-- hashed_password = Case
-- -- @set_hashed_password == sqlc.arg(set_hashed_password)
--   When @set_hashed_password::boolean = True Then @hashed_password
--   else hashed_password
--   end,
-- full_name = Case
--   When @set_full_name::boolean = True Then @full_name
--   else full_name
--   end,
-- email = Case
--   When @set_email::boolean = True Then @email
--   else email
--   end
-- Where
--   username = @username
-- RETURNING *;
-- ##############################################################################