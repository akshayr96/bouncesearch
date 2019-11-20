package utilities

import (
	"fmt"

	"github.com/akshayr96/bounceSearch/types"
)

func ValidateSchema(schema types.Schema, record types.Record) bool {
	for field := range schema {
		fmt.Println(field)
		if record[field] == "" && !schema[field].Optional {
			return false
		} else {
			return true
		}
	}
	fmt.Println("final return")
	return false
}
