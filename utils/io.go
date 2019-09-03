package utils

import (
	"encoding/json"
	"io/ioutil"

	"github.com/akshayr96/bounceSearch/types"
)

func WriteTree(tree *types.Node, fileName string) {
	file, _ := json.MarshalIndent(tree, "", " ")
	_ = ioutil.WriteFile(fileName, file, 0644)
	return
}

func ReadTree(fileName string) types.Node {
	file, _ := ioutil.ReadFile(fileName)
	tree := types.Node{}
	_ = json.Unmarshal([]byte(file), &tree)
	return tree
}
