package db

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/go-sql-driver/mysql"
)

// New establishes a connection to the database and returns the connection pool.
func New(user, password, host, port, dbname string) (*sql.DB, error) {

	cfg := mysql.NewConfig()
	cfg.User = user
	cfg.Passwd = password
	cfg.Net = "tcp"
	cfg.Addr = host + ":" + port
	cfg.DBName = dbname
	cfg.ParseTime = true
	cfg.AllowNativePasswords = true

	var db *sql.DB
	var err error

	// Retry connection for up to 30 seconds
	for i := 0; i < 10; i++ {
		db, err = sql.Open("mysql", cfg.FormatDSN())
		if err == nil {
			err = db.Ping()
			if err == nil {
				return db, nil
			}
		}

		log.Printf("Database not ready, retrying in 3 seconds... (%d/10)", i+1)
		time.Sleep(3 * time.Second)
	}

	return nil, fmt.Errorf("failed to connect to database after retries: %w", err)
}
