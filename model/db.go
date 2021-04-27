package model

import "database/sql"

type Repository struct {
	db *sql.DB
}

func CreateRepository(db *sql.DB) *Repository {
	return &Repository{
		db: db,
	}
}
