package repository

const queryCreateUser = `
	INSERT INTO users (id, username, password_hash, created_at, updated_at)
	VALUES ($1, $2, $3, now(), now())
	RETURNING id, username, password_hash, created_at, updated_at, deleted_at
`

const queryGetUserByUsername = `
	SELECT id, username, password_hash, created_at, updated_at, deleted_at
	FROM users
	WHERE username = $1 AND deleted_at IS NULL
`

const queryUpdatePassword = `
	UPDATE users
	SET password_hash = $1, updated_at = now()
	WHERE id = $2 AND deleted_at IS NULL
`

const queryDeleteUser = `
	UPDATE users
	SET deleted_at = now()
	WHERE id = $1 AND deleted_at IS NULL
`
