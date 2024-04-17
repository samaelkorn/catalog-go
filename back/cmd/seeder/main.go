package main

import (
	"flag"
	"fmt"

	"github.com/samaelkorn/go-catalog/internal/database"
)

type config struct {
	db struct {
		dsn         string
		automigrate bool
	}
}

type application struct {
	db *database.DB
}

func main() {
	var cfg config
	flag.StringVar(&cfg.db.dsn, "db-dsn", "db.sqlite", "sqlite3 DSN")

	db, err := database.New(cfg.db.dsn, cfg.db.automigrate)

	if err != nil {
		fmt.Print("Ошибка подключения к базе")
	}

	defer db.Close()

	app := &application{
		db: db,
	}

	app.seederStatus()
	app.seederColors()
}
