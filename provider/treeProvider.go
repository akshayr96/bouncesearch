package provider

import (
	"encoding/json"
	"os"
	"path"

	"github.com/akshayr96/bounceSearch/constants"
	"github.com/akshayr96/bounceSearch/types"
)

func WriteTree(tree *types.Node, databaseName string, collectionName string) {
	filePath := path.Join(constants.DumpsDir, "storage", databaseName)
	fileName := path.Join(filePath, collectionName+".json")
	os.MkdirAll(filePath, os.ModePerm)
	fileContent, _ := json.MarshalIndent(tree, "", " ")
	dumpContent(fileName, fileContent)
	return
}

func ReadTree(fileName string) types.Node {
	file := loadContent(fileName)
	tree := types.Node{}
	_ = json.Unmarshal([]byte(file), &tree)
	return tree
}
