package schema

import "github.com/go-json-schema/schema/common"

type Option = common.Option

type Schema interface {
	SchemaRef() string
}
