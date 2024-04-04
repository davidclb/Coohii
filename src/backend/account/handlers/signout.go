package handlers

import (
	"time"

	"github.com/gofiber/fiber/v2"
)



func (h *Handler) Signout(ctx *fiber.Ctx) error{

	// Find cookie and token of current user session

	cookie:= fiber.Cookie{
		Name: "jwt",
		Value: "",
		Expires: time.Now().Add(-time.Hour),
	}

	ctx.Cookie(&cookie)


	return ctx.JSON(fiber.Map{
			"message": "Successfully logout",
		})

}