package repository

import (
	"context"

	postgres "github.com/davidclb/Yoshu/db/sqlc"
)


type UserRepository interface {
	CreateUser(User *postgres.User) (*postgres.User, error)
	ListUsers() (*[]postgres.User, error)
	GetUserbyID(id int32) (*postgres.User, error)
	DeleteUser(User *postgres.User) error
}

type repository struct {
	repo *postgres.Repo
}

func NewUserRepository(repo *postgres.Repo) UserRepository {
	return &repository{repo: repo}
}

func (r *repository) CreateUser(User *postgres.User) (*postgres.User, error) {
	userparam := postgres.CreateUserParams{
		User.Username,
		User.Email,
		User.Password,
	}
	result, err := r.repo.CreateUser(context.Background(), userparam )
	if err != nil {
		return nil, err
	}
	return &result, nil
}


func (r *repository) ListUsers() (*[]postgres.User, error) {
	result, err := r.repo.ListUsers(context.Background())
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (r *repository) GetUserbyID(id int32) (*postgres.User, error) {
	result, err := r.repo.GetUserbyID(context.Background(), id)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (r *repository) DeleteUser(User *postgres.User) error {
	err := r.repo.DeleteUser(context.Background(), User.ID)
	if err != nil {
		return err
	}
	return nil
}