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

func (repo *repo) Create({{.Name}} *{{upperFirst .Name}}) error {
	return repo.db.Create(&{{.Name}}).Error
}

func (repo *repo) Get(where {{upperFirst .Name}}) (*{{upperFirst .Name}}, error) {
	var {{.Name}} {{upperFirst .Name}}

	if err := repo.db.Find(&{{.Name}},where).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("%w with id: %s", err, where.ID)
		}
		return nil, err
	}

	return &{{.Name}}, nil
}

func (repo *repo) Count(where map[string]interface{}) (int, error) {
	var total int

	if err := repo.db.Model({{upperFirst .Name}}{}).Where(where).Count(&total).Error; err != nil {
		return 0, err
	}

	return total, nil
}

func (repo *repo) GetAll(where map[string]interface{}, limit int, offset int) ([]{{upperFirst .Name}}, error) {
	var {{.Name}}s []{{upperFirst .Name}}

	if err := repo.db.Limit(limit).Offset(offset).Order("updated_at DESC").Find(&{{.Name}}s,where).Error; err != nil {
		return nil, err
	}

	return {{.Name}}s, nil
}

func (repo *repo) Update({{.Name}} *{{upperFirst .Name}}) error {
	return repo.db.Model(&{{.Name}}).Updates({{.Name}}).Error
}

func (repo *repo) Delete(id string) error {
	return repo.db.Where("id = ?", id).Delete(&{{upperFirst .Name}}{}).Error // soft delete
}
`
)
