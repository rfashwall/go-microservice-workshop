package command

import (
	"context"
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/rfashwall/task-service/internal/models"
)

type TaskCommand interface {
	CreateTask(ctx context.Context, task *models.Task) error
	UpdateTask(ctx context.Context, task *models.Task) error
	DeleteTask(ctx context.Context, id int) error
}

type MySQLTaskCommand struct {
	Conn *sql.DB
}

func NewMySQLTaskCommand(conn *sql.DB) *MySQLTaskCommand {
	return &MySQLTaskCommand{Conn: conn}
}

func (c *MySQLTaskCommand) CreateTask(ctx context.Context, task *models.Task) error {
	_, err := c.Conn.ExecContext(ctx, "INSERT INTO tasks (user_id, title, description, status) VALUES (?, ?, ?, ?)", task.UserID, task.Title, task.Description, task.Status)
	return err
}

func (c *MySQLTaskCommand) UpdateTask(ctx context.Context, task *models.Task) error {
	_, err := c.Conn.ExecContext(ctx, "UPDATE tasks SET title=?, description=?, status=? WHERE id=?", task.Title, task.Description, task.Status, task.ID)
	return err
}

func (c *MySQLTaskCommand) DeleteTask(ctx context.Context, id int) error {
	_, err := c.Conn.ExecContext(ctx, "DELETE FROM tasks WHERE id=?", id)
	return err
}
