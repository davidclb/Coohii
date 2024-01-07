package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/davidclb/Yoshu/account/handlers"
	"github.com/davidclb/Yoshu/account/repository"
	"github.com/davidclb/Yoshu/account/service"
	db "github.com/davidclb/Yoshu/db/sqlc"
	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5/pgxpool"
)
func main(){
	

////////////////////////////////////////// Data source//////////////////////////////////////////////////////////////////////////////////


	log.Printf("Initializing data sources\n")
	// load env variables - we could pass these in,
	// but this is sort of just a top-level (main package)
	// helper function, so I'll just read them in here


	log.Printf("Connecting to Postgresql\n")

	// urlExample := "postgres://username:password@localhost:5432/database_name"
	// Myurl:= postgres://root:root@postgresql:5432/yoshu?sslmode=disable
	//pool, err := pgxpool.Connect(context.Background(), os.Getenv("DATABASE_URL"))

	dsn := fmt.Sprintf("host=localhost user=root password=root dbname=yoshu port=5432 sslmode=disable")
	config2, err := pgxpool.ParseConfig(dsn)


	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to parse url: %v\n", err)
		os.Exit(1)
	}
	pool, err := pgxpool.NewWithConfig(context.Background(), config2)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	if pool.Ping(context.Background())!= nil {
		fmt.Fprintf(os.Stderr, "Unable to ping database: %v\n", pool.Ping(context.Background()))
		os.Exit(1)
	}
	log.Printf("Ping realisé")

	defer pool.Close()


	// SQLStore provides all functions to execute SQL queries and transactions

	// Verify database connection is working
	//if err := pool.Ping(context.Background()); err != nil {
	//	return fmt.Errorf("error connecting to db: %w", err)
	//}
	


////////////////////////////////////////// Dependency injection //////////////////////////////////////////////////////////////////////////////////

// Repository layer 

   repo := db.NewRepo(pool)
   userRepository := repository.NewUserRepository(repo)


// Service layer 

   userService := service.NewUserService(&service.USConfig{
		Repository:  userRepository,
	})

// Handler layer

   app := fiber.New()

	handlers.NewHandler(&handlers.Config{
		R: app,
		UserService: userService,
	})

////////////////////////////////////////// Server initialisation//////////////////////////////////////////////////////////////////////////////////

	log.Println("Starting server...")

    app.Listen(":3001")

	log.Println("Server started")


}