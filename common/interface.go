package common

type Enum interface{}
type EnumList []Enum
type Format string
type PrimitiveType string
type PrimitiveTypeList []PrimitiveType

type Option interface {
	Name() string
	Value() interface{}
}

type Property interface {
	Name() string
	Definition() Schema
}

type Schema interface {
	SchemaRef() string
}

type SchemaList interface {
	Iterator() <-chan Schema
}

type SchemaSet interface {
	Iterator() <-chan *Property
}

const (
	UnspecifiedType PrimitiveType = "unspecified"
	NullType                      = "null"
	BooleanType                   = "boolean"
	ObjectType                    = "object"
	ArrayType                     = "array"
	NumberType                    = "number"
	StringType                    = "string"
	IntegerType                   = "integer"
)
