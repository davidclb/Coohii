package service

import (
	postgres "github.com/davidclb/Yoshu/db/sqlc"
	"github.com/gofiber/fiber/v2"
)


type UserService interface {
	Get(ctx *fiber.Ctx, uid int32) (*postgres.User, error)
	Signup(ctx *fiber.Ctx, u *postgres.User) (*postgres.User, error)
	Signin(ctx *fiber.Ctx, username string) (*postgres.User, error)
}


type StudyService interface {
	LastReview(ctx *fiber.Ctx) (*postgres.Character, error)

}
