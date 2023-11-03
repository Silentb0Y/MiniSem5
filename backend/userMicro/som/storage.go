package som

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type UserData struct {
	Username string
	Password string
}

type PostgresStore struct {
	db *sql.DB
}

func NewPostgresStore() (*PostgresStore, error) {
	connStr := "user=yash password=yash dbname=codeforprogress sslmode=disable"

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}
	return &PostgresStore{
		db: db,
	}, nil
}

type Storage interface {
	CreateUserInDB(string, string) error
	DeleteUserFromDB(*UserData) error
	CheckUservalidityAndGetUser(string, string) (bool, error)

	GetUserProfile(int)
}

func (s *PostgresStore) CreateUserInDB(username string, password string) error {
	query := `
	insert into "UserData" (username, password)
	values ( $1, $2)`

	_, err := s.db.Exec(query, username, password)

	if err != nil {
		return err
	}
	return nil
}

func (s *PostgresStore) DeleteUserFromDB(*UserData) error {
	return nil
}

func (s *PostgresStore) CheckUservalidityAndGetUser(username string, password string) (bool, error) {
	var userFound UserData
	query := `
	select * from "UserData"
	where username=$1 `

	err := s.db.QueryRow(query, username).Scan(&userFound.Username, &userFound.Password)
	if err != nil {
		return false, err
	}
	if password == userFound.Password {
		return true, nil
	}
	return false, fmt.Errorf("wrong password")

}

func (s *PostgresStore) GetUserProfile(id int) {

}
