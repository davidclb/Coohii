package service

import (
	postgres "github.com/davidclb/Yoshu/db/sqlc"
	"github.com/gofiber/fiber/v2"
)


type StudyService interface {

	GetWords(ctx *fiber.Ctx) (*[]postgres.Words, error)
	GetSentences(ctx *fiber.Ctx) (*[]postgres.Sentences, error)
	GetHanzis(ctx *fiber.Ctx) (*[]postgres.Hanzis, error)
	GetRadicals(ctx *fiber.Ctx) (*[]postgres.Radicals, error)
	GetReviews(ctx *fiber.Ctx) (*[]postgres.Reviews, error)
	CreateSession(ctx *fiber.Ctx)(*postgres.Character, error)
	CompleteSession(ctx *fiber.Ctx)(*postgres.Character, error)
	
}
