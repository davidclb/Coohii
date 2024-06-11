package service

import (
	postgres "yoshubackend/db/sqlc"
	"github.com/gofiber/fiber/v2"
)


type UserService interface {
	Get(ctx *fiber.Ctx, uid int32) (*postgres.User, error)
	Signup(ctx *fiber.Ctx, u *postgres.User) ( error)
	Signin(ctx *fiber.Ctx, username string) (*postgres.User, error)
}
