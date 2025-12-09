package main

import (
	"fmt"
	"log"
	"os"
	"server/database"
	"server/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {

	database.ConnectDB()
	database.Migrate()

	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins:     "*",
		AllowMethods:     "GET,POST,PUT,DELETE,OPTIONS",
		AllowHeaders:     "Content-Type, Authorization, Accept, Origin, Cookie",
		AllowCredentials: true,
	}))

	routes.SetupRoutes(app)
	port := os.Getenv("PORT")
	if port == "" {
		port = "3030"
	}

	log.Fatal(app.Listen(":" + port))
	fmt.Println("Server running on port", port)

}
