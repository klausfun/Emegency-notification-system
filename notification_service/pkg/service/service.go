package service

import "EmegencyNotificationSystem/notification_service/pkg/repository"

type Notification interface {
}

type Service struct {
	Notification
}

func NewService(repos *repository.Repository) *Service {
	return &Service{}
}
