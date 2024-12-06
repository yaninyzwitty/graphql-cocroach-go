package database

import (
	"context"
	"fmt"
	"log/slog"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

type DatabaseCfg struct {
	User     string
	Password string
	Host     string
	Port     int
	Database string
	SSLMode  string
}

var (
	pool *pgxpool.Pool
)

// NewPgxPool creates a new pgxpool.Pool with retries and proper error handling.
func NewPgxPool(ctx context.Context, cfg *DatabaseCfg, maxRetries int) (*pgxpool.Pool, error) {
	if cfg.User == "" || cfg.Password == "" || cfg.Host == "" || cfg.Database == "" {
		return nil, fmt.Errorf("missing required database configuration")
	}

	// Construct DSN
	dsn := fmt.Sprintf(
		"postgresql://%s:%s@%s:%d/%s?sslmode=%s",
		cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.Database, cfg.SSLMode,
	)

	var err error
	for retriesLeft := maxRetries; retriesLeft > 0; retriesLeft-- {
		// Check if the context has been canceled
		select {
		case <-ctx.Done():
			return nil, fmt.Errorf("context canceled: %w", ctx.Err())
		default:
		}

		// Parse the configuration
		config, parseErr := pgxpool.ParseConfig(dsn)
		if parseErr != nil {
			return nil, fmt.Errorf("unable to parse the database URL: %w", parseErr)
		}

		// Connection pool settings
		config.MaxConns = 30
		config.MaxConnIdleTime = 5 * time.Minute
		config.HealthCheckPeriod = 1 * time.Minute
		config.MinConns = 1

		// Try to create the connection pool
		pool, err = pgxpool.NewWithConfig(ctx, config)
		if err == nil {
			return pool, nil
		}

		// Log the connection error and retry
		slog.Info("Failed to connect to database, retrying", "error", err, "retriesLeft", retriesLeft-1)
		time.Sleep(500 * time.Millisecond)
	}

	// If no connection was made after maxRetries, return the last error encountered
	return nil, fmt.Errorf("unable to connect to database after %d retries: %w", maxRetries, err)
}

func PingDatabase(ctx context.Context, pool *pgxpool.Pool) error {
	// Attempt to ping the database
	err := pool.Ping(ctx)
	if err != nil {
		// Log the error and return it if the ping fails
		slog.Error("Failed to ping database", "error", err)
		return fmt.Errorf("failed to ping database: %w", err)
	}
	// Log success if the ping was successful
	slog.Info("Successfully pinged the database")
	return nil
}
