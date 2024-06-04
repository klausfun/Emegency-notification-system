package service

import "EmegencyNotificationSystem/pkg/repository"

type Authorization interface {
}

type Notification interface {
}

type Service struct {
	Authorization
	Notification
}

func NewService(repos *repository.Repository) *Service {
	return &Service{}
}
