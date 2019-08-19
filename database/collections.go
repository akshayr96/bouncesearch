package database

type Collections struct {
	databaseName string
}

func (collections Collections) CreateOne(name string, schema string) {

}

func (collections Collections) DeleteOne(modelId string) {

}

func (collections Collections) Collection(collectionName string) Document {
	//Check if the collection <collectionName> exists in collections.databaseName
	return Document{collections.databaseName, collectionName}
}
