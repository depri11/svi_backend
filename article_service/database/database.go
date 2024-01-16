package database

import (
	"database/sql"
	"log"

	mysql "github.com/go-sql-driver/mysql"
)

func InitDatabase() (*sql.DB, error) {
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
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		log.Println(err)
		return nil, err
	}

	err = RunMigrate(db)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return db, nil
}
