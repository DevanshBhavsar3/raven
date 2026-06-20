package tests

import (
	"context"
	"log"
	"path/filepath"
	"testing"

	"github.com/DevanshBhavsar3/raven/internal/config"
	"github.com/DevanshBhavsar3/raven/internal/database"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/require"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/mysql"
)

func GetMySQLTestDB(ctx context.Context, t *testing.T) *sqlx.DB {
	t.Helper()
	cfg := config.Load()

	mysqlContainer, err := mysql.Run(ctx,
		"mysql:8.0.36",
		mysql.WithDatabase(cfg.Database.Name),
		mysql.WithUsername(cfg.Database.User),
		mysql.WithPassword(cfg.Database.Password),
		mysql.WithScripts(filepath.Join("..", "..", "database", "init.sql")),
	)
	require.NoError(t, err)

	dsn, err := mysqlContainer.ConnectionString(ctx)
	require.NoError(t, err)

	db, err := sqlx.Connect("mysql", dsn)
	require.NoError(t, err)

	host, err := mysqlContainer.Host(ctx)
	require.NoError(t, err)

	port, err := mysqlContainer.MappedPort(ctx, "3306/tcp")
	require.NoError(t, err)

	cfg.Database.Host = host
	cfg.Database.Port = int(port.Num())
	database.Migrate(ctx, cfg)

	t.Cleanup(func() {
		if err := testcontainers.TerminateContainer(mysqlContainer); err != nil {
			log.Printf("failed to terminate container: %s", err)
		}
	})

	return db
}
