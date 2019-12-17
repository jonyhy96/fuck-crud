package generator

const (
	// InternalTpl InternalTpl.
	InternalTpl = `
package {{.Name}}

// Repository stands for interface that must be implemented.
type Repository interface {
	Create({{.Name}} {{upperFirst .Name}}) error
	Get(id string) (*{{upperFirst .Name}}, error)
	GetAll() (*[]{{upperFirst .Name}}, error)
	Update({{.Name}} {{upperFirst .Name}}) error
	Delete(id string) error
}

// Service stands for interface that must be implemented.
type Service interface {
	Create({{.Name}} CreateRequest) (string, error)
	Get(id string) (*{{upperFirst .Name}}, error)
	GetAll() (*[]{{upperFirst .Name}}, error)
	Update(id string, {{.Name}} UpdateRequest) error
	Delete(id string) error
}

// {{upperFirst .Name}} handles {{.Name}} logic.
type {{upperFirst .Name}} struct {
{{range $fieldType := .Fields}}{{"\t"}}{{$fieldType.Field}}   {{$fieldType.Type}} {{"\n"}}{{end}} 
}

// CreateRequest stands for a create request model.
type CreateRequest struct {
{{range $fieldType := .CreateRequestFields}}{{"\t"}}{{$fieldType.Field}}   {{$fieldType.Type}} {{"\n"}}{{end}} 
}

// UpdateRequest stands for a update request model.
type UpdateRequest struct {
{{range $fieldType := .UpdateRequestFields}}{{"\t"}}{{$fieldType.Field}}   {{$fieldType.Type}} {{"\n"}}{{end}} 
}
`
)
