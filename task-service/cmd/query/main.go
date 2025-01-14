package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/rfashwall/go-microservice-workshop/pkg/db"
	"github.com/rfashwall/go-microservice-workshop/pkg/middleware"
	"github.com/rfashwall/go-microservice-workshop/pkg/utils"
	"github.com/rfashwall/task-service/internal/handlers"
	"github.com/rfashwall/task-service/internal/query"
)

func main() {
	shutdown := utils.InitTracer()
	defer shutdown()

	app := fiber.New()
	app.Use(logger.New())
	app.Use(middleware.TracingMiddleware("task-query-service"))

	// Connect to the database
	conn := db.MySqlConnect()
	defer conn.Close()

	// Initialize the repository
	taskQuery := query.NewMySQLTaskQuery(conn)

	// Initialize the handler
	taskHandler := handlers.NewTaskQueryHandler(taskQuery)

	// Set up routes
	taskHandler.SetupRoutes(app)

	log.Fatal(app.Listen(":3003"))
}
