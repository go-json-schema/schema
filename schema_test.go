package schema_test

import (
	"encoding/json"
	"io"
	"strings"
	"testing"

	"github.com/go-json-schema/schema"
	"github.com/go-json-schema/schema/draft04"
	"github.com/go-json-schema/schema/draft07"
	"github.com/pmezard/go-difflib/difflib"
	"github.com/stretchr/testify/assert"
)

func TestParse(t *testing.T) {
	list := []struct {
		Name     string
		Source   io.Reader
		Expected string
		Options  []schema.Option
	}{
		{
			Name:     "draft-04",
			Source:   strings.NewReader(draft04.MetaSchema),
			Expected: draft04.MetaSchema,
			Options:  []schema.Option{schema.WithSchemaID(draft04.SchemaID)},
		},
		{
			Name:     "draft-07",
			Source:   strings.NewReader(draft07.MetaSchema),
			Expected: draft07.MetaSchema,
			Options:  []schema.Option{schema.WithSchemaID(draft07.SchemaID)},
		},
	}

	for _, data := range list {
		t.Run(data.Name, func(t *testing.T) {
			s, err := schema.Parse(data.Source, data.Options...)
			if !assert.NoError(t, err, `schema.Parse should succeed`) {
				return
			}

			buf, err := json.MarshalIndent(s, "", "  ")
			if !assert.NoError(t, err, `marshaling schema should succeed`) {
				return
			}

			if string(buf) != data.Expected {
				diff := difflib.UnifiedDiff{
					A:        difflib.SplitLines(data.Expected),
					B:        difflib.SplitLines(string(buf)),
					FromFile: "expected",
					ToFile:   "output",
				}
				result, _ := difflib.GetUnifiedDiffString(diff)
				t.Logf(result)
				t.Errorf("output does not match")
				return
			}
		})
	}
}
