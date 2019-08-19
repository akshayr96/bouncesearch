package database

func Create(databaseName string) {

}

func Connect(databaseName string) Collections {
	if true { // check if such a collection exists
		return Collections{databaseName}
	} else {
		//return error
		return Collections{"Error"}
	}
}

func Drop(databaseName string) {

}
