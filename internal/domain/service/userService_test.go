package service

import (
	"testing"

	"github.com/Kdaito/kinodokuna-be/internal/domain/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// IUserRepositoryのモック
type mockUserRepository struct {
	mock.Mock
}

func (m *mockUserRepository) GetUserList() ([]*model.User, error) {
	ret := m.Called()
	return ret.Get(0).([]*model.User), ret.Error(1)
}
func (m *mockUserRepository) GetUserByID(id int) (*model.User, error) {
	ret := m.Called(id)
	return ret.Get(0).(*model.User), ret.Error(1)
}
func (m *mockUserRepository) CreateUser(request *model.User) (*model.User, error) {
	ret := m.Called(request)
	return ret.Get(0).(*model.User), ret.Error(1)
}

func TestGetUserList(tt *testing.T) {
	tt.Run("正常系", func(t *testing.T) {
		// repository mock setup
		repository := new(mockUserRepository)
		mockedUsers := []*model.User{
			{
				ID:    1,
				Name:  "John Doe",
				Email: "John@example.com",
			},
			{
				ID:    2,
				Name:  "Kathy Smith",
				Email: "Kathy@example.com",
			},
		}
		repository.On("GetUserList").Return(mockedUsers, nil)

		// service setup
		service := NewUserService(repository)

		// assertion
		users, err := service.GetUserList()

		assert.NoError(t, err)
		assert.Equal(t, mockedUsers, users)
	})
}

func TestGetUserByID(tt *testing.T) {
	tt.Run("正常系", func(t *testing.T) {
		// repository mock setup
		repository := new(mockUserRepository)
		mockedUser := &model.User{
			ID:    1,
			Name:  "John Doe",
			Email: "John@example.com",
		}
		repository.On("GetUserByID", 1).Return(mockedUser, nil)

		// service setup
		service := NewUserService(repository)

		// assertion
		user, err := service.GetUserByID(1)

		assert.NoError(t, err)
		assert.Equal(t, mockedUser, user)
	})
}

func TestCreateUser(tt *testing.T) {
	tt.Run("正常系", func(t *testing.T) {
		var newUserName = "John Doe"
		var newUserEmail = "John@example.com"

		// repository mock setup
		repository := new(mockUserRepository)
		mockedCreateRequest := &model.User{
			Name:  newUserName,
			Email: newUserEmail,
		}
		mockedCreateResponse := &model.User{
			ID:    1,
			Name:  newUserName,
			Email: newUserEmail,
		}
		repository.On("CreateUser", mockedCreateRequest).Return(mockedCreateResponse, nil)

		// service setup
		service := NewUserService(repository)

		// assertion
		user, err := service.CreateUser(newUserName, newUserEmail)

		assert.NoError(t, err)
		assert.Equal(t, mockedCreateResponse, user)
	})
}
