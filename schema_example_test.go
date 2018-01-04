package schema_test

import (
	"encoding/json"
	"log"
	"strings"

	"github.com/go-json-schema/schema"
	"github.com/go-json-schema/schema/draft04"
	"github.com/go-json-schema/schema/draft07"
)

func Example() {
	// Parse the draft-07 schema
	s, err := schema.Parse(strings.NewReader(draft07.MetaSchema))
	if err != nil {
		log.Printf("failed to parse schema: %s", err)
		return
	}

	// Check if this is a draft-04 or draft-07 schema
	switch s.SchemaRef() {
	case draft04.SchemaID:
		// do draft-04 specific stuff...
	case draft07.SchemaID:
		// do draft-07 specific stuff...
	default:
		log.Printf("unknown schema")
	}

	// Marshal it back to JSON
	buf, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		log.Printf("failed to marshal schema: %s", err)
		return
	}
	_ = buf
	// OUTPUT:
}
