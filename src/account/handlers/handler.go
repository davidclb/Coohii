package handlers

import (
	"github.com/davidclb/Yoshu/account/service"
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
	c.R.Get("/api/account/user1", h.User1)
	//g.Post("/signup", h.Signup)
	//g.Post("/signin", h.Signin)
	//g.Post("/signout", h.Signout)
	//g.Post("/token", h.Token)
	//g.Put("/details", h.Details)
	 
 }
	



 func (h *Handler) Signout(ctx *fiber.Ctx) error { 
		 return ctx.JSON(fiber.Map{
   			 "hello": "it's me",
  			})


	}

 func (h *Handler) Signin(ctx *fiber.Ctx) error { 
		 return ctx.JSON(fiber.Map{
   			 "hello": "it's me",
  			})

	}

 func (h *Handler) Token(ctx *fiber.Ctx) error { 
		 return ctx.JSON(fiber.Map{
   			 "hello": "it's me",
  			})
	}

 func (h *Handler) Details(ctx *fiber.Ctx) error { 
		 return ctx.JSON(fiber.Map{
   			 "hello": "it's me",
  			})
	}



