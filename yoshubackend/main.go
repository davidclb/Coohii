package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v4/pgxpool"

	"yoshubackend/account/handlers"
	handlerStudy "yoshubackend/study/handlers"

	"yoshubackend/account/repository"
	"yoshubackend/account/service"
	repoStudy "yoshubackend/study/repository"
	serviceStudy "yoshubackend/study/service"

	"yoshubackend/db/sqlc"

	"github.com/gofiber/fiber/v2"
	/* 	"github.com/gofiber/fiber/v2/middleware/cors"
	 */)
	 
func main(){
	


////////////////////////////////////////// Data source//////////////////////////////////////////////////////////////////////////////////


 	log.Printf("Initializing data sources\n")
	// load env variables - we could pass these in,
	// but this is sort of just a top-level (main package)
	// helper function, so I'll just read them in here


	log.Printf("Connecting to Postgresql\n")

	ctx := context.Background()
	connStr := "postgresql://root:root@localhost:5432/yoshu?sslmode=disable"

	// Create a new connection pool
	pool, err := pgxpool.Connect(ctx, connStr)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer pool.Close()

	// Ping the database to ensure a connection can be established
	if err := pool.Ping(ctx); err != nil {
		fmt.Fprintf(os.Stderr, "Unable to ping the database: %v\n", err)
		os.Exit(1)
	}
	fmt.Println("Le ping est passé")

	// Use the connection pool with the generated sqlc queries
	queries := sqlc.New(pool)


////////////////////////////////////////// Dependency injection //////////////////////////////////////////////////////////////////////////////////

// Repository layer 

   userRepository := repository.NewUserRepository(&repository.UserRepo{
		Querie : queries,
	})

	studyRepository := repoStudy.NewStudyRepository(&repoStudy.StudyRepo{
		Querie : queries,
	})


// Service layer 

   userService := service.NewUserService(&service.USConfig{
		Repository:  userRepository,
	}) 

	studyService := serviceStudy.NewStudyService(&serviceStudy.StudyConfig{
		Repository:  studyRepository,
	}) 

// Handler layer

   app := fiber.New()

/*    app.Use(cors.New(cors.Config{
	AllowCredentials: true,
   })) */

 	handlers.NewHandler(&handlers.Config{
		R: app,
		UserService: userService,
	})

	handlerStudy.NewHandler(&handlerStudy.Config{
		R: app,
		StudyService: studyService,
	})
 
////////////////////////////////////////// Server initialisation//////////////////////////////////////////////////////////////////////////////////

	log.Println("Starting server...")

    app.Listen(":3001")

	log.Println("Server started")


}