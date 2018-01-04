//go:generate go run internal/cmd/genschema/genschema.go

package schema

import (
	"io"

	"github.com/go-json-schema/schema/draft04"
	"github.com/go-json-schema/schema/draft07"
	"github.com/pkg/errors"
)

func toDraft04Options(options ...Option) []draft04.Option {
	list := make([]draft04.Option, len(options))
	for i, o := range options {
		list[i] = o
	}
	return list
}

func toDraft07Options(options ...Option) []draft07.Option {
	list := make([]draft07.Option, len(options))
	for i, o := range options {
		list[i] = o
	}
	return list
}

func Parse(src io.Reader, options ...Option) (Schema, error) {
	draftVersion := draft07.SchemaID

	for _, o := range options {
		switch o.Name() {
		case optkeySchemaID:
			draftVersion = o.Value().(string)
		}
	}

	switch draftVersion {
	case draft04.SchemaID:
		return draft04.Parse(src, toDraft04Options(options...)...)
	case draft07.SchemaID:
		return draft07.Parse(src, toDraft07Options(options...)...)
	default:
		return nil, errors.Errorf(`failed to parse: unsupported schema id %s`, draftVersion)
	}
}
