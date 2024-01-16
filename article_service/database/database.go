package database

import (
	"database/sql"
	"log"

	mysql "github.com/go-sql-driver/mysql"
)

func InitDatabase() error {
	config := mysql.Config{
		User:   "root",
		Passwd: "",
		Net:    "tcp",
		Addr:   "localhost:3306",
		DBName: "article",
	}
	dsn := config.FormatDSN()

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Println(err)
		log.Fatal(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Println(err)
		return err
	}

	err = RunMigrate(db)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}
