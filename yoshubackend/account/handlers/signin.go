package handlers

import (
	"fmt"
	"log"
	"strconv"
	"time"

	//postgres "github.com/davidclb/Yoshu/db/sqlc"
	apperrors "yoshubackend/errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)


const SecretKey="secret"

func (h *Handler) Signin(ctx *fiber.Ctx) error {

	
	var data map[string]string
	fmt.Println(data["password"])

	fmt.Println(data["username"])

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

	//password, _ := bcrypt.GenerateFromPassword([]byte(data["password"]), 14)
	
	//fmt.Println("password", password)
	//stringPassword:= string(password)

	//gotoByteApss := []byte(stringPassword)
	//fmt.Println("est-ce que c'est le même ?", gotoByteApss)

	/* user:= postgres.User{
		Username: data["name"],
		Email: data["email"],
		Password: "",

	} */

	user, err:= h.UserService.Signin(ctx, data["username"])


	// In the case we find don't find any user with this username 

	if err != nil {
		log.Printf("Pas d'utilisateur correspondant: %v\n", ctx)
		ctx.Status(fiber.StatusBadRequest)
		return ctx.JSON(fiber.Map{
			"message": "No user found",
		})
	}

	// In the case we find a user with this username 

	if err:= bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(data["password"])) ; err!=nil{
		ctx.Status(fiber.StatusBadRequest)
		return ctx.JSON(fiber.Map{
			"message": "Incorrect password",
		})
	}


	claims :=jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer: strconv.Itoa(int(user.ID)),
		ExpiresAt: time.Now().Add(time.Hour *24 ).Unix(), // 1 day
	})

	token, err:= claims.SignedString([]byte(SecretKey))

	if err != nil {
		ctx.Status(fiber.StatusInternalServerError)
		return ctx.JSON(fiber.Map{
			"message": "could not login",
		})
	}

	cookie := fiber.Cookie{
		Name: "jwt",
		Value: token,
		Expires: time.Now().Add(time.Hour *24 ), // 1 day
		HTTPOnly: true,
	}

	ctx.Cookie(&cookie)
	return ctx.JSON(fiber.Map{
			"message": "Succesfully login",
		})

} 

func (h *Handler) User(ctx *fiber.Ctx) error{

	// Find cookie and token of current user session

	cookie:= ctx.Cookies("jwt")

	token,err:= jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error){
		return []byte(SecretKey),nil
	})

	if err!= nil{
		ctx.Status(fiber.StatusUnauthorized)
		return ctx.JSON(fiber.Map{
			"message": "unauthentificated",
		})
	}

	claims:= token.Claims.(*jwt.StandardClaims)

	// Find the id of the issuer user

	id, err := strconv.Atoi(string(claims.Issuer))

	if err != nil {
        // ... handle error
        panic(err)
    }


	// Find the user

	user, err:= h.UserService.Get(ctx, int32(id))
	 if err != nil {
        log.Printf("Unable to find user: %v\n%v", id, err)
        err := apperrors.NewNotFound("user", "uid")
        ctx.JSON(fiber.Map{
    			"erreur": err.Error(),
    			"status": err.Status(),
			  }, "application/problem+json")

		return err
			}


	return ctx.JSON(user)

}