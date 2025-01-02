package infrastructure

import (
	"database/sql"
	"go-dev-sample/internal/domain/model"
	"go-dev-sample/internal/domain/repository"
)

type User struct {
	db *sql.DB
}

func NewUser(db *sql.DB) repository.UserRepository {
	return &User{db: db}
}

func (u *User) GetUserByID(id int) (*model.User, error) {
	row := u.db.QueryRow("SELECT * FROM user WHERE id = ?", id)
	user := &model.User{}
	if err := row.Scan(&user.ID, &user.Name, &user.Email); err != nil {
		return nil, err
	}
	return user, nil
}
