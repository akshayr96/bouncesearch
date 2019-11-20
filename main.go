package main

import (
	"fmt"

	"github.com/akshayr96/bounceSearch/database"
	"github.com/akshayr96/bounceSearch/types"
)

func main() {
	//create database
	err := database.Create("random_app")
	if err == nil {
		fmt.Println("Database Created")
	} else {
		fmt.Println(err)
	}

	//connect to database
	conn, err := database.Connect("random_app")
	if err == nil {
		//user schema created
		characterSchema := types.Schema{
			"characterName": types.AttributeMeta{Weight: 1.0, Optional: false},
			"description":   types.AttributeMeta{Weight: 1.0, Optional: false},
		}
		characters := conn.Collection("characters", characterSchema)

		newCharacter := types.Record{
			"characterName": "foo",
			"description":   "test",
		}

		characters.InsertOne(newCharacter, "1")
	} else {
		fmt.Println(err)
	}

	err = conn.Delete("characters")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Deleted the collection")
	}

	//drop database
	err = database.Drop("random_app")
	if err == nil {
		fmt.Println("Dropped the database")
	} else {
		fmt.Println(err)
	}

	// API Example
	// log.Fatal(http.ListenAndServe(":8080", routes.Routes()))
}
