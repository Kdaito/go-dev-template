package service

import (
	"go-dev-sample/internal/domain/model"
	"go-dev-sample/internal/domain/repository"
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
