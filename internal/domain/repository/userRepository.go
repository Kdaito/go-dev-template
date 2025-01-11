package repository

import "github.com/Kdaito/go-dev-template/internal/domain/model"

type UserRepository interface {
	GetUserList() ([]*model.User, error)
	GetUserByID(id int) (*model.User, error)
	CreateUser(request *model.User) (*model.User, error)
}
