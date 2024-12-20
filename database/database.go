// Package database implements PostgreSQL database client helpers.
package database

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"net"

	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/jackc/pgx/v5/stdlib" // register as sql driver
	"github.com/jackc/pgx/v5/tracelog"
)

// Connection defines a database connection. It exposes the exported functions of both embedded types.
// The connection is not concurrency-safe.
type Connection struct {
	*pgx.Conn
	*pgxscan.API
}

// Pool defines a connection pool. It exposes the exported functions of both embedded types.
// The pool is concurrency-safe.
type Pool struct {
	*pgxpool.Pool
	*pgxscan.API
}

type (
	// CommandTag is the status text returned by PostgreSQL for a query.
	CommandTag = pgconn.CommandTag
	// Querier is something that pgxscan can query and get the pgx.Rows from.
	// For example, it can be: *pgxpool.Pool, *pgx.Conn or pgx.Tx.
	Querier = pgxscan.Querier
	// Row is a convenience wrapper over Rows that is returned by QueryRow.
	Row = pgx.Row
	// Rows is the result set returned from *Conn.Query. Rows must be closed before
	// the *Conn can be used again. Rows are closed by explicitly calling Close().
	Rows = pgx.Rows
)

var (
	// ErrNoRows occurs when rows are expected but none are returned.
	ErrNoRows = pgx.ErrNoRows
	// ErrRowAlreadyExists occurs when a row already exists when trying to insert.
	ErrRowAlreadyExists = errors.New("row already exists")
	// ErrTooManyRows occurs when more rows than expected are returned.
	ErrTooManyRows = pgx.ErrTooManyRows
)

// New establishes a connection with a PostgreSQL database.
func New(ctx context.Context, cfg Config) (*Connection, error) {
	if err := cfg.Validate(); err != nil {
		return nil, fmt.Errorf("validating config: %w", err)
	}

	hostPort := net.JoinHostPort(cfg.Host, cfg.Port)
	url := fmt.Sprintf("postgres://%s:%s@%s/%s", cfg.User, cfg.Password, hostPort, cfg.Database)
	connConfig, err := pgx.ParseConfig(url) // nolint: contextcheck
	if err != nil {
		return nil, fmt.Errorf("parsing config: %w", err)
	}
	if cfg.Logger != nil {
		connConfig.Tracer = &tracelog.TraceLog{
			Logger:   cfg.Logger,
			LogLevel: cfg.Logger.Level(),
		}
	}

	pgxConn, err := pgx.ConnectConfig(ctx, connConfig)
	if err != nil {
		return nil, fmt.Errorf("connecting to database: %w", err)
	}

	scan, err := pgxscan.NewDBScanAPI()
	if err != nil {
		return nil, fmt.Errorf("creating db scan api: %w", err)
	}
	scanAPI, err := pgxscan.NewAPI(scan)
	if err != nil {
		return nil, fmt.Errorf("creating scan api: %w", err)
	}

	conn := &Connection{
		Conn: pgxConn,
		API:  scanAPI,
	}
	return conn, nil
}

// NewStdlib establishes a connection with a PostgreSQL database and returns a *sql.DB object.
func NewStdlib(ctx context.Context, cfg Config) (*sql.DB, error) {
	if err := cfg.Validate(); err != nil {
		return nil, fmt.Errorf("validating config: %w", err)
	}

	hostPort := net.JoinHostPort(cfg.Host, cfg.Port)
	url := fmt.Sprintf("postgres://%s:%s@%s/%s", cfg.User, cfg.Password, hostPort, cfg.Database)
	db, err := sql.Open("pgx", url)
	if err != nil {
		return nil, fmt.Errorf("creating database connection: %w", err)
	}

	if err := db.PingContext(ctx); err != nil {
		return nil, fmt.Errorf("pinging database: %w", err)
	}

	return db, nil
}

// NewPool establishes a connection pool with a PostgreSQL database.
func NewPool(ctx context.Context, cfg Config) (*Pool, error) {
	if err := cfg.Validate(); err != nil {
		return nil, fmt.Errorf("validating config: %w", err)
	}

	hostPort := net.JoinHostPort(cfg.Host, cfg.Port)
	url := fmt.Sprintf("postgres://%s:%s@%s/%s", cfg.User, cfg.Password, hostPort, cfg.Database)
	connConfig, err := pgxpool.ParseConfig(url) // nolint: contextcheck
	if err != nil {
		return nil, fmt.Errorf("parsing config: %w", err)
	}
	if cfg.Logger != nil {
		connConfig.ConnConfig.Tracer = &tracelog.TraceLog{
			Logger:   cfg.Logger,
			LogLevel: cfg.Logger.Level(),
		}
	}

	pool, err := pgxpool.NewWithConfig(ctx, connConfig)
	if err != nil {
		return nil, fmt.Errorf("connecting to database: %w", err)
	}

	scan, err := pgxscan.NewDBScanAPI()
	if err != nil {
		return nil, fmt.Errorf("creating db scan api: %w", err)
	}
	scanAPI, err := pgxscan.NewAPI(scan)
	if err != nil {
		return nil, fmt.Errorf("creating scan api: %w", err)
	}

	conn := &Pool{
		Pool: pool,
		API:  scanAPI,
	}
	return conn, nil
}

// Close closes all connections in the pool and rejects future Acquire calls.
// Blocks until all connections are returned to pool and closed.
// Implement the same Close function signature to allow Pool and Connection to have the
// same function signatures.
func (p *Pool) Close(_ context.Context) error {
	p.Pool.Close()
	return nil
}
