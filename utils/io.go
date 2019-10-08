package utils

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path"

	"github.com/akshayr96/bounceSearch/constants"

	"github.com/akshayr96/bounceSearch/types"
)

func WriteTree(tree *types.Node, databaseName string, collectionName string) {
	filePath := path.Join(constants.DumpsDir, databaseName)
	fileName := path.Join(filePath, collectionName+".json")
	os.MkdirAll(filePath, os.ModePerm)
	fileContent, _ := json.MarshalIndent(tree, "", " ")
	_ = ioutil.WriteFile(fileName, fileContent, 0644)
	return
}

func ReadTree(fileName string) types.Node {
	file, _ := ioutil.ReadFile(fileName)
	tree := types.Node{}
	_ = json.Unmarshal([]byte(file), &tree)
	return tree
}
