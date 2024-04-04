package handlers

import (
	"github.com/davidclb/Yoshu/model"
	"github.com/gofiber/fiber/v2"
)

// This struct will probably hold the service that have to be used by the handler but empty for now
type Handler struct {

	UserService model.UserService

}

type Config struct{
	R *fiber.App
	UserService model.UserService
 }

 func NewHandler(c *Config){


	h := Handler{
		UserService: c.UserService,
	}

	c.R.Get("/api/study/accueil", h.Accueil)
	c.R.Get("/api/study/kanji", h.Kanji)
	c.R.Get("/api/study/stories", h.Stories)
	 
 }


func (h *Handler) Kanji(ctx *fiber.Ctx) error { 
	return ctx.JSON(fiber.Map{
		"hello": "it's me",
		})
 

}


func (h *Handler) Stories(ctx *fiber.Ctx) error { 
	return ctx.JSON(fiber.Map{
		"hello": "it's me",
		})


}

