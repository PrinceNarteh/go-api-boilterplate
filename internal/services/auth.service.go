package services

import (
	"time"

	"github.com/PrinceNarteh/go-api/internal/models"
	"github.com/PrinceNarteh/go-api/internal/repositories"
)

var _ AuthService = (*authService)(nil)

type AuthService interface {
	Register(user *models.User) error
}

type authService struct {
	repo *repositories.Repositories
}

func (s *authService) Register(user *models.User) error {
	now := time.Now()

	user.CreatedAt = now
	user.UpdatedAt = now

	return s.repo.User.Create(user)
}
