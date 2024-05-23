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

	c.R.Get("/api/explore/radicals", h.Radicals)
	c.R.Get("/api/explore/sentences", h.Sentences)
	c.R.Get("/api/explore/hanzi", h.Hanzi)
	c.R.Get("/api/explore/words", h.Words)


/* 	c.R.Get("/api/study/session", h.Radicals)
	c.R.Get("/api/study/session/:id/cards", h.Sentences)
	c.R.Get("/api/study/session/:id/answer", h.Hanzi)
	c.R.Get("/api/study/session/:id/complete", h.Words)
	

	c.R.Get("/api/stats/:user_id", h.Radicals) */


	 
 }



