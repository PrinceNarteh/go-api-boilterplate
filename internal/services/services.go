// Package services
package services

import "github.com/PrinceNarteh/go-api/internal/repositories"

type Services struct {
	Auth AuthService
}

func NewServices(repo *repositories.Repositories) *Services {
	return &Services{
		Auth: &authService{repo: repo},
	}
}
