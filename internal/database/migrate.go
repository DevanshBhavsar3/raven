package database

import (
	"context"
	"embed"
	"errors"
	"fmt"
	"log"
	"net"
	"net/url"
	"strconv"

	"github.com/DevanshBhavsar3/raven/internal/config"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/golang-migrate/migrate/v4/source/iofs"
)

//go:embed migrations/*.sql
var migrations embed.FS

func Migrate(ctx context.Context, config *config.ApplicationConfig) error {
	hostPort := net.JoinHostPort(config.Database.Host, strconv.Itoa(config.Database.Port))
	encodedPassword := url.QueryEscape(config.Database.Password)
	dsn := fmt.Sprintf("mysql://%s:%s@tcp(%s)/%s?multiStatements=true",
		config.Database.User,
		encodedPassword,
		hostPort,
		config.Database.Name,
	)

	source, err := iofs.New(migrations, "migrations")
	if err != nil {
		return fmt.Errorf("error creating database migrations: %w", err)
	}

	m, err := migrate.NewWithSourceInstance("iofs", source, dsn)
	if err != nil {
		return fmt.Errorf("constructing database migration failed: %w", err)
	}

	from, _, _ := m.Version()

	if err := m.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return fmt.Errorf("error applying database migration: %w", err)
	}

	to, _, _ := m.Version()

	if from == to {
		log.Printf("database schema up to date, version %d", to)
	} else {
		log.Printf("migrated database schema, from %d to %d", from, to)
	}

	return nil
}
