package services

import "github.com/PrinceNarteh/go-api/internal/models"

var _ AuthService = (*authService)(nil)

type AuthService interface {
	Register(user *models.User) error
}

type authService struct{}

func (s *authService) Register(user *models.User) error {
	return nil
}
