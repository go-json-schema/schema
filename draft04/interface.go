package draft04

import (
	"sync"

	"github.com/go-json-schema/schema/common"
)

type Enum = common.Enum
type EnumList = common.EnumList
type Format = common.Format
type Option = common.Option
type PrimitiveType = common.PrimitiveType
type PrimitiveTypeList = common.PrimitiveTypeList

const (
	UnspecifiedType = common.UnspecifiedType
	NullType        = common.NullType
	BooleanType     = common.BooleanType
	ObjectType      = common.ObjectType
	ArrayType       = common.ArrayType
	NumberType      = common.NumberType
	StringType      = common.StringType
	IntegerType     = common.IntegerType
)

const SchemaID = "http://json-schema.org/draft-04/schema#"
const MetaSchema = `{
  "$schema": "http://json-schema.org/draft-04/schema#",
  "default": {},
  "definitions": {
    "positiveInteger": {
      "minimum": 0,
      "type": "integer"
    },
    "positiveIntegerDefault0": {
      "allOf": [
        {
          "$ref": "#/definitions/positiveInteger"
        },
        {
          "default": 0
        }
      ]
    },
    "schemaArray": {
      "items": {
        "$ref": "#"
      },
      "minItems": 1,
      "type": "array"
    },
    "simpleTypes": {
      "enum": [
        "array",
        "boolean",
        "integer",
        "null",
        "number",
        "object",
        "string"
      ]
    },
    "stringArray": {
      "items": {
        "type": "string"
      },
      "minItems": 1,
      "type": "array",
      "uniqueItems": true
    }
  },
  "dependencies": {
    "exclusiveMaximum": [
      "maximum"
    ],
    "exclusiveMinimum": [
      "minimum"
    ]
  },
  "description": "Core schema meta-schema",
  "id": "http://json-schema.org/draft-04/schema#",
  "properties": {
    "$schema": {
      "format": "uri",
      "type": "string"
    },
    "additionalItems": {
      "anyOf": [
        {
          "type": "boolean"
        },
        {
          "$ref": "#"
        }
      ],
      "default": {}
    },
    "additionalProperties": {
      "anyOf": [
        {
          "type": "boolean"
        },
        {
          "$ref": "#"
        }
      ],
      "default": {}
    },
    "allOf": {
      "$ref": "#/definitions/schemaArray"
    },
    "anyOf": {
      "$ref": "#/definitions/schemaArray"
    },
    "default": {},
    "definitions": {
      "additionalProperties": {
        "$ref": "#"
      },
      "default": {},
      "type": "object"
    },
    "dependencies": {
      "additionalProperties": {
        "anyOf": [
          {
            "$ref": "#"
          },
          {
            "$ref": "#/definitions/stringArray"
          }
        ]
      },
      "type": "object"
    },
    "description": {
      "type": "string"
    },
    "enum": {
      "minItems": 1,
      "type": "array",
      "uniqueItems": true
    },
    "exclusiveMaximum": {
      "default": false,
      "type": "boolean"
    },
    "exclusiveMinimum": {
      "default": false,
      "type": "boolean"
    },
    "id": {
      "format": "uri",
      "type": "string"
    },
    "items": {
      "anyOf": [
        {
          "$ref": "#"
        },
        {
          "$ref": "#/definitions/schemaArray"
        }
      ],
      "default": {}
    },
    "maxItems": {
      "$ref": "#/definitions/positiveInteger"
    },
    "maxLength": {
      "$ref": "#/definitions/positiveInteger"
    },
    "maxProperties": {
      "$ref": "#/definitions/positiveInteger"
    },
    "maximum": {
      "type": "number"
    },
    "minItems": {
      "$ref": "#/definitions/positiveIntegerDefault0"
    },
    "minLength": {
      "$ref": "#/definitions/positiveIntegerDefault0"
    },
    "minProperties": {
      "$ref": "#/definitions/positiveIntegerDefault0"
    },
    "minimum": {
      "type": "number"
    },
    "multipleOf": {
      "exclusiveMinimum": true,
      "minimum": 0,
      "type": "number"
    },
    "not": {
      "$ref": "#"
    },
    "oneOf": {
      "$ref": "#/definitions/schemaArray"
    },
    "pattern": {
      "format": "regex",
      "type": "string"
    },
    "patternProperties": {
      "additionalProperties": {
        "$ref": "#"
      },
      "default": {},
      "type": "object"
    },
    "properties": {
      "additionalProperties": {
        "$ref": "#"
      },
      "default": {},
      "type": "object"
    },
    "required": {
      "$ref": "#/definitions/stringArray"
    },
    "title": {
      "type": "string"
    },
    "type": {
      "anyOf": [
        {
          "$ref": "#/definitions/simpleTypes"
        },
        {
          "items": {
            "$ref": "#/definitions/simpleTypes"
          },
          "minItems": 1,
          "type": "array",
          "uniqueItems": true
        }
      ]
    },
    "uniqueItems": {
      "default": false,
      "type": "boolean"
    }
  },
  "type": "object"
}`

type SchemaSet struct {
	mu    sync.RWMutex
	store map[string]*Schema
}

type SchemaList struct {
	mu    sync.RWMutex
	store []*Schema
}

type Schema struct {
	properties *SchemaProperties
}

type Property struct {
	Name       string
	Definition *Schema
}

// DependencyMap contains the dependencies defined within this schema.
// for a given dependency name, you can have either a schema or a
// list of property names
type DependencyMap struct {
	names   map[string][]string
	schemas *SchemaSet
}
