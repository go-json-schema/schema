package draft07

// References:
// https://tools.ietf.org/html/draft-handrews-json-schema-validation-00

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
	NullType    = common.NullType
	BooleanType = common.BooleanType
	ObjectType  = common.ObjectType
	ArrayType   = common.ArrayType
	NumberType  = common.NumberType
	StringType  = common.StringType
	IntegerType = common.IntegerType
)

const SchemaID = "http://json-schema.org/draft-07/schema#"
const MetaSchema = `{
  "$id": "http://json-schema.org/draft-07/schema#",
  "$schema": "http://json-schema.org/draft-07/schema#",
  "default": true,
  "definitions": {
    "nonNegativeInteger": {
      "minimum": 0,
      "type": "integer"
    },
    "nonNegativeIntegerDefault0": {
      "allOf": [
        {
          "$ref": "#/definitions/nonNegativeInteger"
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
      "default": [],
      "items": {
        "type": "string"
      },
      "type": "array",
      "uniqueItems": true
    }
  },
  "properties": {
    "$comment": {
      "type": "string"
    },
    "$id": {
      "format": "uri-reference",
      "type": "string"
    },
    "$ref": {
      "format": "uri-reference",
      "type": "string"
    },
    "$schema": {
      "format": "uri",
      "type": "string"
    },
    "additionalItems": {
      "$ref": "#"
    },
    "additionalProperties": {
      "$ref": "#"
    },
    "allOf": {
      "$ref": "#/definitions/schemaArray"
    },
    "anyOf": {
      "$ref": "#/definitions/schemaArray"
    },
    "const": true,
    "contains": {
      "$ref": "#"
    },
    "contentEncoding": {
      "type": "string"
    },
    "contentMediaType": {
      "type": "string"
    },
    "default": true,
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
    "else": {
      "$ref": "#"
    },
    "enum": {
      "items": true,
      "minItems": 1,
      "type": "array",
      "uniqueItems": true
    },
    "examples": {
      "items": true,
      "type": "array"
    },
    "exclusiveMaximum": {
      "type": "number"
    },
    "exclusiveMinimum": {
      "type": "number"
    },
    "format": {
      "type": "string"
    },
    "if": {
      "$ref": "#"
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
      "default": true
    },
    "maxItems": {
      "$ref": "#/definitions/nonNegativeInteger"
    },
    "maxLength": {
      "$ref": "#/definitions/nonNegativeInteger"
    },
    "maxProperties": {
      "$ref": "#/definitions/nonNegativeInteger"
    },
    "maximum": {
      "type": "number"
    },
    "minItems": {
      "$ref": "#/definitions/nonNegativeIntegerDefault0"
    },
    "minLength": {
      "$ref": "#/definitions/nonNegativeIntegerDefault0"
    },
    "minProperties": {
      "$ref": "#/definitions/nonNegativeIntegerDefault0"
    },
    "minimum": {
      "type": "number"
    },
    "multipleOf": {
      "exclusiveMinimum": 0,
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
      "propertyNames": {
        "format": "regex"
      },
      "type": "object"
    },
    "properties": {
      "additionalProperties": {
        "$ref": "#"
      },
      "default": {},
      "type": "object"
    },
    "propertyNames": {
      "$ref": "#"
    },
    "readOnly": {
      "default": false,
      "type": "boolean"
    },
    "required": {
      "$ref": "#/definitions/stringArray"
    },
    "then": {
      "$ref": "#"
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
  "title": "Core schema meta-schema",
  "type": [
    "object",
    "boolean"
  ]
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
	isEmpty    bool
	isNegated  bool
	properties *SchemaProperties
}

type Property struct {
	Name       string
	Definition *Schema
}
