package main

import (
	"bytes"
	"fmt"
	"go/format"
	"io"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

type Property struct {
	Name  string
	Type  string
	JSON  string
	Deref bool
	Zero  string
}

func main() {
	if err := _main(); err != nil {
		log.Printf("%s", err)
		os.Exit(1)
	}
}

func _main() error {
	if err := GenerateDraft04(); err != nil {
		return errors.Wrap(err, `failed to generate draft-04 schema`)
	}

	if err := GenerateDraft07(); err != nil {
		return errors.Wrap(err, `failed to generate draft-07 schema`)
	}
	return nil
}

func GenerateDraft04() error {
	schemaProperties := []Property{
		{
			Name:  "ID",
			Type:  "*string",
			JSON:  "id",
			Deref: true,
			Zero:  `""`,
		},
		{
			Name:  "Title",
			Type:  "*string",
			JSON:  "title",
			Deref: true,
			Zero:  `""`,
		},
		{
			Name:  "Description",
			Type:  "*string",
			JSON:  "description",
			Deref: true,
			Zero:  `""`,
		},
		{
			Name: "Default",
			Type: "interface{}",
			JSON: "default",
		},
		{
			Name: "Type",
			Type: "PrimitiveTypeList",
			JSON: "type",
		},
		{
			Name:  "SchemaRef",
			Type:  "*string",
			JSON:  "$schema",
			Deref: true,
			Zero:  "SchemaID",
		},
		{
			Name: "Definitions",
			Type: "*SchemaSet",
			JSON: "definitions",
		},
		{
			Name:  "Reference",
			Type:  "*string",
			JSON:  "$ref",
			Deref: true,
			Zero:  `""`,
		},
		{
			Name:  "Format",
			Type:  "*Format",
			JSON:  "format",
			Deref: true,
			Zero:  `Format("")`,
		},

		// NumericValidations
		{
			Name:  "MultipleOf",
			Type:  "*float64",
			JSON:  "multipleOf",
			Deref: true,
			Zero:  "float64(0)",
		},
		{
			Name:  "Minimum",
			Type:  "*float64",
			JSON:  "minimum",
			Deref: true,
			Zero:  "float64(0)",
		},
		{
			Name:  "Maximum",
			Type:  "*float64",
			JSON:  "maximum",
			Deref: true,
			Zero:  "float64(0)",
		},
		{
			Name:  "ExclusiveMinimum",
			Type:  "*bool",
			JSON:  "exclusiveMinimum",
			Deref: true,
			Zero:  "false",
		},
		{
			Name:  "ExclusiveMaximum",
			Type:  "*bool",
			JSON:  "exclusiveMaximum",
			Deref: true,
			Zero:  "false",
		},
		// StringValidation
		{
			Name:  "MaxLength",
			Type:  "*int64",
			JSON:  "maxLength",
			Deref: true,
			Zero:  "int64(0)",
		},
		{
			Name:  "MinLength",
			Type:  "*int64",
			JSON:  "minLength",
			Deref: true,
			Zero:  "int64(0)",
		},
		{
			Name: "Pattern",
			Type: "*regexp.Regexp",
			JSON: "pattern",
		},
		// ArrayValidations
		{
			Name: "AdditionalItems",
			Type: "*SchemaList",
			JSON: "additionalItems",
		},
		{
			Name: "Items",
			Type: "*Schema",
			JSON: "items",
		},
		{
			Name:  "MinItems",
			Type:  "*int64",
			JSON:  "minItems",
			Deref: true,
			Zero:  "int64(0)",
		},
		{
			Name:  "MaxItems",
			Type:  "*int64",
			JSON:  "maxItems",
			Deref: true,
			Zero:  "int64(0)",
		},
		{
			Name:  "UniqueItems",
			Type:  "*bool",
			JSON:  "uniqueItems",
			Deref: true,
			Zero:  "false",
		},

		// ObjectValidations
		{
			Name: "MaxProperties", Type: "*int64", JSON: "maxProperties",
		},
		{
			Name: "MinProperties", Type: "*int64", JSON: "minProperties",
		},
		{
			Name: "Required",
			Type: "[]string",
			JSON: "required",
		},
		{
			Name: "Dependencies",
			Type: "*DependencyMap",
			JSON: "dependencies",
		},
		{
			Name: "Properties",
			Type: "*SchemaSet",
			JSON: "properties",
		},
		{
			Name: "AdditionalProperties",
			Type: "*Schema",
			JSON: "additionalProperties",
		},
		{
			Name: "PatternProperties",
			Type: "*SchemaSet",
			JSON: "patternProperties",
		},
		{
			Name: "Enum",
			Type: "EnumList",
			JSON: "enum",
		},
		{
			Name: "AllOf",
			Type: "*SchemaList",
			JSON: "allOf",
		},
		{
			Name: "AnyOf",
			Type: "*SchemaList",
			JSON: "anyOf",
		},
		{
			Name: "OneOf",
			Type: "*SchemaList",
			JSON: "oneOf",
		},
		{
			Name: "Not",
			Type: "*Schema",
			JSON: "not",
		},
	}
	sort.Slice(schemaProperties, func(i, j int) bool {
		return schemaProperties[i].JSON < schemaProperties[j].JSON
	})

	var buf bytes.Buffer

	fmt.Fprintf(&buf, "package draft04")
	fmt.Fprintf(&buf, "\n\nimport (")
	for _, pkg := range []string{"encoding/json", "regexp"} {
		fmt.Fprintf(&buf, "\n%s", strconv.Quote(pkg))
	}
	fmt.Fprintf(&buf, "\n")
	for _, pkg := range []string{"github.com/pkg/errors"} {
		fmt.Fprintf(&buf, "\n%s", strconv.Quote(pkg))
	}
	fmt.Fprintf(&buf, "\n)") //end import

	fmt.Fprintf(&buf, "\n\ntype SchemaProperties struct {")
	for _, prop := range schemaProperties {
		fmt.Fprintf(&buf, "\n%s %s `json:\"%s,omitempty\"`", prop.Name, prop.Type, prop.JSON)
	}
	fmt.Fprintf(&buf, "\n}") // end type SchemaProperties

	if err := WriteGetters(&buf, schemaProperties); err != nil {
		return errors.Wrap(err, `failed to write getters for draft04`)
	}

	fmt.Fprintf(&buf, "\n\nfunc (s *Schema) MarshalJSON() ([]byte, error) {")
	fmt.Fprintf(&buf, "\nreturn json.Marshal(s.properties)")
	fmt.Fprintf(&buf, "\n}") // end MarshalJSON

	fmt.Fprintf(&buf, "\n\nfunc (s *Schema) UnmarshalJSON(buf []byte) error {")
	fmt.Fprintf(&buf, "\nvar props SchemaProperties")
	fmt.Fprintf(&buf, "\nif err := json.Unmarshal(buf, &props); err != nil {")
	fmt.Fprintf(&buf, "\nreturn errors.Wrap(err, `failed to unmarshal schema`)")
	fmt.Fprintf(&buf, "\n}")
	fmt.Fprintf(&buf, "\ns.properties = &props")
	fmt.Fprintf(&buf, "\nreturn nil")
	fmt.Fprintf(&buf, "\n}") // end MarshalJSON

	WriteSchemaHelpers(&buf)
	return WriteFormattedSource(`draft04/schema_gen.go`, &buf)
}

func GenerateDraft07() error {
	var schemaProperties = []Property{
		{
			Name:  "SchemaRef",
			Type:  "*string",
			JSON:  "$schema",
			Deref: true,
			Zero:  "SchemaID",
		},
		{
			Name:  "Reference",
			Type:  "*string",
			JSON:  "$ref",
			Deref: true,
			Zero:  `""`,
		},

		{
			Name:  "ID",
			Type:  "*string",
			JSON:  "$id",
			Deref: true,
			Zero:  `""`,
		},
		{
			Name:  "Comment",
			Type:  "*string",
			JSON:  "$comment",
			Deref: true,
			Zero:  `""`,
		},
		{
			Name:  "Title",
			Type:  "*string",
			JSON:  "title",
			Deref: true,
			Zero:  `""`,
		},
		{
			Name:  "Description",
			Type:  "*string",
			JSON:  "description",
			Deref: true,
			Zero:  `""`,
		},
		{
			Name: "Definitions",
			Type: "*SchemaSet",
			JSON: "definitions",
		},
		{
			Name: "Default",
			Type: "interface{}",
			JSON: "default",
		},

		// 6.1.  Validation Keywords for Any Instance Type
		{
			Name: "Type",
			Type: "PrimitiveTypeList",
			JSON: "type",
		},
		{
			Name: "Enum",
			Type: "EnumList",
			JSON: "enum",
		},
		{
			Name: "Const",
			Type: "interface{}",
			JSON: "const",
		},

		// 6.2.  Validation Keywords for Numeric Instances (number and integer)
		{
			Name:  "MultipleOf",
			Type:  "*float64",
			JSON:  "multipleOf",
			Deref: true,
			Zero:  "float64(0)",
		},
		{
			Name:  "Maximum",
			Type:  "*float64",
			JSON:  "maximum",
			Deref: true,
			Zero:  "float64(0)",
		},
		{
			Name:  "Minimum",
			Type:  "*float64",
			JSON:  "minimum",
			Deref: true,
			Zero:  "float64(0)",
		},
		{
			Name:  "ExclusiveMaximum",
			Type:  "*float64",
			JSON:  "exclusiveMaximum",
			Deref: true,
			Zero:  "float64(0)",
		},
		{
			Name:  "ExclusiveMinimum",
			Type:  "*float64",
			JSON:  "exclusiveMinimum",
			Deref: true,
			Zero:  "float64(0)",
		},

		// 6.3.  Validation Keywords for Strings
		{
			Name:  "MaxLength",
			Type:  "*int64",
			JSON:  "maxLength",
			Deref: true,
			Zero:  "int64(0)",
		},
		{
			Name:  "MinLength",
			Type:  "*int64",
			JSON:  "minLength",
			Deref: true,
			Zero:  "int64(0)",
		},
		{
			Name: "Pattern",
			Type: "*regexp.Regexp",
			JSON: "pattern",
		},

		// 6.4.  Validation Keywords for Arrays
		{
			Name: "Items",
			Type: "*SchemaList",
			JSON: "items",
		},
		{
			Name: "AdditionalItems",
			Type: "*Schema",
			JSON: "additionalItems",
		},
		{
			Name:  "MaxItems",
			Type:  "*int64",
			JSON:  "maxItems",
			Deref: true,
			Zero:  "int64(0)",
		},
		{
			Name:  "MinItems",
			Type:  "*int64",
			JSON:  "minItems",
			Deref: true,
			Zero:  "int64(0)",
		},
		{
			Name:  "UniqueItems",
			Type:  "*bool",
			JSON:  "uniqueItems",
			Deref: true,
			Zero:  "false",
		},
		{
			Name: "Contains",
			Type: "*SchemaList",
			JSON: "contains",
		},

		// 6.5.  Validation Keywords for Objects
		{
			Name:  "MaxProperties",
			Type:  "*int64",
			JSON:  "maxProperties",
			Deref: true,
			Zero:  "int64(0)",
		},
		{
			Name:  "MinProperties",
			Type:  "*int64",
			JSON:  "minProperties",
			Deref: true,
			Zero:  "int64(0)",
		},
		{
			Name: "Required",
			Type: "[]string",
			JSON: "required",
		},
		{
			Name: "Properties",
			Type: "*SchemaSet",
			JSON: "properties",
		},
		{
			Name: "PatternProperties",
			Type: "*SchemaSet",
			JSON: "patternProperties",
		},
		{
			Name: "AdditionalProperties",
			Type: "*Schema",
			JSON: "additionalProperties",
		},
		{
			Name: "Dependencies",
			Type: "*SchemaSet",
			JSON: "dependencies",
		},
		{
			Name: "PropertyNames",
			Type: "*Schema",
			JSON: "propertyNames",
		},

		// 6.6.  Keywords for Applying Subschemas Conditionally
		{
			Name: "If",
			Type: "*Schema",
			JSON: "if",
		},
		{
			Name: "Then",
			Type: "*Schema",
			JSON: "then",
		},
		{
			Name: "Else",
			Type: "*Schema",
			JSON: "else",
		},

		// 6.7.  Keywords for Applying Subschemas With Boolean Logic
		{
			Name: "AllOf",
			Type: "*SchemaList",
			JSON: "allOf",
		},
		{
			Name: "AnyOf",
			Type: "*SchemaList",
			JSON: "anyOf",
		},
		{
			Name: "OneOf",
			Type: "*SchemaList",
			JSON: "oneOf",
		},
		{
			Name: "Not",
			Type: "*Schema",
			JSON: "not",
		},

		// 7.  Semantic Validation With "format"
		{
			Name:  "Format",
			Type:  "*Format",
			JSON:  "format",
			Deref: true,
			Zero:  `Format("")`,
		},
	}

	sort.Slice(schemaProperties, func(i, j int) bool {
		return schemaProperties[i].JSON < schemaProperties[j].JSON
	})

	var buf bytes.Buffer

	fmt.Fprintf(&buf, "package draft07")
	fmt.Fprintf(&buf, "\n\nimport (")
	for _, pkg := range []string{"bytes", "encoding/json", "regexp"} {
		fmt.Fprintf(&buf, "\n%s", strconv.Quote(pkg))
	}
	fmt.Fprintf(&buf, "\n")
	for _, pkg := range []string{"github.com/pkg/errors"} {
		fmt.Fprintf(&buf, "\n%s", strconv.Quote(pkg))
	}
	fmt.Fprintf(&buf, "\n)") //end import

	fmt.Fprintf(&buf, "\n\ntype SchemaProperties struct {")
	for _, prop := range schemaProperties {
		fmt.Fprintf(&buf, "\n%s %s `json:\"%s,omitempty\"`", prop.Name, prop.Type, prop.JSON)
	}
	fmt.Fprintf(&buf, "\n}") // end type SchemaProperties

	if err := WriteGetters(&buf, schemaProperties); err != nil {
		return errors.Wrap(err, `failed to write getters for draft07`)
	}

	for _, prop := range schemaProperties {
		if strings.HasPrefix(prop.Type, `[]`) {
			fmt.Fprintf(&buf, "\n\nfunc (s *Schema) Add%s(list ...%s) *Schema {", prop.Name, prop.Type[2:])
			fmt.Fprintf(&buf, "\ns.properties.%s = append(s.properties.%s, list...)", prop.Name, prop.Name)
			fmt.Fprintf(&buf, "\nreturn s")
			fmt.Fprintf(&buf, "\n}") // end setter
		} else if strings.HasSuffix(prop.Type, `List`) {
			fmt.Fprintf(&buf, "\n\nfunc (s *Schema) Add%s(list ...%s) *Schema {", prop.Name, prop.Type[:len(prop.Type)-4])
			fmt.Fprintf(&buf, "\ns.properties.%s.Append(list...)", prop.Name)
			fmt.Fprintf(&buf, "\nreturn s")
			fmt.Fprintf(&buf, "\n}") // end setter
		} else {
			setType := prop.Type
			setValue := "v"
			switch setType {
			case "*int64", "*float64", "*bool":
				setType = setType[1:]
				setValue = "&v"
			}

			fmt.Fprintf(&buf, "\n\nfunc(s *Schema) Set%s(v %s) *Schema {", prop.Name, setType)
			fmt.Fprintf(&buf, "\ns.properties.%s = %s", prop.Name, setValue)
			fmt.Fprintf(&buf, "\nreturn s")
			fmt.Fprintf(&buf, "\n}") // end setter
		}
	}

	fmt.Fprintf(&buf, "\n\nvar trueBytes = []byte(\"true\")")
	fmt.Fprintf(&buf, "\nvar falseBytes = []byte(\"false\")")
	fmt.Fprintf(&buf, "\n\nfunc (s *Schema) MarshalJSON() ([]byte, error) {")
	fmt.Fprintf(&buf, "\nif s.isNegated {")
	fmt.Fprintf(&buf, "\nreturn falseBytes, nil")
	fmt.Fprintf(&buf, "\n}") // end isNegated
	fmt.Fprintf(&buf, "\n\nif s.isEmpty {")
	fmt.Fprintf(&buf, "\nreturn trueBytes, nil")
	fmt.Fprintf(&buf, "\n}") // end isEmpty
	fmt.Fprintf(&buf, "\n\nreturn json.Marshal(s.properties)")
	fmt.Fprintf(&buf, "\n}") // end MarshalJSON

	fmt.Fprintf(&buf, "\n\nfunc (s *Schema) UnmarshalJSON(buf []byte) error {")
	fmt.Fprintf(&buf, "\nif bytes.HasPrefix(buf, trueBytes) {")
	fmt.Fprintf(&buf, "\n*s = Schema{isEmpty: true}")
	fmt.Fprintf(&buf, "\nreturn nil")
	fmt.Fprintf(&buf, "\n} else if bytes.HasPrefix(buf, falseBytes) {")
	fmt.Fprintf(&buf, "\n*s = Schema{isNegated: true}")
	fmt.Fprintf(&buf, "\nreturn nil")
	fmt.Fprintf(&buf, "\n}")
	fmt.Fprintf(&buf, "\n\nvar fs Schema")
	fmt.Fprintf(&buf, "\nif err := json.Unmarshal(buf, &fs.properties); err != nil {")
	fmt.Fprintf(&buf, "\nreturn errors.Wrap(err, \"failed to unmarshal schema\")")
	fmt.Fprintf(&buf, "\n}")
	fmt.Fprintf(&buf, "\n*s = fs")
	fmt.Fprintf(&buf, "\nreturn nil")
	fmt.Fprintf(&buf, "\n}")

	WriteSchemaHelpers(&buf)
	return WriteFormattedSource(`draft07/schema_gen.go`, &buf)
}

func WriteFormattedSource(file string, buf *bytes.Buffer) error {
	formatted, err := format.Source(buf.Bytes())
	if err != nil {
		log.Printf("%s", buf.Bytes())
		return errors.Wrap(err, `failed to format source`)
	}

	fh, err := os.Create(file)
	if err != nil {
		return errors.Wrapf(err, `failed to open file %s for writing`, file)
	}
	defer fh.Close()

	fh.Write(formatted)
	return nil
}

func WriteSchemaHelpers(dst io.Writer) {
	fmt.Fprintf(dst, `
func (s *SchemaSet) Iterator() <-chan *Property {
	s.mu.RLock()
	ch := make(chan *Property, len(s.store))
	go func() {
		defer s.mu.RUnlock()
		defer close(ch)
		for k, v := range s.store {
			ch <- &Property{Name: k, Definition: v}
		}
	}()
	return ch
}

func (s *SchemaSet) MarshalJSON() ([]byte, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return json.Marshal(s.store)
}

func (s *SchemaSet) UnmarshalJSON(buf []byte) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	return json.Unmarshal(buf, &s.store)
}

func (s *SchemaList) Append(list ...*Schema) {
	s.store = append(s.store, list...)
}

func (s *SchemaList) MarshalJSON() ([]byte, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	if len(s.store) == 1 {
		return json.Marshal(s.store[0])
	}
	return json.Marshal(s.store)
}

func (s *SchemaList) UnmarshalJSON(buf []byte) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	if len(buf) > 0 && buf[0] == '[' {
		return json.Unmarshal(buf, &s.store)
	}
	var v Schema
	if err := json.Unmarshal(buf, &v); err != nil {
		return errors.Wrap(err, "failed to unmarshal json schema list")
	}
	s.store = []*Schema{&v}
	return nil
}`)
}

func WriteGetters(dst io.Writer, properties []Property) error {
	for _, prop := range properties {
		typ := prop.Type
		if prop.Deref {
			typ = strings.TrimPrefix(typ, "*")
		}

		fmt.Fprintf(dst, "\n\nfunc(s *Schema) %s() %s {", prop.Name, typ)
		if prop.Deref {
			if prop.Zero == "" {
				return errors.Errorf(`property %s needs dereferencing, but does not have a zero value for fallback`)
			}
			fmt.Fprintf(dst, "\nif !s.Has%s() {", prop.Name)
			fmt.Fprintf(dst, "\nreturn %s", prop.Zero)
			fmt.Fprintf(dst, "\n}") // end if !s.Has%s

			fmt.Fprintf(dst, "\nreturn *(s.properties.%s)", prop.Name)
		} else {
			fmt.Fprintf(dst, "\nreturn s.properties.%s", prop.Name)
		}
		fmt.Fprintf(dst, "\n}") // end getter

		// we need to create a HasFoo() method
		fmt.Fprintf(dst, "\n\nfunc (s *Schema) Has%s() bool {", prop.Name)
		fmt.Fprintf(dst, "\nreturn s.properties.%s != nil", prop.Name)
		fmt.Fprintf(dst, "\n}") // end Has%s
	}
	return nil
}
