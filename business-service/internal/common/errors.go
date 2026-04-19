package common

import (
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

var (
	ErrDatabaseConnection  = errors.New("failed to connect to db")
	ErrNotFound            = errors.New("not found")
	ErrAlreadyExists       = errors.New("already exists")
	ErrForeignKeyViolation = errors.New("foreign key violation")
	ErrCheckViolation      = errors.New("check violation")
	ErrNotNullViolation    = errors.New("not null violation")
)

func MapRepositoryError(operation string, err error) error {
	if errors.Is(err, pgx.ErrNoRows) {
		return fmt.Errorf("%s: %w", operation, ErrNotFound)
	}

	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) {
		switch pgErr.Code {
		case "23505":
			return fmt.Errorf("%s: %w", operation, ErrAlreadyExists)
		case "23503":
			return fmt.Errorf("%s: %w", operation, ErrForeignKeyViolation)
		case "23514":
			return fmt.Errorf("%s: %w", operation, ErrCheckViolation)
		case "23502":
			return fmt.Errorf("%s: %w", operation, ErrNotNullViolation)
		}
	}

	return fmt.Errorf("%s: %w", operation, err)
}
