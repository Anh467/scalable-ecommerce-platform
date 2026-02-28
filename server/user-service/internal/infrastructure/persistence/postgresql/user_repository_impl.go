package postgresql

import (
	"context"
	"database/sql"

	"github.com/Anh467/scalable-ecommerce-platform/server/user-service/internal/domain/entity"
	"github.com/Anh467/scalable-ecommerce-platform/server/user-service/internal/domain/repository"
	"github.com/google/uuid"
)

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) repository.UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) Create(ctx context.Context, user *entity.User) error {
	query := `
		INSERT INTO user_temp (id, email, password_hash, first_name, last_name, role)
		VALUES ($1, $2, $3, $4, $5, $6)
	`

	_, err := r.db.ExecContext(ctx, query,
		user.ID,
		user.Email,
		user.PasswordHash,
		user.FirstName,
		user.LastName,
		user.Role,
	)

	return err
}

func (r *userRepository) GetByID(ctx context.Context, id uuid.UUID) (*entity.User, error) {
	query := `
		SELECT id, email, password_hash, first_name, last_name, role
		FROM user_temp
		WHERE id = $1
	`

	row := r.db.QueryRowContext(ctx, query, id)

	var user entity.User
	err := row.Scan(
		&user.ID,
		&user.Email,
		&user.PasswordHash,
		&user.FirstName,
		&user.LastName,
		&user.Role,
	)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *userRepository) GetByEmail(ctx context.Context, email string) (*entity.User, error) {
	query := `
		SELECT id, email, password_hash, first_name, last_name, role
		FROM user_temp
		WHERE email = $1
	`

	row := r.db.QueryRowContext(ctx, query, email)

	var user entity.User
	err := row.Scan(
		&user.ID,
		&user.Email,
		&user.PasswordHash,
		&user.FirstName,
		&user.LastName,
		&user.Role,
	)

	if err != nil {
		return nil, err
	}

	return &user, nil
}