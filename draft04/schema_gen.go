package draft04

import (
	"encoding/json"
	"regexp"

	"github.com/pkg/errors"
)

type SchemaProperties struct {
	Reference            string            `json:"$ref,omitempty"`
	SchemaRef            string            `json:"$schema,omitempty"`
	AdditionalItems      *SchemaList       `json:"additionalItems,omitempty"`
	AdditionalProperties *Schema           `json:"additionalProperties,omitempty"`
	AllOf                *SchemaList       `json:"allOf,omitempty"`
	AnyOf                *SchemaList       `json:"anyOf,omitempty"`
	Default              interface{}       `json:"default,omitempty"`
	Definitions          *SchemaSet        `json:"definitions,omitempty"`
	Dependencies         *DependencyMap    `json:"dependencies,omitempty"`
	Description          string            `json:"description,omitempty"`
	Enum                 EnumList          `json:"enum,omitempty"`
	ExclusiveMaximum     *bool             `json:"exclusiveMaximum,omitempty"`
	ExclusiveMinimum     *bool             `json:"exclusiveMinimum,omitempty"`
	Format               Format            `json:"format,omitempty"`
	ID                   string            `json:"id,omitempty"`
	Items                *Schema           `json:"items,omitempty"`
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
	Required             []string          `json:"required,omitempty"`
	Title                string            `json:"title,omitempty"`
	Type                 PrimitiveTypeList `json:"type,omitempty"`
	UniqueItems          *bool             `json:"uniqueItems,omitempty"`
}

func (s *Schema) Reference() string {
	return s.properties.Reference
}

func (s *Schema) SchemaRef() string {
	if v := s.properties.SchemaRef; v != "" {
		return v
	}
	return draft04.SchemaID
}

func (s *Schema) AdditionalItems() *SchemaList {
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

func (s *Schema) Default() interface{} {
	return s.properties.Default
}

func (s *Schema) Definitions() *SchemaSet {
	return s.properties.Definitions
}

func (s *Schema) Dependencies() *DependencyMap {
	return s.properties.Dependencies
}

func (s *Schema) Description() string {
	return s.properties.Description
}

func (s *Schema) Enum() EnumList {
	return s.properties.Enum
}

func (s *Schema) ExclusiveMaximum() *bool {
	return s.properties.ExclusiveMaximum
}

func (s *Schema) ExclusiveMinimum() *bool {
	return s.properties.ExclusiveMinimum
}

func (s *Schema) Format() Format {
	return s.properties.Format
}

func (s *Schema) ID() string {
	return s.properties.ID
}

func (s *Schema) Items() *Schema {
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

func (s *Schema) Required() []string {
	return s.properties.Required
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

func (s *Schema) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.properties)
}

func (s *Schema) UnmarshalJSON(buf []byte) error {
	var props SchemaProperties
	if err := json.Unmarshal(buf, &props); err != nil {
		return errors.Wrap(err, `failed to unmarshal schema`)
	}
	s.properties = &props
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
