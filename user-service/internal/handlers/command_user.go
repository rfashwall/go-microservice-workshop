package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rfashwall/user-service/internal/command"
	"github.com/rfashwall/user-service/internal/models"
)

type UserCommandHandler struct {
	UserCommand command.UserCommand
}

func NewUserCommandHandler(userCommand command.UserCommand) *UserCommandHandler {
	return &UserCommandHandler{UserCommand: userCommand}
}

func (h *UserCommandHandler) SetupRoutes(app *fiber.App) {
	app.Post("/users", h.createUser)
	app.Put("/users/:id", h.updateUser)
	app.Delete("/users/:id", h.deleteUser)
}

func (h *UserCommandHandler) createUser(c *fiber.Ctx) error {
	user := new(models.User)
	if err := c.BodyParser(user); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid input")
	}

	err := h.UserCommand.CreateUser(c.Context(), user)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	return c.Status(fiber.StatusCreated).JSON(user)
}

func (h *UserCommandHandler) updateUser(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid user ID")
	}

	user := new(models.User)
	if err := c.BodyParser(user); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid input")
	}
	user.ID = id

	err = h.UserCommand.UpdateUser(c.Context(), user)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	return c.JSON(user)
}

func (h *UserCommandHandler) deleteUser(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid user ID")
	}

	err = h.UserCommand.DeleteUser(c.Context(), id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	return c.SendStatus(fiber.StatusNoContent)
}
