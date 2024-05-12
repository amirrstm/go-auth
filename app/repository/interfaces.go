package repository

import (
	"github.com/amirrstm/go-auth/app/dto"
)

type UserRepository interface {
	Create(b *dto.CreateUser) error
}
