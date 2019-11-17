package provider

import (
	"io/ioutil"
	"os"
	"path/filepath"
)

//Writes the contents of a file
func dumpContent(fileName string, fileContent []byte) {
	_ = ioutil.WriteFile(fileName, fileContent, 0644)
	return
}

//Reads the contents of a file
func loadContent(fileName string) []byte {
	file, _ := ioutil.ReadFile(fileName)
	return file
}

//Checks if a given directory exists
func checkIfDirectoryExists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return true
}

//Creates a folder in the given name
func createFolder(path string) error {
	err := os.Mkdir(path, 'd')
	if err != nil {
		return err
	}
	return nil
}

//Deletes a given folder
func deleteFolder(path string) error {
	directory, err := os.Open(path)
	if err != nil {
		return err
	}
	defer directory.Close()
	names, err := directory.Readdirnames(-1)
	if err != nil {
		return err
	}
	for _, name := range names {
		err = os.RemoveAll(filepath.Join(path, name))
		if err != nil {
			return err
		}
	}
	err = os.Remove(path)
	if err != nil {
		return err
	}
	return nil
}

func deleteFile(path string) error {
	err := os.Remove(path)
	if err != nil {
		return err
	}
	return nil
}
