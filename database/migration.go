package database

import (
	"database/sql"
	"fmt"
	"io/fs"
	"net/http"

	migrate "github.com/rubenv/sql-migrate"
)

// MigrationDirection defines the direction of the migration.
type MigrationDirection = migrate.MigrationDirection

const (
	// UpMigration executes migrations that are new and have not been applied yet.
	UpMigration = migrate.Up
	// DownMigration migrates down.
	DownMigration = migrate.Down
)

// Migrate the given database with the embedded migrations and direction.
func Migrate(db *sql.DB, migrations fs.FS, direction MigrationDirection) (int, error) {
	return MigrateMax(db, migrations, direction, 0)
}

// MigrateMax migrates the given database with the embedded migrations and direction,
// will apply at most `max` migrations, pass 0 for no limit.
func MigrateMax(db *sql.DB, migrations fs.FS, direction MigrationDirection, maxAmount int) (int, error) {
	source := &migrate.HttpFileSystemMigrationSource{
		FileSystem: http.FS(migrations),
	}

	applied, err := migrate.ExecMax(db, "postgres", source, direction, maxAmount)
	if err != nil {
		return 0, fmt.Errorf("executing migrations: %w", err)
	}

	return applied, nil
}
