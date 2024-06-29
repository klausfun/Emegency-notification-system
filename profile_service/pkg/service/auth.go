package service

import (
	"EmegencyNotificationSystem/profile_service/models"
	"EmegencyNotificationSystem/profile_service/pkg/repository"
	"crypto/sha1"
	"errors"
	"fmt"
	"os"
)

const (
	salt = "hbcwnjcjkn8038u8bvheff8vuhdih98"
)

type AuthService struct {
	repo repository.Authorization
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) CreateUser(user models.User, privateKey string) (int, error) {
	if privateKey != os.Getenv("PRIVATE_KEY") {
		return 0, errors.New("you do not have access to this system, provide another 'privateKey'")
	}

	user.Password = generatePasswordHash(user.Password)
	return s.repo.CreateUser(user)
}

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
