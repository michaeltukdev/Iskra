package database

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
	"github.com/pressly/goose/v3"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/sqlitedialect"
)

var DB *bun.DB

func Init() {
	sqldb, err := sql.Open("sqlite3", "./internal/database/iskra.db")
	if err != nil {
		log.Fatalf("failed to open SQLite database: %v", err)
	}

	if err := goose.SetDialect("sqlite3"); err != nil {
		log.Fatalf("failed to set goose dialect: %v", err)
	}

	if err := goose.Up(sqldb, "./internal/database/migrations"); err != nil {
		log.Fatalf("failed to run migrations: %v", err)
	}

	log.Println("Migrations ran successfully")

	DB = bun.NewDB(sqldb, sqlitedialect.New())
	log.Println("Bun DB initialized successfully")
}
