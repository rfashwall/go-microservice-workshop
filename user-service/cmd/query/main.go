package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/rfashwall/go-microservice-workshop/pkg/db"
	"github.com/rfashwall/go-microservice-workshop/pkg/middleware"
	"github.com/rfashwall/go-microservice-workshop/pkg/utils"
	"github.com/rfashwall/user-service/internal/handlers"
	"github.com/rfashwall/user-service/internal/query"
)

func main() {
	shutdown := utils.InitTracer()
	defer shutdown()

	app := fiber.New()
	app.Use(logger.New())
	app.Use(middleware.TracingMiddleware("user-query-service"))

	// Connect to the database
	conn := db.MySqlConnect()
	defer conn.Close()

	db.SeedData(conn)

	// Initialize the repository
	userQuery := query.NewMySQLUserQuery(conn)

	// Initialize the handler
	userHandler := handlers.NewUserQueryHandler(userQuery)

	// Set up routes
	userHandler.SetupRoutes(app)

	log.Fatal(app.Listen(":3000"))
}