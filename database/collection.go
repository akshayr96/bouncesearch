package database

import "fmt"

type Collection struct {
	databaseName, name string
}

//Inserts a document into a collection
func (collection Collection) InsertOne(record string, id string) {

}

//Finds a phrase in a collection
func (collection Collection) Search(query string) {

}

//Deletes the current collection
func (collection Collection) Delete(query string) {

}

func (collection Collection) PrintStats() {
	fmt.Println("stats: Connected to", collection.databaseName, "database,", collection.name, "collection")
}
