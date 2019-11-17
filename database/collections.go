package database

import "github.com/akshayr96/bounceSearch/types"

type Collections struct {
	databaseName string
}

func (collections Collections) CreateOne(name string, schema types.Schema) {
	//check if schema already exists
	//Persist schema (io)
	//return success
}

func (collections Collections) DeleteOne(modelId string) {
	//Check if collection exists
	//Remove collection
	//return success
}

func (collections Collections) Collection(collectionName string) Collection {
	//Check if the collection <collectionName> exists in collections.databaseName
	return Collection{collections.databaseName, collectionName}
}
