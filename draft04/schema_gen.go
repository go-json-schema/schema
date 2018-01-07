package draft04

import (
	"encoding/json"
	"regexp"

	"github.com/pkg/errors"
)

type SchemaProperties struct {
	Reference            *string           `json:"$ref,omitempty"`
	SchemaRef            *string           `json:"$schema,omitempty"`
	AdditionalItems      *SchemaList       `json:"additionalItems,omitempty"`
	AdditionalProperties *Schema           `json:"additionalProperties,omitempty"`
	AllOf                *SchemaList       `json:"allOf,omitempty"`
	AnyOf                *SchemaList       `json:"anyOf,omitempty"`
	Default              interface{}       `json:"default,omitempty"`
	Definitions          *SchemaSet        `json:"definitions,omitempty"`
	Dependencies         *DependencyMap    `json:"dependencies,omitempty"`
	Description          *string           `json:"description,omitempty"`
	Enum                 EnumList          `json:"enum,omitempty"`
	ExclusiveMaximum     *bool             `json:"exclusiveMaximum,omitempty"`
	ExclusiveMinimum     *bool             `json:"exclusiveMinimum,omitempty"`
	Format               *Format           `json:"format,omitempty"`
	ID                   *string           `json:"id,omitempty"`
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
	Required             []string          `json:"required,omitempty"`
	Title                *string           `json:"title,omitempty"`
	Type                 PrimitiveTypeList `json:"type,omitempty"`
	UniqueItems          *bool             `json:"uniqueItems,omitempty"`
}

func (s *Schema) Reference() string {
	if !s.HasReference() {
		return ""
	}
	return *(s.properties.Reference)
}

func (s *Schema) HasReference() bool {
	return s.properties.Reference != nil
}

func (s *Schema) SchemaRef() string {
	if !s.HasSchemaRef() {
		return SchemaID
	}
	return *(s.properties.SchemaRef)
}

func (s *Schema) HasSchemaRef() bool {
	return s.properties.SchemaRef != nil
}

func (s *Schema) AdditionalItems() *SchemaList {
	return s.properties.AdditionalItems
}

func (s *Schema) HasAdditionalItems() bool {
	return s.properties.AdditionalItems != nil
}

func (s *Schema) AdditionalProperties() *Schema {
	return s.properties.AdditionalProperties
}

func (s *Schema) HasAdditionalProperties() bool {
	return s.properties.AdditionalProperties != nil
}

func (s *Schema) AllOf() *SchemaList {
	return s.properties.AllOf
}

func (s *Schema) HasAllOf() bool {
	return s.properties.AllOf != nil
}

func (s *Schema) AnyOf() *SchemaList {
	return s.properties.AnyOf
}

func (s *Schema) HasAnyOf() bool {
	return s.properties.AnyOf != nil
}

func (s *Schema) Default() interface{} {
	return s.properties.Default
}

func (s *Schema) HasDefault() bool {
	return s.properties.Default != nil
}

func (s *Schema) Definitions() *SchemaSet {
	return s.properties.Definitions
}

func (s *Schema) HasDefinitions() bool {
	return s.properties.Definitions != nil
}

func (s *Schema) Dependencies() *DependencyMap {
	return s.properties.Dependencies
}

func (s *Schema) HasDependencies() bool {
	return s.properties.Dependencies != nil
}

func (s *Schema) Description() string {
	if !s.HasDescription() {
		return ""
	}
	return *(s.properties.Description)
}

func (s *Schema) HasDescription() bool {
	return s.properties.Description != nil
}

func (s *Schema) Enum() EnumList {
	return s.properties.Enum
}

func (s *Schema) HasEnum() bool {
	return s.properties.Enum != nil
}

func (s *Schema) ExclusiveMaximum() bool {
	if !s.HasExclusiveMaximum() {
		return false
	}
	return *(s.properties.ExclusiveMaximum)
}

func (s *Schema) HasExclusiveMaximum() bool {
	return s.properties.ExclusiveMaximum != nil
}

func (s *Schema) ExclusiveMinimum() bool {
	if !s.HasExclusiveMinimum() {
		return false
	}
	return *(s.properties.ExclusiveMinimum)
}

func (s *Schema) HasExclusiveMinimum() bool {
	return s.properties.ExclusiveMinimum != nil
}

func (s *Schema) Format() Format {
	if !s.HasFormat() {
		return Format("")
	}
	return *(s.properties.Format)
}

func (s *Schema) HasFormat() bool {
	return s.properties.Format != nil
}

func (s *Schema) ID() string {
	if !s.HasID() {
		return ""
	}
	return *(s.properties.ID)
}

func (s *Schema) HasID() bool {
	return s.properties.ID != nil
}

func (s *Schema) Items() *SchemaList {
	return s.properties.Items
}

func (s *Schema) HasItems() bool {
	return s.properties.Items != nil
}

func (s *Schema) MaxItems() int64 {
	if !s.HasMaxItems() {
		return int64(0)
	}
	return *(s.properties.MaxItems)
}

func (s *Schema) HasMaxItems() bool {
	return s.properties.MaxItems != nil
}

func (s *Schema) MaxLength() int64 {
	if !s.HasMaxLength() {
		return int64(0)
	}
	return *(s.properties.MaxLength)
}

func (s *Schema) HasMaxLength() bool {
	return s.properties.MaxLength != nil
}

func (s *Schema) MaxProperties() int64 {
	if !s.HasMaxProperties() {
		return int64(0)
	}
	return *(s.properties.MaxProperties)
}

func (s *Schema) HasMaxProperties() bool {
	return s.properties.MaxProperties != nil
}

func (s *Schema) Maximum() float64 {
	if !s.HasMaximum() {
		return float64(0)
	}
	return *(s.properties.Maximum)
}

func (s *Schema) HasMaximum() bool {
	return s.properties.Maximum != nil
}

func (s *Schema) MinItems() int64 {
	if !s.HasMinItems() {
		return int64(0)
	}
	return *(s.properties.MinItems)
}

func (s *Schema) HasMinItems() bool {
	return s.properties.MinItems != nil
}

func (s *Schema) MinLength() int64 {
	if !s.HasMinLength() {
		return int64(0)
	}
	return *(s.properties.MinLength)
}

func (s *Schema) HasMinLength() bool {
	return s.properties.MinLength != nil
}

func (s *Schema) MinProperties() int64 {
	if !s.HasMinProperties() {
		return int64(0)
	}
	return *(s.properties.MinProperties)
}

func (s *Schema) HasMinProperties() bool {
	return s.properties.MinProperties != nil
}

func (s *Schema) Minimum() float64 {
	if !s.HasMinimum() {
		return float64(0)
	}
	return *(s.properties.Minimum)
}

func (s *Schema) HasMinimum() bool {
	return s.properties.Minimum != nil
}

func (s *Schema) MultipleOf() float64 {
	if !s.HasMultipleOf() {
		return float64(0)
	}
	return *(s.properties.MultipleOf)
}

func (s *Schema) HasMultipleOf() bool {
	return s.properties.MultipleOf != nil
}

func (s *Schema) Not() *Schema {
	return s.properties.Not
}

func (s *Schema) HasNot() bool {
	return s.properties.Not != nil
}

func (s *Schema) OneOf() *SchemaList {
	return s.properties.OneOf
}

func (s *Schema) HasOneOf() bool {
	return s.properties.OneOf != nil
}

func (s *Schema) Pattern() *regexp.Regexp {
	return s.properties.Pattern
}

func (s *Schema) HasPattern() bool {
	return s.properties.Pattern != nil
}

func (s *Schema) PatternProperties() *SchemaSet {
	return s.properties.PatternProperties
}

func (s *Schema) HasPatternProperties() bool {
	return s.properties.PatternProperties != nil
}

func (s *Schema) Properties() *SchemaSet {
	return s.properties.Properties
}

func (s *Schema) HasProperties() bool {
	return s.properties.Properties != nil
}

func (s *Schema) Required() []string {
	return s.properties.Required
}

func (s *Schema) HasRequired() bool {
	return s.properties.Required != nil
}

func (s *Schema) Title() string {
	if !s.HasTitle() {
		return ""
	}
	return *(s.properties.Title)
}

func (s *Schema) HasTitle() bool {
	return s.properties.Title != nil
}

func (s *Schema) Type() PrimitiveTypeList {
	return s.properties.Type
}

func (s *Schema) HasType() bool {
	return s.properties.Type != nil
}

func (s *Schema) UniqueItems() bool {
	if !s.HasUniqueItems() {
		return false
	}
	return *(s.properties.UniqueItems)
}

func (s *Schema) HasUniqueItems() bool {
	return s.properties.UniqueItems != nil
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
			ch <- &Property{name: k, definition: v}
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

func (s *SchemaList) Iterator() <-chan *Schema {
	s.mu.RLock()
	defer s.mu.RUnlock()
	ch := make(chan *Schema, len(s.store))
	go func() {
		defer close(ch)
		s.mu.RLock()
		defer s.mu.RUnlock()
		for _, e := range s.store {
			ch <- e
		}
	}()
	return ch
}

func (s *SchemaList) Append(list ...*Schema) {
	s.mu.Lock()
	defer s.mu.Unlock()
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
