package provider

import "io/ioutil"

func dumpContent(fileName string, fileContent []byte) {
	_ = ioutil.WriteFile(fileName, fileContent, 0644)
	return
}

func loadContent(fileName string) []byte {
	file, _ := ioutil.ReadFile(fileName)
	return file
}
