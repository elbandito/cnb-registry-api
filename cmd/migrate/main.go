package main

import (
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	"os"
)

func main() {
	m, err := migrate.New(
		"db/migrations",
		fmt.Sprintf("%s?sslmode=enable", os.Getenv("DATABASE_URL")))

	if err != nil {
		panic(err)
	}

	m.Up()
}