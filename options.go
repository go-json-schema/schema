package schema

const (
	optkeySchemaID = "schema-id"
)

type option struct {
	name string
	value interface{}
}

func (o *option) Name() string {
	return o.name
}

func (o *option) Value() interface{} {
	return o.value
}

func WithSchemaID(s string) Option {
	return &option{
		name: optkeySchemaID,
		value: s,
	}
}
