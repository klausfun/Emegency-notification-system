package service

import "EmegencyNotificationSystem/profile_service/pkg/repository"

type Authorization interface {
}

type Profile interface{}

type Service struct {
	Authorization
	Profile
}

func NewService(repos *repository.Repository) *Service {
	return &Service{}
}
