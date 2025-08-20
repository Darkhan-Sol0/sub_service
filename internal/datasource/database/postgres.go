package database

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Config interface {
	GetDBEnv() string
	GetDBPort() string
	GetDBHost() string
	GetDBDatabase() string
	GetDBUsername() string
	GetDBPassword() string
}

func ConnectDB(ctx context.Context, cfg Config) (pool *pgxpool.Pool, err error) {
	dns := fmt.Sprintf("postgres://%s:%s@%s:%s/%s",
		cfg.GetDBUsername(),
		cfg.GetDBPassword(),
		cfg.GetDBHost(),
		cfg.GetDBPort(),
		cfg.GetDBDatabase(),
	)
	pool, err = pgxpool.New(ctx, dns)
	if err != nil {
		return nil, fmt.Errorf("database connection failed: %v", err)
	}
	if err := pool.Ping(ctx); err != nil {
		return nil, fmt.Errorf("database ping failed: %v", err)
	}
	return pool, nil
	// return nil, nil
}
