package draft07_test

import (
	"log"

	"github.com/lestrrat/go-jsschema/draft07"
)

func Example() {
	s, err := draft07.ParseFile("schema.json")
	if err != nil {
		log.Printf("failed to read schema: %s", err)
		return
	}

	for prop := range s.Properties().Iterator() {
		// Do what you will with `pdef`, which contain
		// Schema information for `name` property
		_ = prop.Name
		_ = prop.Definition
	}

	/*
		// Create a validator
		v := validator.New(s)

		// You can also validate an arbitrary piece of data
		var p interface{} // initialize using json.Unmarshal...
		if err := v.Validate(p); err != nil {
			log.Printf("failed to validate data: %s", err)
		}
	*/
}
