package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rfashwall/task-service/internal/command"
	"github.com/rfashwall/task-service/internal/models"
)

type TaskCommandHandler struct {
	TaskCommand command.TaskCommand
}

func NewTaskCommandHandler(taskCommand command.TaskCommand) *TaskCommandHandler {
	return &TaskCommandHandler{TaskCommand: taskCommand}
}

func (h *TaskCommandHandler) SetupRoutes(app *fiber.App) {
	app.Post("/tasks", h.createTask)
	app.Put("/tasks/:id", h.updateTask)
	app.Delete("/tasks/:id", h.deleteTask)
}

func (h *TaskCommandHandler) createTask(c *fiber.Ctx) error {
	task := new(models.Task)
	if err := c.BodyParser(task); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid input")
	}

	err := h.TaskCommand.CreateTask(c.Context(), task)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	return c.Status(fiber.StatusCreated).JSON(task)
}

func (h *TaskCommandHandler) updateTask(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid task ID")
	}

	task := new(models.Task)
	if err := c.BodyParser(task); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid input")
	}
	task.ID = id

	err = h.TaskCommand.UpdateTask(c.Context(), task)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	return c.JSON(task)
}

func (h *TaskCommandHandler) deleteTask(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid task ID")
	}

	err = h.TaskCommand.DeleteTask(c.Context(), id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	return c.SendStatus(fiber.StatusNoContent)
}
