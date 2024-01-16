package main

import (
	"log"
	"sharing_vasion_indonesia/article_service/database"
)

func main() {

	err := database.InitDatabase()
	if err != nil {
		log.Println(err)
	}

}
