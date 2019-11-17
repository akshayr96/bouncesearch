package constants

import (
	"os"
	"path"
)

//Folder paths
var Pwd, _ = os.Getwd()
var DumpsDir = path.Join(Pwd, "dumps")

//Error messages
const (
	DBAlreadyExistsError    = "Database already exists"
	DBNotFoundError         = "No such database found"
	UnableToCreateDB        = "Unable to create database"
	ErrorDroppingDatabase   = "Error dropping the database"
	CollectionDoesntExist   = "No such collection exists"
	ErrorDeletingCollection = "Unable to delete collection"
)
