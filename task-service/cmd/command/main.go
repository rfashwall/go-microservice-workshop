package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/rfashwall/go-microservice-workshop/pkg/db"
	"github.com/rfashwall/go-microservice-workshop/pkg/middleware"
	"github.com/rfashwall/go-microservice-workshop/pkg/utils"
	"github.com/rfashwall/task-service/internal/command"
	"github.com/rfashwall/task-service/internal/handlers"
)

func main() {
	shutdown := utils.InitTracer()
	defer shutdown()

	app := fiber.New()
	app.Use(logger.New())
	app.Use(middleware.TracingMiddleware("task-command-service"))

	// Connect to the database
	conn := db.MySqlConnect()
	defer conn.Close()

	// Initialize the repository
	taskCommand := command.NewMySQLTaskCommand(conn)

	// Initialize the handler
	taskHandler := handlers.NewTaskCommandHandler(taskCommand)

	// Set up routes
	taskHandler.SetupRoutes(app)

	log.Fatal(app.Listen(":3002"))
}
