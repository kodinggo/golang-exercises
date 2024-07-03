package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	migrate "github.com/rubenv/sql-migrate"
)

func main() {
	// db environment
	var (
		dbUsername = "root"
		dbPassword = ""
		dbName     = "kodinggo"
		dbHost     = "localhost:3306"
	)

	// prepare connection string
	// charset=utf8mb4 agar dapat menyimpan semua karakter unicode
	// parseTime=true agar dapat diparsing dari timestamp ke tipe time.Time
	connStr := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=true",
		dbUsername,
		dbPassword,
		dbHost,
		dbName)
	connDB, err := sql.Open("mysql", connStr)
	if err != nil {
		log.Fatalf("failed to connect server db, error: %s", err.Error())
	}

	// Perform migration
	var (
		direction = "down"
		step      = 0
	)
	migrations := &migrate.FileMigrationSource{Dir: "."}
	var n int // "n" berapa migration yg di-applied
	if direction == "down" {
		n, err = migrate.ExecMax(connDB, "mysql", migrations, migrate.Down, step)
	} else {
		n, err = migrate.ExecMax(connDB, "mysql", migrations, migrate.Up, step)
	}
	if err != nil {
		log.Fatalf("failed to run migration, error: %s", err.Error())
	}

	log.Printf("successfully applied %d migration(s)", n)
}
