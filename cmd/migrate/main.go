package main

import (
	"fmt"
	"os"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
	m, err := migrate.New(
		"file://db/migrations",
		fmt.Sprintf("%s?sslmode=enable", os.Getenv("DATABASE_URL")))

	if err != nil {
		panic(err)
	}

	err = m.Up()
	if err != nil {
		panic(err)
	}
}