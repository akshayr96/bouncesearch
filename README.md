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

# Getting Started

**Database API**

Create database
```go
package main

import "github.com/akshayr96/bounceSearch"

//create a database
err := bounceSearch.Create("rick_and_morty")
```

Connect to a database
```go
if err == nil  {
  //connect to the database
  conn, err := bounceSearch.Connect("rick_and_morty")
}

```

Drop database
```go
err := database.Drop("rick_and_morty")
```



**Schema Definition**
```go
//create a schema
charactersSchema := bounceSearch.types.Schema{
  "name": bounceSearch.types.AttributeMeta{defaultValue: "", Weight: 1.0, Optional: false},
  "description":  bounceSearch.AttributeMeta{defaultValue: "", Weight: 1.0, Optional: false},
}
 ```



**Collections API**
Create collection
```go
characters := conn.Collection("characters", charactersSchema)
```

Delete collection
```go
err = conn.Delete("characters")
```

# Tasks to be completed

- [x] Create tree generating and parsing functions
- [x] Add Levenshtein Distance computing function
- [x] Add Unit Test Cases
- [x] Implement the File write Layer to persist data 
- [ ] Create the handlers for indexing and searching
- [ ] Create REST APIs to accept queries
- [ ] Add Message Queue to queue the request