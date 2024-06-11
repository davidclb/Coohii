package handlers

import (
	"github.com/gofiber/fiber/v2"
)




func (h *Handler) Radicals(ctx *fiber.Ctx) error { 
	return ctx.JSON(fiber.Map{
		"hello": "it's me",
		})
 

}
