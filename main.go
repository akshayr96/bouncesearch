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
	utils.InsertWord(tree, "app", "6516516516", 5.5)
	utils.InsertWord(tree, "app", "561919189", 5.4)
	utils.InsertWord(tree, "apple", "161641640", 1.0)
	utils.InsertWord(tree, "bat", "56316516516", 0.5)

	// Searching for a words in the tree
	fmt.Println(types.CompleteMatch == utils.Parser(tree, "app"))
	fmt.Println(types.NoMatch == utils.Parser(tree, "ball"))
	fmt.Println(types.PartialMatch == utils.Parser(tree, "ap"))
	fmt.Println(types.NoMatch == utils.Parser(tree, "x"))

	utils.WriteTree(tree, "vbiview", "users")

	//Getting the Levenshtien distance between two words
	fmt.Println(utils.LevenshteinDistance("app", "map"))

	// log.Fatal(http.ListenAndServe(":8080", routes.Routes()))

	//API Example
	conn := database.Connect("RandomApp")
	users := conn.Collection("users")
	users.Me()
}
