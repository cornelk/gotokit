package database

import (
	"context"

	"github.com/jackc/pgx/v5/tracelog"
)

// Logger defines the logger interface used by the database implementation.
type Logger interface {
	Log(ctx context.Context, level tracelog.LogLevel, msg string, data map[string]interface{})
	Level() tracelog.LogLevel
}
