package model

import (
	"database/sql"
)

type DBTX interface {
	Exec(string, ...interface{}) (sql.Result, error)
	Prepare(string) (*sql.Stmt, error)
	Query(string, ...interface{}) (*sql.Rows, error)
	QueryRow(string, ...interface{}) *sql.Row
}

type Repository struct {
	db DBTX
}

func CreateRepository(tx *sql.DB) *Repository {
	return &Repository{
		db: tx,
	}
}

func CreateRepositoryWithTx(tx *sql.Tx) *Repository {
	return &Repository{
		db: tx,
	}
}

type SQLRepository struct {
	db *sql.DB
}

func NewSQLRepository(db *sql.DB) *SQLRepository {
	return &SQLRepository{
		db: db,
	}
}