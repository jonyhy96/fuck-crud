// Package transform transforms config into go struct.
package transform // todo rename this.

// Config stands for files abstract.
type Config struct {
	Name                string
	Fields              []FieldType
	CreateRequestFields []FieldType
	UpdateRequestFields []FieldType
}

// FieldType combines field with type.
type FieldType struct {
	Field string
	Type  string // this won't validate wether the type is of build in types.
}
