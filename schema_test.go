package schema_test

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/go-json-schema/schema"
	"github.com/go-json-schema/schema/draft04"
	"github.com/go-json-schema/schema/draft07"
	"github.com/pmezard/go-difflib/difflib"
	"github.com/stretchr/testify/assert"
)

func TestParseDraft07(t *testing.T) {
	s, err := schema.Parse(strings.NewReader(draft07.MetaSchema), schema.WithSchemaID(draft07.SchemaID))
	if !assert.NoError(t, err, `schema.Parse should succeed`) {
		return
	}

	buf, err := json.MarshalIndent(s, "", "  ")
	if !assert.NoError(t, err, `marshaling schema should succeed`) {
		return
	}

	if string(buf) != draft07.MetaSchema {
		diff := difflib.UnifiedDiff{
			A:        difflib.SplitLines(draft07.MetaSchema),
			B:        difflib.SplitLines(string(buf)),
			FromFile: "expected",
			ToFile:   "output",
		}
		result, _ := difflib.GetUnifiedDiffString(diff)
		t.Logf(result)
		t.Errorf("output does not match")
		return
	}
}

func TestParseDraft04(t *testing.T) {
	s, err := schema.Parse(strings.NewReader(draft04.MetaSchema), schema.WithSchemaID(draft04.SchemaID))
	if !assert.NoError(t, err, `schema.Parse should succeed`) {
		return
	}

	buf, err := json.MarshalIndent(s, "", "  ")
	if !assert.NoError(t, err, `marshaling schema should succeed`) {
		return
	}

	if string(buf) != draft04.MetaSchema {
		diff := difflib.UnifiedDiff{
			A:        difflib.SplitLines(draft04.MetaSchema),
			B:        difflib.SplitLines(string(buf)),
			FromFile: "expected",
			ToFile:   "output",
		}
		result, _ := difflib.GetUnifiedDiffString(diff)
		t.Logf(result)
		t.Errorf("output does not match")
		return
	}
}
