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
func (s *service) Create(createRequest CreateRequest) (*{{upperFirst .Name}}, error) {
	uuid, _ := uuid.NewV4()
	id := uuid.String()

	{{.Name}} := createRequest.override({{upperFirst .Name}}{
		ID: id,
	})

	return &{{.Name}}, s.repository.Create(&{{.Name}})
}

// Get get a {{.Name}}.
func (s *service) Get(id string) (*{{upperFirst .Name}}, error) {
	return s.repository.Get({{upperFirst .Name}}{
		ID: id,
	})
}

// GetAll get all {{.Name}}.
func (s *service) GetAll(arguments map[string][]string) (*GetAllResponse, error) {
	var (
		total      int
		err        error
		{{.Name}}s []{{upperFirst .Name}}
	)

	limit, offset, err := dbutil.Paginate(arguments)
	if err != nil {
		if !errors.Is(err, dbutil.ErrNoPageInfo) {
			return nil, err
		}
	}

	total, err = s.repository.Count({{upperFirst .Name}}{})
	if err != nil {
		return nil, err
	}

	{{.Name}}s, err = s.repository.GetAll({{upperFirst .Name}}{}, limit, offset)
	if err != nil {
		return nil, err
	}

	return &GetAllResponse{
		Total: total,
		Data:  {{.Name}}s,
	}, nil
}

// Update update a {{.Name}}.
func (s *service) Update(id string, updateRequest UpdateRequest) (*{{upperFirst .Name}}, error)  {
	origin, err := s.repository.Get({{upperFirst .Name}}{
		ID: id,
	})
	if err != nil {
		return nil, err
	}

	{{.Name}} := updateRequest.override(*origin)

	return &{{.Name}}, s.repository.Update(&{{.Name}})
}

// Delete delete a {{.Name}}.
func (s *service) Delete(id string) error {
	if _, err := s.repository.Get({{upperFirst .Name}}{
		ID: id,
	}); err != nil {
		return err
	}
	
	return s.repository.Delete(id)
}
`
)
