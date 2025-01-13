package repository

import "github.com/Kdaito/kinodokuna-be/internal/domain/model"

type IUserRepository interface {
	GetUserList() ([]*model.User, error)
	GetUserByID(id int) (*model.User, error)
	CreateUser(request *model.User) (*model.User, error)
}
