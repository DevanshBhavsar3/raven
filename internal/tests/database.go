package tests

import (
	"context"
	"log"
	"path/filepath"
	"runtime"

	"github.com/DevanshBhavsar3/raven/internal/config"
	"github.com/DevanshBhavsar3/raven/internal/database"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/mysql"
)

type MySQLContainer struct {
	db        *sqlx.DB
	container testcontainers.Container
}

func NewMySQLTestDB(ctx context.Context) *MySQLContainer {
	cfg := config.Load()

	_, file, _, ok := runtime.Caller(0)
	if !ok {
		panic("unable to determine caller")
	}

	helperDir := filepath.Dir(file)

	mysqlContainer, err := mysql.Run(ctx,
		"mysql:8.0.36",
		mysql.WithDatabase(cfg.Database.Name),
		mysql.WithUsername(cfg.Database.User),
		mysql.WithPassword(cfg.Database.Password),
		mysql.WithScripts(filepath.Join(helperDir, "../database/init.sql")),
	)
	if err != nil {
		log.Fatalf("failed to start MySQL container: %v", err)
	}

	dsn, err := mysqlContainer.ConnectionString(ctx, "parseTime=true&multiStatements=true")
	if err != nil {
		log.Fatalf("failed to get connection string: %v", err)
	}

	db, err := sqlx.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("failed to connect to MySQL: %v", err)
	}

	host, err := mysqlContainer.Host(ctx)
	if err != nil {
		log.Fatalf("failed to get MySQL container host: %v", err)
	}

	port, err := mysqlContainer.MappedPort(ctx, "3306/tcp")
	if err != nil {
		log.Fatalf("failed to get MySQL container port: %v", err)
	}

	cfg.Database.Host = host
	cfg.Database.Port = int(port.Num())

	err = database.Migrate(ctx, cfg)
	if err != nil {
		log.Fatalf("error migrating database: %v", err)
	}

	mysqlTestDB := &MySQLContainer{
		db:        db,
		container: mysqlContainer,
	}

	return mysqlTestDB
}

func (m *MySQLContainer) GetDB() *sqlx.DB {
	return m.db
}

func (m *MySQLContainer) Terminate(ctx context.Context) {
	if err := testcontainers.TerminateContainer(m.container); err != nil {
		log.Printf("failed to terminate container: %s", err)
	}
}
