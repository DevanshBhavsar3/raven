package tests

import (
	"context"
	"fmt"
	"log"
	"path/filepath"
	"testing"

	"github.com/DevanshBhavsar3/raven/internal/config"
	"github.com/DevanshBhavsar3/raven/internal/database"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/moby/moby/api/types/network"
	"github.com/stretchr/testify/require"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/mysql"
	"github.com/testcontainers/testcontainers-go/wait"
)

func GetMySQLTestDB(ctx context.Context, t *testing.T) *sqlx.DB {
	t.Helper()
	cfg := config.Load()

	mysqlContainer, err := mysql.Run(ctx,
		"mysql",
		mysql.WithDatabase(cfg.Database.Name),
		mysql.WithUsername(cfg.Database.User),
		mysql.WithPassword(cfg.Database.Password),
		mysql.WithScripts(filepath.Join("..", "..", "database", "init.sql")),
		testcontainers.WithWaitStrategy(
			wait.ForSQL("3306/tcp", "mysql", func(host string, port network.Port) string {
				return fmt.Sprintf(
					"%s:%s@tcp(%s:%s)/%s",
					cfg.Database.User,
					cfg.Database.Password,
					host,
					port.Port(),
					cfg.Database.Name,
				)
			}),
		),
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

	err = database.Migrate(ctx, cfg)
	if err != nil {
		log.Fatalf("error migrating database: %v", err)
	}

	t.Cleanup(func() {
		if err := testcontainers.TerminateContainer(mysqlContainer); err != nil {
			log.Printf("failed to terminate container: %s", err)
		}
	})

	return db
}
