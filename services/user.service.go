package services

import (
	"database/sql"
	"fmt"

	"go.uber.org/zap"

	"github.com/michelm117/cycling-coach-lab/model"
)

type UserServicer interface {
	GetById(id int) (*model.User, error)
	GetByEmail(email string) (*model.User, error)
	GetAllUsers() ([]*model.User, error)
	AddUser(user model.User) (*model.User, error)
	DeleteUser(id int) error
	Count() (int, error)
}
type UserService struct {
	db     *sql.DB
	logger *zap.SugaredLogger
}

func NewUserServicer(db *sql.DB, logger *zap.SugaredLogger) UserServicer {
	return &UserService{
		db:     db,
		logger: logger,
	}
}

func (s *UserService) GetById(id int) (*model.User, error) {
	row := s.db.QueryRow("SELECT * FROM users WHERE users.id = $1", id)

	var user model.User
	err := row.Scan(
		&user.ID,
		&user.Email,
		&user.Firstname,
		&user.Lastname,
		&user.DateOfBirth,
		&user.PasswordHash,
		&user.Status,
		&user.Role,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (s *UserService) GetByEmail(email string) (*model.User, error) {
	row := s.db.QueryRow("SELECT * FROM users WHERE users.email = $1", email)

	var user model.User
	err := row.Scan(
		&user.ID,
		&user.Email,
		&user.Firstname,
		&user.Lastname,
		&user.DateOfBirth,
		&user.PasswordHash,
		&user.Status,
		&user.Role,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (s *UserService) GetAllUsers() ([]*model.User, error) {
	rows, err := s.db.Query("SELECT * FROM users")
	if err != nil {
		return nil, fmt.Errorf("error while trying to execute query: %s", err)
	}

	var users []*model.User
	for rows.Next() {
		var user model.User
		err := rows.Scan(
			&user.ID,
			&user.Email,
			&user.Firstname,
			&user.Lastname,
			&user.DateOfBirth,
			&user.PasswordHash,
			&user.Status,
			&user.Role,
			&user.CreatedAt,
			&user.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		users = append(users, &user)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	defer rows.Close()
	return users, nil
}

func (s *UserService) AddUser(user model.User) (*model.User, error) {
	_, err := s.db.Exec(
		"INSERT INTO users (email, firstname, lastname, date_of_birth, password_hash, status, role) VALUES ($1, $2, $3, $4, $5, $6, $7)",
		user.Email,
		user.Firstname,
		user.Lastname,
		user.DateOfBirth,
		user.PasswordHash,
		user.Status,
		user.Role,
	)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (s *UserService) DeleteUser(id int) error {
	_, err := s.db.Exec(
		"DELETE FROM users WHERE users.id = $1",
		id,
	)
	if err != nil {
		return fmt.Errorf("error while trying to execute query: %s", err)
	}
	return nil
}

func (s *UserService) Count() (int, error) {
	row := s.db.QueryRow("SELECT count(*) FROM users")
	var count int
	err := row.Scan(&count)
	if err != nil {
		return -1, fmt.Errorf("error while trying to execute query: %s", err)
	}
	return count, nil
}
