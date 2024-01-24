package dbrepo

import (
	"database/sql"

	"github.com/salmaqnsGH/unit-test/pkg/data"
)

type TestDBRepo struct{}

func (m *TestDBRepo) Connection() *sql.DB {
	return nil
}

func (m *TestDBRepo) AllUsers() ([]*data.User, error) {
	var users []*data.User

	return users, nil
}

func (m *TestDBRepo) GetUser(id int) (*data.User, error) {
	user := data.User{ID: 1}

	return &user, nil
}

func (m *TestDBRepo) GetUserByEmail(email string) (*data.User, error) {
	user := data.User{ID: 1}

	return &user, nil
}

func (m *TestDBRepo) UpdateUser(u data.User) error {
	return nil
}

func (m *TestDBRepo) DeleteUser(id int) error {
	return nil
}

func (m *TestDBRepo) InsertUser(user data.User) (int, error) {
	newID := 1

	return newID, nil
}

func (m *TestDBRepo) ResetPassword(id int, password string) error {
	return nil
}

func (m *TestDBRepo) InsertUserImage(i data.UserImage) (int, error) {
	newID := 1

	return newID, nil
}
