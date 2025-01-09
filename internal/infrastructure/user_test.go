package infrastructure

import (
	"go-dev-sample/internal/domain/model"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestGetUserList(t *testing.T) {
	t.Run("正常系", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()

		rows := sqlmock.NewRows([]string{"id", "name", "email"}).
			AddRow("1", "John Doe", "John@example.com").
			AddRow("2", "Kathy Smith", "Kathy@example.com")

		mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM user")).WillReturnRows(rows)

		repository := NewUser(db)
		users, err := repository.GetUserList()

		assert.NoError(t, err)
		assert.Len(t, users, 2)
		assert.Equal(t, "John Doe", users[0].Name)
		assert.Equal(t, "Kathy Smith", users[1].Name)
	})

	t.Run("異常系 - Queryメソッドでエラー", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()

		mock.ExpectQuery("SELECT * FROM user").WillReturnError(assert.AnError)

		repository := NewUser(db)
		users, err := repository.GetUserList()

		assert.Error(t, err)
		assert.Nil(t, users)
	})
}

func TestGetUserByID(t *testing.T) {
	t.Run("正常系", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()

		rows := sqlmock.NewRows([]string{"id", "name", "email"}).
			AddRow("1", "John Doe", "John@example.com")

		mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM user WHERE id = ?")).WithArgs(1).WillReturnRows(rows)

		repository := NewUser(db)
		user, err := repository.GetUserByID(1)

		assert.NoError(t, err)
		assert.Equal(t, "John Doe", user.Name)
	})

	t.Run("異常系 - QueryRowメソッドでエラー", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()

		rows := sqlmock.NewRows([]string{"id", "name", "email"}).
			AddRow("1", "John Doe", "John@example.com").RowError(0, assert.AnError)

		mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM user WHERE id = ?")).WithArgs(1).WillReturnRows(rows)

		repository := NewUser(db)
		user, err := repository.GetUserByID(1)

		assert.Error(t, err)
		assert.Nil(t, user)
	})
}

func TestCreateUser(t *testing.T) {
	t.Run("正常系", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()

		mock.ExpectExec(regexp.QuoteMeta("INSERT INTO user (name, email) VALUES (?, ?)")).WithArgs("John Doe", "John@example.com").WillReturnResult(sqlmock.NewResult(1, 1))

		repository := NewUser(db)
		user, err := repository.CreateUser(&model.User{Name: "John Doe", Email: "John@example.com"})

		assert.NoError(t, err)
		assert.NotNil(t, user)
		assert.Equal(t, 1, user.ID)
	})

	t.Run("異常系 - Execメソッドでエラー", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()

		mock.ExpectExec(regexp.QuoteMeta("INSERT INTO user (name, email) VALUES (?, ?)")).WithArgs("John Doe", "John@example.com").WillReturnError(assert.AnError)

		repository := NewUser(db)
		user, err := repository.CreateUser(&model.User{Name: "John Doe", Email: "John@example.com"})

		assert.Error(t, err)
		assert.Nil(t, user)
	})

	t.Run("異常系 - LastInsertIdメソッドでエラー", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		assert.NoError(t, err)
		defer db.Close()

		mock.ExpectExec(regexp.QuoteMeta("INSERT INTO user (name, email) VALUES (?, ?)")).WithArgs("John Doe", "John@example.com").WillReturnResult(sqlmock.NewErrorResult(assert.AnError))

		repository := NewUser(db)
		user, err := repository.CreateUser(&model.User{Name: "John Doe", Email: "John@example.com"})

		assert.Error(t, err)
		assert.Nil(t, user)
	})
}
