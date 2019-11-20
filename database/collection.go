package database

import (
	"errors"
	"fmt"

	"github.com/akshayr96/bounceSearch/constants"
	"github.com/akshayr96/bounceSearch/types"
	"github.com/akshayr96/bounceSearch/utilities"
)

type Collection struct {
	databaseName string
	name         string
	schema       types.Schema
}

//Inserts a document into a collection
func (collection Collection) InsertOne(record types.Record, id string) error {
	if utilities.ValidateSchema(collection.schema, record) {
		fmt.Println("Valid Schema")
	} else {
		return errors.New(constants.InvalidRecord)
	}
	return nil
}

//Finds a phrase in a collection
func (collection Collection) Search(query string) {

}

//Deletes the current collection
func (collection Collection) Delete(query string) {
	//@todo: reindex upon delete
}

func (collection Collection) PrintStats() {
	fmt.Println("stats: Connected to", collection.databaseName, "database,", collection.name, "collection")

}
