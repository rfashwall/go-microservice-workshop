package repository

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/rfashwall/user-service/internal/models"
)

type Repository interface {
	GetUserByID(ctx context.Context, id int) (*models.User, error)
	// Additional CRUD methods can be defined here
}
type PostgresUserRepository struct {
	Conn *pgx.Conn
}

func NewPostgresUserRepository(conn *pgx.Conn) *PostgresUserRepository {
	return &PostgresUserRepository{Conn: conn}
}

func (r *PostgresUserRepository) GetUserByID(ctx context.Context, id int) (*models.User, error) {
	var user models.User
	err := r.Conn.QueryRow(ctx, "SELECT id, name, email, password FROM users WHERE id=$1", id).Scan(&user.ID, &user.Name, &user.Email, &user.Password)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// Additional CRUD methods can be implemented here
