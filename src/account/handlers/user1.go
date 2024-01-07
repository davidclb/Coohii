package handlers

import (
	"log"

	apperrors "github.com/davidclb/Yoshu/errors"
	"github.com/gofiber/fiber/v2"
)




func (h *Handler) User1(q *fiber.Ctx) error{

	
	u, err:= h.UserService.Get(q, 1)
	 if err != nil {
        log.Printf("Unable to find user: %v\n%v", 1, err)
        err := apperrors.NewNotFound("user", "uid")
        q.JSON(fiber.Map{
    			"erreur": err.Error(),
    			"status": err.Status(),
			  }, "application/problem+json")

		return err
			}

		q.JSON(fiber.Map{

			"user":u,
			  })
		return nil

}
