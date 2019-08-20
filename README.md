# Bounce Search

It is Golang based search database. I am developing it just to learn and understand Golang and the internals of a database more, it is not meant to be used in production.


# Jargon

Just for the sake of simplicity, the terminologies are made much similar to MongoDB

* **Document** - A group of key value pairs of a predefined keynames and value types
* **Collection** - A group of homogenious documents
* **Database** - Duh


# Scope of BS

* Index documents in a collection of a predefined schema with weights into a ternary search tree 

* On search return either top scored paginated list of collection that matched the search phrase or suggested similar search phrase with the closest levestian distance in the collection if no results are found,


# Tasks to be completed

- [x] Create tree generating and parsing functions
- [x] Add Levenshtein Distance computing function
- [ ] Create the API handler methods
- [ ] Implement the File write Layer to persist data 
- [ ] Create REST APIs to accept queries
- [ ] Add Message Queue to queue the request
- [ ] Add Unit Test Cases