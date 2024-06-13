package handlers

import (
	"log"
	"strconv"
	apperrors "yoshubackend/errors"

	"github.com/gofiber/fiber/v2"
)




func (h *Handler) Hanzi(ctx *fiber.Ctx) error { 

	limit, err := strconv.Atoi(ctx.Query("limit", "20"))
    if err != nil || limit < 1 {
        limit = 20
    }

    lastID, err := strconv.Atoi(ctx.Query("last_id", "1"))
    if err != nil {
        lastID = 0
    }

    filter := ctx.Query("filter", "")


    cards, err:= h.StudyService.GetHanzi(ctx, filter, int32(limit), int32(lastID) )

    if err != nil {
        log.Printf("Unable get Hanzi")
        err := apperrors.NewNotFound("user", "uid")
        ctx.JSON(fiber.Map{
    			"erreur": err.Error(),
    			"status": err.Status(),
			  }, "application/problem+json")

		return err
			}

    return ctx.JSON(cards) 

	

}
