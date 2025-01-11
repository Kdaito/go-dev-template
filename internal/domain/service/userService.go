package service

import (
	"github.com/Kdaito/go-dev-template/internal/domain/model"
	"github.com/Kdaito/go-dev-template/internal/domain/repository"
)

type UserService struct {
	repository repository.UserRepository
}

func NewUserService(repository repository.UserRepository) *UserService {
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
