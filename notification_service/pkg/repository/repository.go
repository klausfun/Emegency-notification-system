package repository

import "github.com/jmoiron/sqlx"

type Notification interface {
}

type Repository struct {
	Notification
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{}
}
