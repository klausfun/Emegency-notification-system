package service

import (
	"EmegencyNotificationSystem/profile_service/models"
	"EmegencyNotificationSystem/profile_service/pkg/repository"
)

type Authorization interface {
	CreateUser(user models.User, privateKey string) (int, error)
}

type Profile interface{}

type Service struct {
	Authorization
	Profile
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
	}
}
