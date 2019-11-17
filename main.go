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
		fmt.Println("Created")
	} else {
		fmt.Println(err)
	}

	//connect to database
	conn, err := database.Connect("random_app")
	if err == nil {
		//user schema created
		userSchema := types.Schema{
			"description": types.AttributeMeta{Weight: 1.0, Optional: false},
			"reportName":  types.AttributeMeta{Weight: 1.0, Optional: false},
		}
		users := conn.Collection("users", userSchema)
		users.PrintStats()
	} else {
		fmt.Println(err)
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
