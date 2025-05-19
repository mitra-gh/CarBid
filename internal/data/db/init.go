package db

import (
	"context"
	"fmt"
	"sync"

	"github.com/jackc/pgx/v5"
	"github.com/mitra-gh/CarBid/configs"
	"github.com/mitra-gh/CarBid/internal/data/db/sqlc"
)

var (
	queries     *sqlc.Queries
	conn        *pgx.Conn
	initOnce    sync.Once
	initErr     error
)

func InitDB(cfg *configs.Config) error {
    initOnce.Do(func() {
        initErr = initializeDB(cfg)
    })
    return initErr
}

func initializeDB(cfg *configs.Config) error {
	ctx := context.Background()

	NewConn, err := pgx.Connect(ctx, buildConnectionString(cfg))
	if err != nil {
		return fmt.Errorf("failed to connect to database: %w", err)
	}

	conn = NewConn


	// Verify connection with a ping
	if err := conn.Ping(ctx); err != nil {
		conn.Close(ctx)
		return fmt.Errorf("database ping failed: %w", err)
	}
	
	queries = sqlc.New(conn)
	return nil
}

func buildConnectionString(cfg *configs.Config) string {
	return fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		cfg.Postgres.Host,
		cfg.Postgres.Port,
		cfg.Postgres.User,
		cfg.Postgres.Password,
		cfg.Postgres.DBName,
		cfg.Postgres.SSLMode,
	)
}

func GetQueries() (*sqlc.Queries, error) {
	if queries == nil {
		return nil, fmt.Errorf("database not initialized - call InitDB first")
	}
	return queries, nil
}

func CloseDB() error {
	if conn == nil {
		return fmt.Errorf("database not initialized - call InitDB first")
	}
	return conn.Close(context.Background())
}