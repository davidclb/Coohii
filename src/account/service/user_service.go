package service

import (
	"github.com/davidclb/Yoshu/account/repository"
	"github.com/davidclb/Yoshu/db/sqlc"
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

