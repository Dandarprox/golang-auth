package interfaces

import "github.com/Dandarprox/golang-auth/entities"

type JwtI interface {
	GenerateToken(user *entities.User) string
	ValidateToken(token string) (bool, error)
	GetTokenUser(token string) (*entities.User, error)
}
