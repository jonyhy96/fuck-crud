package generator

const (
	// InternalTpl InternalTpl.
	InternalTpl = `
package {{.Name}}

// Repository stands for interface that must be implemented.
type Repository interface {
	Create({{.Name}} *{{upperFirst .Name}}) error
	Get({{.Name}} {{upperFirst .Name}}) (*{{upperFirst .Name}}, error)
	GetAll(where map[string]interface{}, limit int, offset int) ([]{{upperFirst .Name}}, error)
	Update({{.Name}} *{{upperFirst .Name}}) error
	Delete(id string) error
	Count(where map[string]interface{}) (int, error)
}

// Service stands for interface that must be implemented.
type Service interface {
	Create({{.Name}} CreateRequest) (*{{upperFirst .Name}}, error)
	Get(id string) (*{{upperFirst .Name}}, error)
	GetAll(arguments map[string][]string) (*GetAllResponse, error)
	Update(id string, {{.Name}} UpdateRequest) (*{{upperFirst .Name}}, error)
	Delete(id string) error
}

//go:generate gomodifytags -file $GOFILE -struct {{upperFirst .Name}} -add-tags json --skip-unexported -transform camelcase -w
// {{upperFirst .Name}} handles {{.Name}} logic.
type {{upperFirst .Name}} struct {
{{range $fieldType := .Fields}}{{"\t"}}{{$fieldType.Field}}   {{$fieldType.Type}} {{"\n"}}{{end}} 
}

//go:generate gomodifytags -file $GOFILE -struct CreateRequest -add-tags json --skip-unexported -transform camelcase -w
// CreateRequest stands for a create request model.
type CreateRequest struct {
{{range $fieldType := .CreateRequestFields}}{{"\t"}}{{$fieldType.Field}}   {{$fieldType.Type}}  {{"\n"}}{{end}} 
}

func (c CreateRequest) override(in {{upperFirst .Name}}) {{upperFirst .Name}} {
{{range $fieldType := .Fields}} 
	if c.{{$fieldType.Field}} != "" {
		in.{{$fieldType.Field}} = c.{{$fieldType.Field}}
	}{{end}}
	return in
}

// Validate implements validate.Validator.
func (c CreateRequest) Validate() []validate.Error {
	var errors []validate.Error

	return errors
}

//go:generate gomodifytags -file $GOFILE -struct UpdateRequest -add-tags json --skip-unexported -transform camelcase -w
// UpdateRequest stands for a update request model.
type UpdateRequest struct {
{{range $fieldType := .UpdateRequestFields}}{{"\t"}}{{$fieldType.Field}}   {{$fieldType.Type}}  {{"\n"}}{{end}} 
}

func (u UpdateRequest) override(in {{upperFirst .Name}}) {{upperFirst .Name}} {
{{range $fieldType := .Fields}} 
	if u.{{$fieldType.Field}} != "" {
		in.{{$fieldType.Field}} = u.{{$fieldType.Field}}
	}{{end}}
	return in
}

// Validate implements validate.Validator.
func (u UpdateRequest) Validate() []validate.Error {
	var errors []validate.Error

	return errors
}

// GetAllResponse GetAllResponse.
type GetAllResponse struct {
	Data  []{{upperFirst .Name}}
	Total int
}
`
)
