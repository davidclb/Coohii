package service

import (
	postgres "yoshubackend/db/sqlc"
	"github.com/gofiber/fiber/v2"
)


type StudyService interface {

 //	GetWords(ctx *fiber.Ctx) (*[]postgres.Word, error)
 // GetSentences(ctx *fiber.Ctx) (*[]postgres.Sentences, error)
 	GetHanzi(ctx *fiber.Ctx, filter string, limit int32, last_id int32 ) ([]*postgres.Character, error)
//	GetRadicals(ctx *fiber.Ctx) (*[]postgres.Radical, error)
//	GetReviews(ctx *fiber.Ctx) (*[]postgres.Review, error)
//	CreateSession(ctx *fiber.Ctx)(*postgres.Character, error)
//	CompleteSession(ctx *fiber.Ctx)(*postgres.Character, error)
	
}
