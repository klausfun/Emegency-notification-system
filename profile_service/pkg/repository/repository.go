package repository

import "github.com/jmoiron/sqlx"

type Authorization interface {
}

type Profile interface{}

type Repository struct {
	Authorization
	Profile
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{}
}
