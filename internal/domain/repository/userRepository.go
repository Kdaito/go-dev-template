package repository

import "go-dev-sample/internal/domain/model"

type UserRepository interface {
	GetUserList() ([]*model.User, error)
	GetUserByID(id int) (*model.User, error)
}
