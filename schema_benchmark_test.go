// +build bench

package schema_test

import (
	"strings"
	"testing"

	"github.com/go-json-schema/schema"
	"github.com/go-json-schema/schema/draft04"
	"github.com/go-json-schema/schema/draft07"
	oldschema "github.com/lestrrat/go-jsschema"
	"github.com/lestrrat/go-jsschema/validator"
	"github.com/xeipuuv/gojsonschema"
)

func BenchmarkGojsonschemaParse(b *testing.B) {
	l := gojsonschema.NewStringLoader(draft04.MetaSchema)
	for i := 0; i < b.N; i++ {
		s, _ := gojsonschema.NewSchema(l)
		_ = s
	}
}

func BenchmarkGoJsschemaParse(b *testing.B) {
	r := strings.NewReader(draft04.MetaSchema)
	for i := 0; i < b.N; i++ {
		s, _ := oldschema.Read(r)
		_ = s
	}
}

func BenchmarkGoJsschemaParseAndMakeValidator(b *testing.B) {
	r := strings.NewReader(draft04.MetaSchema)
	for i := 0; i < b.N; i++ {
		s, _ := oldschema.Read(r)
		_ = s
		v := validator.New(s)
		v.Compile() // force compiling for comparison
	}
}

func BenchmarkGoJSONSchemaParseDraft04(b *testing.B) {
	r := strings.NewReader(draft04.MetaSchema)
	for i := 0; i < b.N; i++ {
		s, _ := schema.Parse(r, schema.WithSchemaID(draft04.SchemaID))
		_ = s
	}
}

func BenchmarkGoJSONSchemaParseDraft07(b *testing.B) {
	r := strings.NewReader(draft07.MetaSchema)
	for i := 0; i < b.N; i++ {
		s, _ := schema.Parse(r, schema.WithSchemaID(draft07.SchemaID))
		_ = s
	}
}
