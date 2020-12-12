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
		fmt.Sprintf("%s?sslmode=require", os.Getenv("DATABASE_URL")))

	if err != nil {
		panic(err)
	}

	println("Running migrations...")
	err = m.Up()
	if err != nil {
		fmt.Printf("at=migrate status=skipped reason='%s'\n", err)
	} else {
		println("at=migrate status=done")
	}
}