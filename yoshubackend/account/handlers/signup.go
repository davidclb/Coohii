package handlers

import (
	"fmt"
	"log"

	postgres "yoshubackend/db/sqlc"
	apperrors "yoshubackend/errors"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)


func (h *Handler) Signup(ctx *fiber.Ctx) error {
	
	var data map[string]string

	err :=ctx.BodyParser(&data)
	
	if err != nil {
		log.Printf("Unable to extract user from request context for unknown reason: %v\n", ctx)
		err := apperrors.NewInternal()
		ctx.JSON(fiber.Map{
    			"erreur": err.Error(),
    			"status": err.Status(),
			  }, "application/problem+json")

		return err
	}

	password, _ := bcrypt.GenerateFromPassword([]byte(data["password"]), 14)
	
	fmt.Println("password", password)
	stringPassword:= string(password)

	//gotoByteApss := []byte(stringPassword)
	//fmt.Println("est-ce que c'est le même ?", gotoByteApss)

	user:= postgres.User{
		Username: data["username"],
		Email: data["email"],
		Password: stringPassword,

	}

	err2 := h.UserService.Signup(ctx, & user)
	if err2 != nil {
		log.Printf("Unable to extract user from request context for unknown reason: %v\n", ctx)
		err := apperrors.NewInternal()
		ctx.JSON(fiber.Map{
    			"erreur": err.Error(),
    			"status": err.Status(),
			  }, "application/problem+json")

		return err
	}


	//log.Printf("User created: %v\n", ctx)
	return ctx.JSON(user)

} 