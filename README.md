# github.com/go-json-schema/schema

JSON Schema for Go: handles definition, and marshaling/unmarshaling JSON schemas

[![Build Status](https://travis-ci.org/go-json-schema/schema.png?branch=master)](https://travis-ci.org/go-json-schema/schema)

[![GoDoc](https://godoc.org/github.com/go-json-schema/schema?status.svg)](https://godoc.org/github.com/go-json-schema/schema)
# SYNOPSIS

```go
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
    // s.(*draft04.Schema)
  case draft07.SchemaID:
    // do draft-07 specific stuff...
    // s.(*draft07.Schema)
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
```

# DESCRIPTION

This library can parse various JSON schema versions into a
struct, which you can then use to manipulate, such as to
build a validator.

# SUPPORTED VERSIONS

* draft-04
* draft-07

