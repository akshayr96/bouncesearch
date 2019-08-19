package main

import (

	// "log"
	// "net/http"

	"fmt"

	"github.com/akshayr96/bounceSearch/database"
	"github.com/akshayr96/bounceSearch/types"
	"github.com/akshayr96/bounceSearch/utils"
)

func main() {
	tree := &types.Node{}

	//Inserting words into a ternary tree
	utils.InsertWord(tree, "app")
	utils.InsertWord(tree, "apple")
	utils.InsertWord(tree, "bat")

	// Searching for a words in the tree
	fmt.Println(types.CompleteMatch == utils.Parser(tree, "app"))
	fmt.Println(types.CompleteMatch == utils.Parser(tree, "ball"))

	//Getting the Levenshtien distance between two words
	fmt.Println(utils.LevenshteinDistance("app", "map"))

	// log.Fatal(http.ListenAndServe(":8080", routes.Routes()))

	//API Example
	conn := database.Connect("RandomApp")
	users := conn.Collection("users")
	users.Me()
}
