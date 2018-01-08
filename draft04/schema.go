package draft04

import (
	"encoding/json"
	"io"
	"os"

	pdebug "github.com/lestrrat/go-pdebug"
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

func Parse(src io.Reader, options ...Option) (s *Schema, err error) {
	if pdebug.Enabled {
		g := pdebug.Marker("draft04.Parse").BindError(&err)
		defer g.End()
	}

	var s1 Schema
	if err := json.NewDecoder(src).Decode(&s1); err != nil {
		return nil, errors.Wrap(err, `failed to unmarshal schema`)
	}
	return &s1, nil
}
