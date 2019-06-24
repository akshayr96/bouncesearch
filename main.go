package main

import (

	// "log"
	// "net/http"

	"fmt"

	"github.com/akshayr96/bounceSearch/types"
	"github.com/akshayr96/bounceSearch/utils"
)

func main() {
	tree := &types.Node{}
	utils.InsertWord(tree, "app")
	utils.InsertWord(tree, "apple")
	utils.InsertWord(tree, "bat")
	fmt.Println(tree)

	fmt.Println(utils.Parser(tree, "app"))  // Prints true
	fmt.Println(utils.Parser(tree, "ball")) // Prints false

	// log.Fatal(http.ListenAndServe(":8080", routes.Routes()))
}