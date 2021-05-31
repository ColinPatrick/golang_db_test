package main

import (
	"fmt"
	"log"
	"os"
	"database/sql"
	"github.com/joho/godotenv"
	_ "github.com/go-sql-driver/mysql"
)

func goDotEnvVariable(key string) string {

	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)

}

func main() {
	fmt.Println("Go MySQL Tutorial")

	dotenv := goDotEnvVariable("DB_INFO")

	db, err := sql.Open("mysql", dotenv)

	if err != nil {
		panic(err.Error())		
	}

	defer db.Close()

	results, err := db.Query("SELECT * FROM places")

	if err != nil {
		panic(err.Error())
	}

	defer results.Close()

	for results.Next() {
		var (
			id int
			name string
			location string
		)
		err = results.Scan(&id, &name, &location)
		if err != nil {
			panic(err.Error())
		}
		fmt.Printf("ID: %d, Name: %s, Location: %s\n", id, name, location)
	}

	var (
		id int
		name string
		location string	
	)

	err = db.QueryRow("SELECT * FROM places WHERE id = 1").Scan(&id, &name, &location)

	if err != nil {
		panic(err.Error())
	}

	fmt.Printf("ID: %d, Name: %s, Location: %s\n", id, name, location)

	// insert, err := db.Query("INSERT INTO places (name, location) VALUES ('Mordor', 'Middle Earth')")

	// if err != nil {
	// 	panic(err.Error())
	// }

	// defer insert.Close()

	// fmt.Println("You entered something into a database!")
}