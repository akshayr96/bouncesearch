package database

import "github.com/akshayr96/bounceSearch/types"

type Collections struct {
	databaseName string
}

func (collections Collections) Delete(modelId string) {
	//Check if collection exists
	//Remove collection
	//return success
}

func (collections Collections) Collection(collectionName string, schema types.Schema) Collection {
	//Check if the collection <collectionName> exists in collections.databaseName
	return Collection{collections.databaseName, collectionName, schema}
}
