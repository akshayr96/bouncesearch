package provider

import (
	"encoding/json"
	"os"
	"path"

	"github.com/akshayr96/bounceSearch/constants"
	"github.com/akshayr96/bounceSearch/types"
)

//Writes the given tree into a its file
func WriteTree(tree *types.Node, databaseName string, collectionName string) {
	filePath := path.Join(constants.DumpsDir, "storage", databaseName)
	fileName := path.Join(filePath, collectionName+".json")
	os.MkdirAll(filePath, os.ModePerm)
	fileContent, _ := json.MarshalIndent(tree, "", " ")
	dumpContent(fileName, fileContent)
	return
}

//Reads the given tree from its file
func ReadTree(fileName string) types.Node {
	file := loadContent(fileName)
	tree := types.Node{}
	_ = json.Unmarshal([]byte(file), &tree)
	return tree
}

//Checks if the given database exists
func CheckIfDatabaseExists(databaseName string) bool {
	directoryPath := path.Join(constants.DumpsDir, "storage", databaseName)
	return checkIfDirectoryExists(directoryPath)
}

//Creates the database
func CreateDatabase(databaseName string) error {
	directoryPath := path.Join(constants.DumpsDir, "storage", databaseName)
	return createFolder(directoryPath)
}

//Dropps the database
func DropDatabase(databaseName string) error {
	directoryPath := path.Join(constants.DumpsDir, "storage", databaseName)
	return deleteFolder(directoryPath)
}
