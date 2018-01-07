package draft04

import (
	"encoding/json"
	"io"
	"os"

	"github.com/go-json-schema/schema/common"
	"github.com/pkg/errors"
)

func ParseFile(fn string) (*Schema, error) {
	f, err := os.Open(fn)
	if err != nil {
		return nil, errors.Wrap(err, `failed to open schema file`)
	}
	defer f.Close()

	return Parse(f)
}

func Parse(src io.Reader, options ...Option) (*Schema, error) {
	var s Schema
	if err := json.NewDecoder(src).Decode(&s); err != nil {
		return nil, errors.Wrap(err, `failed to unmarshal schema`)
	}
	return &s, nil
}

func (p *Property) Name() string {
	return p.name
}

func (p *Property) Definition() common.Schema {
	return p.definition
}
