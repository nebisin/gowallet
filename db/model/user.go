package model

type CreateUserParams struct {
	Username       string `json:"username"`
	HashedPassword string `json:"password"`
	FullName       string `json:"full_name"`
	Email          string `json:"email"`
}

const createUser = `INSERT INTO users (
	username,
	hashed_password,
	full_name,
	email
) VALUES (
	$1, $2, $3, $4
) RETURNING username, hashed_password, full_name, email, created_at`

func (r *Repository) CreateUser(args CreateUserParams) (User, error) {
	row := r.db.QueryRow(createUser, args.Username, args.HashedPassword, args.FullName, args.Email)

	var i User
	err := row.Scan(
		&i.Username,
		&i.HashedPassword,
		&i.FullName,
		&i.Email,
		&i.CreatedAt,
	)
	return i, err
}

const getUser = `SELECT username, hashed_password, full_name, email, created_at
FROM users WHERE username = $1 LIMIT 1`

func (r *Repository) GetUser(username string) (User, error) {
	row := r.db.QueryRow(getUser, username)

	var i User
	err := row.Scan(
		&i.Username,
		&i.HashedPassword,
		&i.FullName,
		&i.Email,
		&i.CreatedAt,
	)
	return i, err
}
