package database

import (
	"errors"

	"github.com/akshayr96/bounceSearch/constants"
	"github.com/akshayr96/bounceSearch/provider"
)

// @todo: add validations to database names in all functions

// Created the database
// Takes the database name and returns the error message if any
func Create(databaseName string) error {
	if provider.CheckIfDatabaseExists(databaseName) {
		return errors.New(constants.DBAlreadyExistsError)
	} else {
		err := provider.CreateDatabase(databaseName)
		if err != nil {
			return errors.New(constants.UnableToCreateDB)
		}
		return nil
	}
}

//Creates a connection to a database
//Takes the database name and returns connection object or error message
func Connect(databaseName string) (Collections, error) {
	if provider.CheckIfDatabaseExists(databaseName) {
		return Collections{databaseName}, nil
	} else {
		return Collections{""}, errors.New(constants.DBNotFoundError)
	}
}

//Drops a database
//Takes the database name and returns error if any
func Drop(databaseName string) error {
	if provider.CheckIfDatabaseExists(databaseName) {
		err := provider.DropDatabase(databaseName)
		if err != nil {
			return errors.New(constants.ErrorDroppingDatabase)
		}
		return nil
	}
	return errors.New(constants.DBNotFoundError)
}
