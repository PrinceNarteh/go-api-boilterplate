package utils

import (
	"time"

	"github.com/PrinceNarteh/go-api/config"
	"github.com/PrinceNarteh/go-api/internal/models"
	"github.com/golang-jwt/jwt/v5"
)

type AccessClaims struct {
	UserId string
	Email  string
	Roles  []string
	jwt.RegisteredClaims
}

type RefreshClaims struct {
	UserId string
	jwt.RegisteredClaims
}

func GenerateTokens(user *models.User) (*models.Tokens, error) {
	now := time.Now()

	accessClaims := AccessClaims{
		UserId: user.ID.Hex(),
		Email:  user.Email,
		Roles:  user.Roles,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "go-api",
			Subject:   user.ID.Hex(),
			IssuedAt:  jwt.NewNumericDate(now),
			NotBefore: jwt.NewNumericDate(now),
			ExpiresAt: jwt.NewNumericDate(now.Add(config.Env.JWT.AccessTokenTTL)),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, accessClaims)
	accessToken, err := token.SignedString(config.Env.JWT.AccessTokenSecret)
	if err != nil {
		return nil, err
	}

	refreshClaims := RefreshClaims{
		UserId: user.ID.Hex(),
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "go-api",
			Subject:   user.ID.Hex(),
			IssuedAt:  jwt.NewNumericDate(now),
			NotBefore: jwt.NewNumericDate(now),
			ExpiresAt: jwt.NewNumericDate(now.Add(config.Env.JWT.RefreshTokenTTL)),
		},
	}
	token = jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims)
	refreshToken, err := token.SignedString(config.Env.JWT.RefreshTokenSecret)
	if err != nil {
		return nil, err
	}

	return &models.Tokens{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}
