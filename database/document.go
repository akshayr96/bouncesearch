package database

import "fmt"

type Document struct {
	databaseName, collectionName string
}

func (document Document) InsertOne(record string, documentName string) {

}

func (document Document) FindOne(query string) {

}

func (document Document) UpdateOne(recordId string, updates string, documentName string) {

}

func (document Document) DeleteOne(recordId string) {

}

func (document Document) Me() {
	fmt.Println(document.databaseName, document.collectionName)
}
