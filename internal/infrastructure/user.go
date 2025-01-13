package infrastructure

import (
	"database/sql"

	"github.com/Kdaito/kinodokuna-be/internal/domain/model"
	"github.com/Kdaito/kinodokuna-be/internal/domain/repository"
)

type User struct {
	db *sql.DB
}

func NewUser(db *sql.DB) repository.IUserRepository {
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

func (u *User) CreateUser(user *model.User) (*model.User, error) {
	stmt, err := u.db.Exec("INSERT INTO user (name, email) VALUES (?, ?)", user.Name, user.Email)
	if err != nil {
		return nil, err
	}

	id, err := stmt.LastInsertId()
	if err != nil {
		return nil, err
	}

	user.ID = int(id)
	return user, nil
}
