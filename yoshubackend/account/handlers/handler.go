package handlers

import (
	"yoshubackend/account/service"
	"github.com/gofiber/fiber/v2"
)

// This struct will probably hold the service that have to be used by the handler but empty for now
type Handler struct {

	UserService service.UserService

}

type Config struct{
	R *fiber.App
	UserService service.UserService
 }

 func NewHandler(c *Config){


	h := Handler{
		UserService: c.UserService,
	}

	c.R.Get("/api/account/me", h.Me)
	c.R.Post("/signup", h.Signup)
	c.R.Post("/signin", h.Signin)
	c.R.Post("/signout", h.Signout)
	c.R.Get("/user", h.User)
	 
 }
	


 func (h *Handler) Details(ctx *fiber.Ctx) error { 
		 return ctx.JSON(fiber.Map{
   			 "hello": "it's me",
  			})
	}



