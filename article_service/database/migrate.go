package database

import (
	"database/sql"
	"log"
	"os"
	"path/filepath"

	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func RunMigrate(db *sql.DB) error {
	getwd, err := os.Getwd()
	if err != nil {
		log.Println(err)
		return err
	}

	path := "file://" + filepath.Join(getwd, "article_service/migrations")

	driver, err := mysql.WithInstance(db, &mysql.Config{})
	if err != nil {
		log.Println(err)
		return err
	}

	m, err := migrate.NewWithDatabaseInstance(
		path,
		"mysql",
		driver,
	)
	if err != nil {
		log.Println(err)
		return err
	}

	log.Println("migration up...")
	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatal(err)
	}
	log.Println("migration up done")

	return nil

}
