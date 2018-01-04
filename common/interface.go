package common

type Enum interface{}
type EnumList []Enum
type Format string
type PrimitiveType string
type PrimitiveTypeList []PrimitiveType

const (
	NullType    PrimitiveType = "null"
	BooleanType               = "boolean"
	ObjectType                = "object"
	ArrayType                 = "array"
	NumberType                = "number"
	StringType                = "string"
	IntegerType               = "integer"
)


