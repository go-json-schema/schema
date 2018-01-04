package schema

type Option interface {
	Name() string
	Value() interface{}
}

type Schema interface {
}
