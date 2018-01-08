package draft04

import (
	"encoding/json"

	"github.com/pkg/errors"
)

func (m *DependencyMap) Names() map[string][]string {
	return m.names
}

func (m *DependencyMap) Schemas() *SchemaSet {
	return m.schemas
}

func (m *DependencyMap) MarshalJSON() ([]byte, error) {
	out := make(map[string]interface{})
	for prop := range m.schemas.Iterator() {
		out[prop.Name()] = prop.Definition()
	}
	for k, v := range m.names {
		out[k] = v
	}
	return json.Marshal(out)
}

func (m *DependencyMap) UnmarshalJSON(buf []byte) error {
	var tmp map[string]json.RawMessage
	if err := json.Unmarshal(buf, &tmp); err != nil {
		return errors.Wrap(err, `failed to parse dependencies`)
	}

	depnames := make(map[string][]string)
	schemas := make(map[string]*Schema)
	for k, raw := range tmp {
		if len(raw) <= 0 {
			return errors.Errorf(`empty value found for key %s`, k)
		}

		switch raw[0] {
		case '[':
			var l []string
			if err := json.Unmarshal([]byte(raw), &l); err != nil {
				return errors.Wrap(err, `failed to unmarhsal list dependency`)
			}
			depnames[k] = l
		case '{':
			var s Schema
			if err := json.Unmarshal([]byte(raw), &s); err != nil {
				return errors.Wrap(err, `failed to unmarhsal schema dependency`)
			}
			schemas[k] = &s
		default:
			return errors.Errorf(`invalid dependency %s`, k)
		}
	}

	*m = DependencyMap{
		names:   depnames,
		schemas: &SchemaSet{store: schemas},
	}
	return nil
}
