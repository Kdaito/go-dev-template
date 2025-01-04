package repository

import (
	"go-dev-sample/internal/domain/model"
)

type UserRepository interface {
	GetUserList() ([]*model.User, error)
	GetUserByID(id int) (*model.User, error)
	CreateUser(request *model.User) (*model.User, error)
}
