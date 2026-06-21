package database

import (
	"fmt"
	"log"
	"net"
	"net/url"
	"strconv"

	"github.com/DevanshBhavsar3/raven/internal/config"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

func New(config *config.ApplicationConfig) *sqlx.DB {
	hostPort := net.JoinHostPort(config.Database.Host, strconv.Itoa(config.Database.Port))
	encodedPassword := url.QueryEscape(config.Database.Password)
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=true&multiStatements=true",
		config.Database.User,
		encodedPassword,
		hostPort,
		config.Database.Name,
	)

	db, err := sqlx.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("error connecting to database: %v", err)
	}
	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(3)

	return db
}
