package repository

type Authorization interface {
}

type Notification interface {
}

type Repository struct {
	Authorization
	Notification
}

func NewRepository() *Repository {
	return &Repository{}
}
