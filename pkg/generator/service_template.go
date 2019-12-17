package generator

const (
	// ServiceTpl ServiceTpl.
	ServiceTpl = `
package {{.Name}}

import (
	"github.com/gofrs/uuid"
)

type service struct {
	repository Repository
}

// NewService creates a service of user.
func NewService(rep Repository) Service {
	return &service{
		repository: rep,
	}
}

// CreateUser creates a new user.
func (s *service) Create(createRequest CreateRequest) (string, error) {
	uuid, _ := uuid.NewV4()
	id := uuid.String()
	{{.Name}} := {{upperFirst .Name}}{
		ID:               id,
{{range $fieldType := .CreateRequestFields}}{{"\t"}}{{"\t"}}{{$fieldType.Field}}:{{"\t"}}createRequest.{{$fieldType.Field}},{{"\n"}}{{end}} 
	}
	if err := s.repository.Create({{.Name}}); err != nil {
		return "", err
	}
	return id, nil
}

// Get get a {{.Name}}.
func (s *service) Get(id string) (*{{upperFirst .Name}}, error) {
	{{.Name}}, err := s.repository.Get(id)
	if err != nil {
		return nil, err
	}
	return {{.Name}}, nil
}

// GetAll get all {{.Name}}.
func (s *service) GetAll() (*[]{{upperFirst .Name}}, error) {
	{{.Name}}s, err := s.repository.GetAll()
	if err != nil {
		return nil, err
	}
	return {{.Name}}s, nil
}

// Update update a {{.Name}}.
func (s *service) Update(id string, updateRequest UpdateRequest) error {
	_, err := s.repository.Get(id)
	if err != nil {
		return err
	}
	return s.repository.Update({{upperFirst .Name}}{
		ID:               id,
{{range $fieldType := .UpdateRequestFields}}{{"\t"}}{{"\t"}}{{$fieldType.Field}}:{{"\t"}}updateRequest.{{$fieldType.Field}},{{"\n"}}{{end}} 
	})
}

// Delete delete a {{.Name}}.
func (s *service) Delete(id string) error {
	_, err := s.repository.Get(id)
	if err != nil {
		return err
	}
	return s.repository.Delete(id)
}
`
)

// type serviceGenerator struct {
// 	source transform.Internal
// }

// // NewServiceGenerator returns a new serviceGenerator.
// func NewServiceGenerator(source transform.Internal) Generator {
// 	return &serviceGenerator{source}
// }

// func (h *serviceGenerator) Generate(out io.Writer) error {
// 	t := templateutil.GetTemplate()
// 	t, err := t.Parse(serviceTpl)
// 	if err != nil {
// 		return err
// 	}
// 	return t.Execute(out, h.source)
// }
