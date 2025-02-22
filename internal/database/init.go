package database

import (
	"database/sql"
	"iskra/centralized/internal/config"
	"iskra/centralized/internal/helpers"
	"log"

	_ "github.com/mattn/go-sqlite3"
	"github.com/pressly/goose/v3"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/sqlitedialect"
)

func Init() (*bun.DB, error) {
	config := config.Initialize()

	var sqlDB *sql.DB
	var err error

	if config.APPLICATION_STATUS == "testing" {
		sqlDB, err = sql.Open("sqlite3", "file:?mode=memory&cache=shared")
	} else {
		sqlDB, err = sql.Open("sqlite3", "./internal/database/iskra.db")
	}

	if err != nil {
		return nil, err
	}

	if err := goose.SetDialect("sqlite3"); err != nil {
		return nil, err
	}

	if err := goose.Up(sqlDB, helpers.GetProjectRoot()+"/internal/database/migrations"); err != nil {
		return nil, err
	}

	log.Println("Migrations ran successfully")

	db := bun.NewDB(sqlDB, sqlitedialect.New())
	log.Println("Bun DB initialized successfully")

	return db, nil
}
