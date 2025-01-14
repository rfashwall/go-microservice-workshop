// File: /home/rabieh/projects/go-microservice-workshop/user-service/internal/handlers/user_handler.go

package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rfashwall/user-service/internal/query"
)

type UserQueryHandler struct {
	UserQuery query.UserQuery
}

func NewUserQueryHandler(userQuery query.UserQuery) *UserQueryHandler {
	return &UserQueryHandler{UserQuery: userQuery}
}

func (h *UserQueryHandler) SetupRoutes(app *fiber.App) {
	app.Get("/users/:id", h.getUser)
	app.Get("/users", h.listUsers)
}

func (h *UserQueryHandler) getUser(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid user ID")
	}

	user, err := h.UserQuery.GetUserByID(c.Context(), id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	return c.JSON(user)
}

func (h *UserQueryHandler) listUsers(c *fiber.Ctx) error {
	users, err := h.UserQuery.ListUsers(c.Context())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	return c.JSON(users)
}
