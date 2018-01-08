package draft07

import (
	"bytes"
	"encoding/json"

	"github.com/pkg/errors"
)

type Property struct {
	name       string
	definition *Schema
}

func (p *Property) Name() string {
	return p.name
}
func (p *Property) Definition() *Schema {
	return p.definition
}

type schemaProperties struct {
	Comment              *string           `json:"$comment,omitempty"`
	ID                   *string           `json:"$id,omitempty"`
	Reference            *string           `json:"$ref,omitempty"`
	SchemaRef            *string           `json:"$schema,omitempty"`
	AdditionalItems      *Schema           `json:"additionalItems,omitempty"`
	AdditionalProperties *Schema           `json:"additionalProperties,omitempty"`
	AllOf                *SchemaList       `json:"allOf,omitempty"`
	AnyOf                *SchemaList       `json:"anyOf,omitempty"`
	Const                interface{}       `json:"const,omitempty"`
	Contains             *SchemaList       `json:"contains,omitempty"`
	Default              interface{}       `json:"default,omitempty"`
	Definitions          *SchemaSet        `json:"definitions,omitempty"`
	Dependencies         *SchemaSet        `json:"dependencies,omitempty"`
	Description          *string           `json:"description,omitempty"`
	Else                 *Schema           `json:"else,omitempty"`
	Enum                 EnumList          `json:"enum,omitempty"`
	ExclusiveMaximum     *float64          `json:"exclusiveMaximum,omitempty"`
	ExclusiveMinimum     *float64          `json:"exclusiveMinimum,omitempty"`
	Format               *Format           `json:"format,omitempty"`
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
	Pattern              *string           `json:"pattern,omitempty"`
	PatternProperties    *SchemaSet        `json:"patternProperties,omitempty"`
	Properties           *SchemaSet        `json:"properties,omitempty"`
	PropertyNames        *Schema           `json:"propertyNames,omitempty"`
	Required             []string          `json:"required,omitempty"`
	Then                 *Schema           `json:"then,omitempty"`
	Title                *string           `json:"title,omitempty"`
	Type                 PrimitiveTypeList `json:"type,omitempty"`
	UniqueItems          *bool             `json:"uniqueItems,omitempty"`
}

func (s *Schema) Comment() string {
	if !s.HasComment() {
		return ""
	}
	return *(s.properties.Comment)
}

func (s *Schema) HasComment() bool {
	return s.properties.Comment != nil
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

func (s *Schema) AdditionalItems() *Schema {
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

func (s *Schema) Const() interface{} {
	return s.properties.Const
}

func (s *Schema) HasConst() bool {
	return s.properties.Const != nil
}

func (s *Schema) Contains() *SchemaList {
	return s.properties.Contains
}

func (s *Schema) HasContains() bool {
	return s.properties.Contains != nil
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

func (s *Schema) Dependencies() *SchemaSet {
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

func (s *Schema) Else() *Schema {
	return s.properties.Else
}

func (s *Schema) HasElse() bool {
	return s.properties.Else != nil
}

func (s *Schema) Enum() EnumList {
	return s.properties.Enum
}

func (s *Schema) HasEnum() bool {
	return s.properties.Enum != nil
}

func (s *Schema) ExclusiveMaximum() float64 {
	if !s.HasExclusiveMaximum() {
		return float64(0)
	}
	return *(s.properties.ExclusiveMaximum)
}

func (s *Schema) HasExclusiveMaximum() bool {
	return s.properties.ExclusiveMaximum != nil
}

func (s *Schema) ExclusiveMinimum() float64 {
	if !s.HasExclusiveMinimum() {
		return float64(0)
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

func (s *Schema) If() *Schema {
	return s.properties.If
}

func (s *Schema) HasIf() bool {
	return s.properties.If != nil
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

func (s *Schema) Pattern() string {
	if !s.HasPattern() {
		return ""
	}
	return *(s.properties.Pattern)
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

func (s *Schema) PropertyNames() *Schema {
	return s.properties.PropertyNames
}

func (s *Schema) HasPropertyNames() bool {
	return s.properties.PropertyNames != nil
}

func (s *Schema) Required() []string {
	return s.properties.Required
}

func (s *Schema) HasRequired() bool {
	return s.properties.Required != nil
}

func (s *Schema) Then() *Schema {
	return s.properties.Then
}

func (s *Schema) HasThen() bool {
	return s.properties.Then != nil
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

func (s *Schema) SetComment(v *string) *Schema {
	s.properties.Comment = v
	return s
}

func (s *Schema) SetID(v *string) *Schema {
	s.properties.ID = v
	return s
}

func (s *Schema) SetReference(v *string) *Schema {
	s.properties.Reference = v
	return s
}

func (s *Schema) SetSchemaRef(v *string) *Schema {
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

func (s *Schema) SetDescription(v *string) *Schema {
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

func (s *Schema) SetFormat(v *Format) *Schema {
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

func (s *Schema) SetPattern(v *string) *Schema {
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

func (s *Schema) SetTitle(v *string) *Schema {
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
