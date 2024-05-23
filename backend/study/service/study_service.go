package service

import (
	"github.com/davidclb/Yoshu/account/repository"
	"github.com/davidclb/Yoshu/db/sqlc"
	"github.com/gofiber/fiber/v2"
)

type studyService struct {
	StudyRepository repository.StudyRepository
}



type StudyConfig struct{
	Repository repository.UserRepository
 }

func NewUserService(c *StudyConfig) StudyService {
	return &studyService{
		StudyRepository:  c.Repository,
		}
}


func (US *studyService ) GetRadicals(ctx *fiber.Ctx, uid int32) (*sqlc.User, error) {
	user, err := US.StudyRepository.GetUserbyID( uid)
	return user, err
}

func (US *studyService ) GetHanzis(ctx *fiber.Ctx, user *sqlc.User) (*sqlc.User, error) {
	user, err := US.StudyRepository.CreateUser(user)

	return user, err
}

func (US *studyService ) GetWords(ctx *fiber.Ctx, username string) (*sqlc.User, error) {
	user, err := US.UserRepository.GetUserbyUsername( username)

	if err != nil {

		return nil, err
	}


	return user, err
}

func (US *studyService ) GetSentence(ctx *fiber.Ctx, user *sqlc.User) (*sqlc.User, error) {
	user, err := US.StudyRepository.CreateUser(user)

	return user, err
}

 