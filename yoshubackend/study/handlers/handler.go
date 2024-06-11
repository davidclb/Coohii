package handlers

import (
	"yoshubackend/study/service"

	"github.com/gofiber/fiber/v2"
)

// This struct will probably hold the service that have to be used by the handler but empty for now
type Handler struct {

	StudyService service.StudyService

}

type Config struct{
	R *fiber.App
	StudyService service.StudyService
 }

 func NewHandler(c *Config){


	h := Handler{
		StudyService: c.StudyService,
	}

	c.R.Get("/api/explore/radicals", h.Radicals)
	c.R.Get("/api/explore/sentences", h.Sentences)
	c.R.Get("/api/explore/hanzi", h.Hanzi)
	c.R.Get("/api/explore/words", h.Words)
	c.R.Get("/api/study/reviews", h.Reviews)

	 
 }



