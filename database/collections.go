package database

import (
	"errors"

	"github.com/akshayr96/bounceSearch/constants"
	"github.com/akshayr96/bounceSearch/provider"
	"github.com/akshayr96/bounceSearch/types"
)

type Collections struct {
	databaseName string
}

func (collections Collections) Collection(collectionName string, schema types.Schema) Collection {
	return Collection{collections.databaseName, collectionName, schema}
}

func (collections Collections) Delete(collectionName string) error {
	if provider.CheckIfCollectionExists(collections.databaseName, collectionName) {
		err := provider.DeleteCollection(collections.databaseName, collectionName)
		if err != nil {
			return errors.New(constants.ErrorDeletingCollection)
		}
		return nil
	} else {
		return errors.New(constants.CollectionDoesntExist)
	}
}
