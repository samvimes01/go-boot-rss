package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	"github.com/samvimes01/go-rss/internal/env"
)

func InitDb(e *env.Env) *Queries {
	ssl := ""
	if e.AppEnv != "production" {
		ssl = "?sslmode=disable"
	}
	dbURL := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s%s", e.DBUser, e.DBPass, e.DBHost, e.DBPort, e.DBName, ssl)

	db, err := sql.Open("postgres", dbURL)
  if err != nil {
    panic(err)
  }

	return New(db)
}
