package generator

const (
	// RepositoryTpl RepositoryTpl.
	RepositoryTpl = `
package {{.Name}}

import (
	"errors"
	"fmt"

	"github.com/jinzhu/gorm"
)

type repo struct {
	db *gorm.DB
}

// NewRepo creates a new repository.
func NewRepo(db *gorm.DB) Repository {
	return &repo{
		db: db,
	}
}

func (repo *repo) Create({{.Name}} {{upperFirst .Name}}) error {
	err := repo.db.Create(&{{.Name}}).Error
	if err != nil {
		return err
	}
	return nil
}

func (repo *repo) Get(id string) (*{{upperFirst .Name}}, error) {
	{{.Name}} := &{{upperFirst .Name}}{
		ID: id,
	}
	if err := repo.db.Find(&{{.Name}}).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("%w with id: %s", err, id)
		}
		return nil, err
	}
	return {{.Name}}, nil
}

func (repo *repo) GetAll() (*[]{{upperFirst .Name}}, error) {
	{{.Name}}s := &[]{{upperFirst .Name}}{}
	if err := repo.db.Find(&{{.Name}}s).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("%w", err)
		}
		return nil, err
	}
	return {{.Name}}s, nil
}

func (repo *repo) Update({{.Name}} {{upperFirst .Name}}) error {
	return repo.db.Model(&{{.Name}}).Updates({{.Name}}).Error
}

func (repo *repo) Delete(id string) error {
	return repo.db.Where("id = ?", id).Delete(&{{upperFirst .Name}}{}).Error // soft delete
}
`
)

// type repositoryGenerator struct {
// 	source transform.Internal
// }

// // NewRepositoryGenerator returns a new repositoryGenerator.
// func NewRepositoryGenerator(source transform.Internal) Generator {
// 	return &repositoryGenerator{source}
// }

// func (h *repositoryGenerator) Generate(out io.Writer) error {
// 	t := templateutil.GetTemplate()
// 	t, err := t.Parse(repositoryTpl)
// 	if err != nil {
// 		return err
// 	}
// 	return t.Execute(out, h.source)
// }
