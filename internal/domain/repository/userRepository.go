package repository

import "go-dev-sample/internal/domain/model"

type UserRepository interface {
	GetUserByID(id int) (*model.User, error)
}
