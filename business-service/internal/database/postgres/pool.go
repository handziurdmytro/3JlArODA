package postgres

import (
	"context"
	"errors"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
)

var (
	ErrPingFailed = errors.New("ping to the db failed")
)

func GetNewPool(ctx context.Context, dsn string) (*pgxpool.Pool, error) {
	pool, err := pgxpool.New(ctx, dsn)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	if err := pool.Ping(ctx); err != nil {
		log.Println(err)
		return nil, ErrPingFailed
	}

	return pool, nil
}
