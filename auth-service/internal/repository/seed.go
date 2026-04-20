package repository

import (
	"context"
	"log/slog"

	"github.com/google/uuid"
)

func (r *Repository) SeedDefaultAdmin(ctx context.Context, username, password string, hashFn func(string) (string, error)) error {

	exists, err := r.GetUserByUsername(ctx, username)
	if err == nil && exists != nil {
		slog.Info("default admin already exists, skipping seed")
		return nil
	}

	hash, err := hashFn(password)
	if err != nil {
		return err
	}

	_, err = r.db.Exec(ctx, queryCreateUser, uuid.New(), username, hash)
	if err != nil {
		return err
	}

	slog.Info("successfully seeded default admin", slog.String("username", username))
	return nil
}
