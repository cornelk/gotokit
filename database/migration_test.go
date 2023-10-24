package database

import (
	"context"
	"database/sql"
	"testing"
	"testing/fstest"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	_ "modernc.org/sqlite" // register driver
)

func TestMigrate(t *testing.T) {
	db, err := sql.Open("sqlite", ":memory:")
	require.NoError(t, err)

	ctx := context.Background()
	instruction := `INSERT INTO test(id) VALUES(1)`
	_, err = db.ExecContext(ctx, instruction)
	require.Error(t, err)

	var migrationFs = fstest.MapFS{
		"1234567890_init.sql": {
			Data: []byte(`
-- +migrate Up

CREATE TABLE test
(
    id INTEGER PRIMARY KEY NOT NULL
);

-- +migrate Down

DROP TABLE IF EXISTS test;
`),
		},
	}

	applied, err := Migrate(db, migrationFs, UpMigration)
	require.NoError(t, err)
	assert.Equal(t, 1, applied)

	_, err = db.ExecContext(ctx, instruction)
	require.NoError(t, err)

	require.NoError(t, db.Close())

	_, err = Migrate(db, migrationFs, UpMigration)
	require.Error(t, err)
}
