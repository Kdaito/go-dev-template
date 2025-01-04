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

func (u *User) GetUserList() ([]*model.User, error) {
	rows, err := u.db.Query("SELECT * FROM user")
	if err != nil {
		return nil, err
	}

	users := []*model.User{}

	for rows.Next() {
		user := &model.User{}
		if err := rows.Scan(&user.ID, &user.Name, &user.Email); err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

func (u *User) GetUserByID(id int) (*model.User, error) {
	row := u.db.QueryRow("SELECT * FROM user WHERE id = ?", id)
	user := &model.User{}
	if err := row.Scan(&user.ID, &user.Name, &user.Email); err != nil {
		return nil, err
	}
	return user, nil
}
