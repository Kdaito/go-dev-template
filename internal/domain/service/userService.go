package service

import (
	"github.com/Kdaito/kinodokuna-be/internal/domain/model"
	"github.com/Kdaito/kinodokuna-be/internal/domain/repository"
)

type IUserService interface {
	GetUserList() ([]*model.User, error)
	GetUserByID(id int) (*model.User, error)
	CreateUser(name string, email string) (*model.User, error)
}

type UserService struct {
	repository repository.IUserRepository
}

func NewUserService(repository repository.IUserRepository) IUserService {
	return &UserService{repository: repository}
}

func (u *UserService) GetUserList() ([]*model.User, error) {
	return u.repository.GetUserList()
}

func (u *UserService) GetUserByID(id int) (*model.User, error) {
	return u.repository.GetUserByID(id)
}

func (u *UserService) CreateUser(name string, email string) (*model.User, error) {
	newUser := &model.User{
		Name:  name,
		Email: email,
	}

	return u.repository.CreateUser(newUser)
}
