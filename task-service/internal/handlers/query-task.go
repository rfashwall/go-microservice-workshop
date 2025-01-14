package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rfashwall/task-service/internal/query"
)

type TaskQueryHandler struct {
	TaskQuery query.TaskQuery
}

func NewTaskQueryHandler(taskQuery query.TaskQuery) *TaskQueryHandler {
	return &TaskQueryHandler{TaskQuery: taskQuery}
}

func (h *TaskQueryHandler) SetupRoutes(app *fiber.App) {
	app.Get("/tasks/:id", h.getTask)
	app.Get("/users/:user_id/tasks", h.listTasksByUserID)
}

func (h *TaskQueryHandler) getTask(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid task ID")
	}

	task, err := h.TaskQuery.GetTaskByID(c.Context(), id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	return c.JSON(task)
}

func (h *TaskQueryHandler) listTasksByUserID(c *fiber.Ctx) error {
	userID, err := c.ParamsInt("user_id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid user ID")
	}

	tasks, err := h.TaskQuery.ListTasksByUserID(c.Context(), userID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	return c.JSON(tasks)
}
