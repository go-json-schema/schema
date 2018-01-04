package draft07

import (
	"bytes"
	"encoding/json"
	"regexp"

	"github.com/pkg/errors"
)

type SchemaProperties struct {
	Comment              string            `json:"$comment,omitempty"`
	ID                   string            `json:"$id,omitempty"`
	Reference            string            `json:"$ref,omitempty"`
	SchemaRef            string            `json:"$schema,omitempty"`
	AdditionalItems      *Schema           `json:"additionalItems,omitempty"`
	AdditionalProperties *Schema           `json:"additionalProperties,omitempty"`
	AllOf                *SchemaList       `json:"allOf,omitempty"`
	AnyOf                *SchemaList       `json:"anyOf,omitempty"`
	Const                interface{}       `json:"const,omitempty"`
	Contains             *SchemaList       `json:"contains,omitempty"`
	Default              interface{}       `json:"default,omitempty"`
	Definitions          *SchemaSet        `json:"definitions,omitempty"`
	Dependencies         *SchemaSet        `json:"dependencies,omitempty"`
	Description          string            `json:"description,omitempty"`
	Else                 *Schema           `json:"else,omitempty"`
	Enum                 EnumList          `json:"enum,omitempty"`
	ExclusiveMaximum     *float64          `json:"exclusiveMaximum,omitempty"`
	ExclusiveMinimum     *float64          `json:"exclusiveMinimum,omitempty"`
	Format               Format            `json:"format,omitempty"`
	If                   *Schema           `json:"if,omitempty"`
	Items                *SchemaList       `json:"items,omitempty"`
	MaxItems             *int64            `json:"maxItems,omitempty"`
	MaxLength            *int64            `json:"maxLength,omitempty"`
	MaxProperties        *int64            `json:"maxProperties,omitempty"`
	Maximum              *float64          `json:"maximum,omitempty"`
	MinItems             *int64            `json:"minItems,omitempty"`
	MinLength            *int64            `json:"minLength,omitempty"`
	MinProperties        *int64            `json:"minProperties,omitempty"`
	Minimum              *float64          `json:"minimum,omitempty"`
	MultipleOf           *float64          `json:"multipleOf,omitempty"`
	Not                  *Schema           `json:"not,omitempty"`
	OneOf                *SchemaList       `json:"oneOf,omitempty"`
	Pattern              *regexp.Regexp    `json:"pattern,omitempty"`
	PatternProperties    *SchemaSet        `json:"patternProperties,omitempty"`
	Properties           *SchemaSet        `json:"properties,omitempty"`
	PropertyNames        *Schema           `json:"propertyNames,omitempty"`
	Required             []string          `json:"required,omitempty"`
	Then                 *Schema           `json:"then,omitempty"`
	Title                string            `json:"title,omitempty"`
	Type                 PrimitiveTypeList `json:"type,omitempty"`
	UniqueItems          *bool             `json:"uniqueItems,omitempty"`
}

func (s *Schema) Comment() string {
	return s.properties.Comment
}

func (s *Schema) ID() string {
	return s.properties.ID
}

func (s *Schema) Reference() string {
	return s.properties.Reference
}

func (s *Schema) SchemaRef() string {
	return s.properties.SchemaRef
}

func (s *Schema) AdditionalItems() *Schema {
	return s.properties.AdditionalItems
}

func (s *Schema) AdditionalProperties() *Schema {
	return s.properties.AdditionalProperties
}

func (s *Schema) AllOf() *SchemaList {
	return s.properties.AllOf
}

func (s *Schema) AnyOf() *SchemaList {
	return s.properties.AnyOf
}

func (s *Schema) Const() interface{} {
	return s.properties.Const
}

func (s *Schema) Contains() *SchemaList {
	return s.properties.Contains
}

func (s *Schema) Default() interface{} {
	return s.properties.Default
}

func (s *Schema) Definitions() *SchemaSet {
	return s.properties.Definitions
}

func (s *Schema) Dependencies() *SchemaSet {
	return s.properties.Dependencies
}

func (s *Schema) Description() string {
	return s.properties.Description
}

func (s *Schema) Else() *Schema {
	return s.properties.Else
}

func (s *Schema) Enum() EnumList {
	return s.properties.Enum
}

func (s *Schema) ExclusiveMaximum() *float64 {
	return s.properties.ExclusiveMaximum
}

func (s *Schema) ExclusiveMinimum() *float64 {
	return s.properties.ExclusiveMinimum
}

func (s *Schema) Format() Format {
	return s.properties.Format
}

func (s *Schema) If() *Schema {
	return s.properties.If
}

func (s *Schema) Items() *SchemaList {
	return s.properties.Items
}

func (s *Schema) MaxItems() *int64 {
	return s.properties.MaxItems
}

func (s *Schema) MaxLength() *int64 {
	return s.properties.MaxLength
}

func (s *Schema) MaxProperties() *int64 {
	return s.properties.MaxProperties
}

func (s *Schema) Maximum() *float64 {
	return s.properties.Maximum
}

func (s *Schema) MinItems() *int64 {
	return s.properties.MinItems
}

func (s *Schema) MinLength() *int64 {
	return s.properties.MinLength
}

func (s *Schema) MinProperties() *int64 {
	return s.properties.MinProperties
}

func (s *Schema) Minimum() *float64 {
	return s.properties.Minimum
}

func (s *Schema) MultipleOf() *float64 {
	return s.properties.MultipleOf
}

func (s *Schema) Not() *Schema {
	return s.properties.Not
}

func (s *Schema) OneOf() *SchemaList {
	return s.properties.OneOf
}

func (s *Schema) Pattern() *regexp.Regexp {
	return s.properties.Pattern
}

func (s *Schema) PatternProperties() *SchemaSet {
	return s.properties.PatternProperties
}

func (s *Schema) Properties() *SchemaSet {
	return s.properties.Properties
}

func (s *Schema) PropertyNames() *Schema {
	return s.properties.PropertyNames
}

func (s *Schema) Required() []string {
	return s.properties.Required
}

func (s *Schema) Then() *Schema {
	return s.properties.Then
}

func (s *Schema) Title() string {
	return s.properties.Title
}

func (s *Schema) Type() PrimitiveTypeList {
	return s.properties.Type
}

func (s *Schema) UniqueItems() *bool {
	return s.properties.UniqueItems
}

func (s *Schema) SetComment(v string) *Schema {
	s.properties.Comment = v
	return s
}

func (s *Schema) SetID(v string) *Schema {
	s.properties.ID = v
	return s
}

func (s *Schema) SetReference(v string) *Schema {
	s.properties.Reference = v
	return s
}

func (s *Schema) SetSchemaRef(v string) *Schema {
	s.properties.SchemaRef = v
	return s
}

func (s *Schema) SetAdditionalItems(v *Schema) *Schema {
	s.properties.AdditionalItems = v
	return s
}

func (s *Schema) SetAdditionalProperties(v *Schema) *Schema {
	s.properties.AdditionalProperties = v
	return s
}

func (s *Schema) AddAllOf(list ...*Schema) *Schema {
	s.properties.AllOf.Append(list...)
	return s
}

func (s *Schema) AddAnyOf(list ...*Schema) *Schema {
	s.properties.AnyOf.Append(list...)
	return s
}

func (s *Schema) SetConst(v interface{}) *Schema {
	s.properties.Const = v
	return s
}

func (s *Schema) AddContains(list ...*Schema) *Schema {
	s.properties.Contains.Append(list...)
	return s
}

func (s *Schema) SetDefault(v interface{}) *Schema {
	s.properties.Default = v
	return s
}

func (s *Schema) SetDefinitions(v *SchemaSet) *Schema {
	s.properties.Definitions = v
	return s
}

func (s *Schema) SetDependencies(v *SchemaSet) *Schema {
	s.properties.Dependencies = v
	return s
}

func (s *Schema) SetDescription(v string) *Schema {
	s.properties.Description = v
	return s
}

func (s *Schema) SetElse(v *Schema) *Schema {
	s.properties.Else = v
	return s
}

func (s *Schema) AddEnum(list ...Enum) *Schema {
	s.properties.Enum.Append(list...)
	return s
}

func (s *Schema) SetExclusiveMaximum(v float64) *Schema {
	s.properties.ExclusiveMaximum = &v
	return s
}

func (s *Schema) SetExclusiveMinimum(v float64) *Schema {
	s.properties.ExclusiveMinimum = &v
	return s
}

func (s *Schema) SetFormat(v Format) *Schema {
	s.properties.Format = v
	return s
}

func (s *Schema) SetIf(v *Schema) *Schema {
	s.properties.If = v
	return s
}

func (s *Schema) AddItems(list ...*Schema) *Schema {
	s.properties.Items.Append(list...)
	return s
}

func (s *Schema) SetMaxItems(v int64) *Schema {
	s.properties.MaxItems = &v
	return s
}

func (s *Schema) SetMaxLength(v int64) *Schema {
	s.properties.MaxLength = &v
	return s
}

func (s *Schema) SetMaxProperties(v int64) *Schema {
	s.properties.MaxProperties = &v
	return s
}

func (s *Schema) SetMaximum(v float64) *Schema {
	s.properties.Maximum = &v
	return s
}

func (s *Schema) SetMinItems(v int64) *Schema {
	s.properties.MinItems = &v
	return s
}

func (s *Schema) SetMinLength(v int64) *Schema {
	s.properties.MinLength = &v
	return s
}

func (s *Schema) SetMinProperties(v int64) *Schema {
	s.properties.MinProperties = &v
	return s
}

func (s *Schema) SetMinimum(v float64) *Schema {
	s.properties.Minimum = &v
	return s
}

func (s *Schema) SetMultipleOf(v float64) *Schema {
	s.properties.MultipleOf = &v
	return s
}

func (s *Schema) SetNot(v *Schema) *Schema {
	s.properties.Not = v
	return s
}

func (s *Schema) AddOneOf(list ...*Schema) *Schema {
	s.properties.OneOf.Append(list...)
	return s
}

func (s *Schema) SetPattern(v *regexp.Regexp) *Schema {
	s.properties.Pattern = v
	return s
}

func (s *Schema) SetPatternProperties(v *SchemaSet) *Schema {
	s.properties.PatternProperties = v
	return s
}

func (s *Schema) SetProperties(v *SchemaSet) *Schema {
	s.properties.Properties = v
	return s
}

func (s *Schema) SetPropertyNames(v *Schema) *Schema {
	s.properties.PropertyNames = v
	return s
}

func (s *Schema) AddRequired(list ...string) *Schema {
	s.properties.Required = append(s.properties.Required, list...)
	return s
}

func (s *Schema) SetThen(v *Schema) *Schema {
	s.properties.Then = v
	return s
}

func (s *Schema) SetTitle(v string) *Schema {
	s.properties.Title = v
	return s
}

func (s *Schema) AddType(list ...PrimitiveType) *Schema {
	s.properties.Type.Append(list...)
	return s
}

func (s *Schema) SetUniqueItems(v bool) *Schema {
	s.properties.UniqueItems = &v
	return s
}

var trueBytes = []byte("true")
var falseBytes = []byte("false")

func (s *Schema) MarshalJSON() ([]byte, error) {
	if s.isNegated {
		return falseBytes, nil
	}

	if s.isEmpty {
		return trueBytes, nil
	}

	return json.Marshal(s.properties)
}

func (s *Schema) UnmarshalJSON(buf []byte) error {
	if bytes.HasPrefix(buf, trueBytes) {
		*s = Schema{isEmpty: true}
		return nil
	} else if bytes.HasPrefix(buf, falseBytes) {
		*s = Schema{isNegated: true}
		return nil
	}

	var fs Schema
	if err := json.Unmarshal(buf, &fs.properties); err != nil {
		return errors.Wrap(err, "failed to unmarshal schema")
	}
	*s = fs
	return nil
}
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
}
