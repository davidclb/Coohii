package service

import (

	"yoshubackend/account/repository"
	"yoshubackend/db/sqlc"
	"github.com/gofiber/fiber/v2"
)

type userService struct {
	UserRepository repository.UserRepository
}



type USConfig struct{
	Repository repository.UserRepository
 }

func NewUserService(c *USConfig) UserService {
	return &userService{
		UserRepository:  c.Repository,
		}
}


func (US *userService ) Get(ctx *fiber.Ctx, uid int32) (*sqlc.User, error) {
	user, err := US.UserRepository.GetUserbyID( uid)
	return user, err
}

func (US *userService ) Signup(ctx *fiber.Ctx, user *sqlc.User) ( error) {


    err := US.UserRepository.CreateUser(user)


	return  err
}

func (US *userService ) Signin(ctx *fiber.Ctx, username string) (*sqlc.User, error) {
	user, err := US.UserRepository.GetUserbyUsername( username)

	if err != nil {

		return nil, err
	}


	return user, err
}

 