package repository

import (
	"context"
	"log"

	postgres "github.com/davidclb/Yoshu/db/sqlc"
)


type StudyRepository interface {
	CreateUser(User *postgres.User) (*postgres.User, error)
	ListUsers() (*[]postgres.User, error)
	GetUserbyID(id int32) (*postgres.User, error)
	GetUserbyUsername(username string) (*postgres.User, error)
	DeleteUser(User *postgres.User) error
}

type UserRepo struct {
	Querie *postgres.Queries
}

func NewUserRepository(r *UserRepo) UserRepository {
	return &UserRepo{Querie: r.Querie}
}

func (r *UserRepo) CreateUser(User *postgres.User) (*postgres.User, error) {
	userparam := postgres.CreateUserParams{
		Username: User.Username,
		Email: User.Email,
		Password: User.Password,
	}
	
	result, err := r.Querie.CreateUser(context.Background(), userparam )
	if err != nil {
		return nil, err
	}
	log.Printf("User created in db")
	return &result, nil
}


func (r *UserRepo) ListUsers() (*[]postgres.User, error) {
	result, err := r.Querie.ListUsers(context.Background())
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (r *UserRepo) GetUserbyID(id int32) (*postgres.User, error) {
	result, err := r.Querie.GetUserbyID(context.Background(), id)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (r *UserRepo) GetUserbyUsername(username string) (*postgres.User, error) {
	result, err := r.Querie.GetUserbyUsername(context.Background(), username)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (r *UserRepo) DeleteUser(User *postgres.User) error {
	err := r.Querie.DeleteUser(context.Background(), User.ID)
	if err != nil {
		return err
	}
	return nil
}  