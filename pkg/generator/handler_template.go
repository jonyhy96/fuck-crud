package generator

const (
	// HandlerTpl HandlerTpl.
	HandlerTpl = `
package handler

import (
	"errors"

	"nitrogen/pkg/internal/{{.Name}}"
	"nitrogen/pkg/storage"
	"nitrogen/pkg/util/log"
	"nitrogen/pkg/util/validate/sql"
	"nitrogen/pkg/util/validate/uuid"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)
	
// {{upperFirst .Name}}Handler handles {{.Name}} requests.
type {{upperFirst .Name}}Handler struct {
	base // use some default functions.
}

// Init init route and storage.
func (u *{{upperFirst .Name}}Handler) Init(router *gin.RouterGroup, engine storage.Storage) {
	u.engine = engine
	u.logger = log.Logger("{{.Name}}")
	{{.Name}}Group := router.Group("/{{.Name}}s")
	{{.Name}}Group.POST("", u.Create{{upperFirst .Name}})
	{{.Name}}Group.GET("", u.GetAll{{upperFirst .Name}})
	{{.Name}}Group.GET("/:id", u.Get{{upperFirst .Name}})
	{{.Name}}Group.PUT("/:id", u.Update{{upperFirst .Name}})
	{{.Name}}Group.DELETE("/:id", u.Delete{{upperFirst .Name}})
}

// Create{{upperFirst .Name}} creates a new {{.Name}}.
// @Summary creates a new {{.Name}}
// @Description creates a new {{.Name}}.
// @Tags {{.Name}}
// @ID {{.Name}}-create
// @Accept  json
// @Produce  json
// @Param {{.Name}} body {{.Name}}.CreateRequest true "{{.Name}}"
// @Success 200 {object} {{.Name}}.{{upperFirst .Name}}
// @Failure 400 {object} errorResponse
// @Failure 404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router /{{.Name}}s [post]
func (u *{{upperFirst .Name}}Handler) Create{{upperFirst .Name}}(ctx *gin.Context) {
	var (
		req {{.Name}}.CreateRequest
	)
	err := ctx.ShouldBind(&req)
	if err != nil {
		u.logger.Error(err)
		u.badRequest(ctx, err)
		return
	}

	if validateErrors := u.validate(req); validateErrors != nil {
		u.logger.Error(validateErrors)
		u.unprocessableEntity(ctx, validateErrors)
		return
	}

	db, err := u.getDB()
	if err != nil {
		u.serviceError(ctx, err)
		return
	}
	repository := {{.Name}}.NewRepo(db)
	response, err := {{.Name}}.NewService(repository).Create(req)
	if err != nil {
		u.logger.Error(err)
		u.serviceError(ctx, err)
		return
	}
	u.success(ctx, response)
}

// Get{{upperFirst .Name}} get {{.Name}}.
// @Summary Get a {{.Name}}
// @Description get string by ID
// @Tags {{.Name}}
// @ID {{.Name}}-get
// @Accept  json
// @Produce  json
// @Param id path string true "{{.Name}} id"
// @Success 200 {object} {{.Name}}.{{upperFirst .Name}}
// @Failure 400 {object} errorResponse
// @Failure 404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router /{{.Name}}s/{id} [get]
func (u *{{upperFirst .Name}}Handler) Get{{upperFirst .Name}}(ctx *gin.Context) {
	var (
		id = ctx.Param("id")
	) 
	if err := uuid.ValidateUUID(id); err != nil {
		u.logger.Error(err)
		u.badRequest(ctx, err)
		return
	}
	db, err := u.getDB()
	if err != nil {
		u.serviceError(ctx, err)
		return
	}
	repository := {{.Name}}.NewRepo(db)
	response, err := {{.Name}}.NewService(repository).Get(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			u.logger.Error(err)
			u.notFound(ctx, err)
			return
		}
		u.logger.Error(err)
		u.serviceError(ctx, err)
		return
	}
	u.success(ctx, response)
}

// GetAll{{upperFirst .Name}} get {{.Name}}.
// @Summary Get all {{.Name}}
// @Description get all {{.Name}}s
// @Tags {{.Name}}
// @ID {{.Name}}-get-all
// @Accept  json
// @Produce  json
// @Success 200 {object} {{.Name}}.GetAllResponse
// @Failure 400 {object} errorResponse
// @Failure 404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router /{{.Name}}s [get]
func (u *{{upperFirst .Name}}Handler) GetAll{{upperFirst .Name}}(ctx *gin.Context) {
	var (
		arguments = ctx.Request.URL.Query()
	)

	db, err := u.getDB()
	if err != nil {
		u.serviceError(ctx, err)
		return
	}
	repository := {{.Name}}.NewRepo(db)
	response, err := {{.Name}}.NewService(repository).GetAll(arguments)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			u.logger.Error(err)
			u.notFound(ctx, err)
			return
		}
		u.logger.Error(err)
		u.serviceError(ctx, err)
		return
	}
	u.success(ctx, response)
}

// Update{{upperFirst .Name}} update {{.Name}}.
// @Summary Update a {{.Name}}
// @Description update {{.Name}}
// @Tags {{.Name}}
// @ID {{.Name}}-update
// @Accept  json
// @Produce  json
// @Param id path string true "{{.Name}} id"
// @Param {{.Name}} body {{.Name}}.UpdateRequest true "{{.Name}}"
// @Success 200 {object} {{.Name}}.{{upperFirst .Name}}
// @Failure 400 {object} errorResponse
// @Failure 404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router /{{.Name}}s/{id} [put]
func (u *{{upperFirst .Name}}Handler) Update{{upperFirst .Name}}(ctx *gin.Context) {
	var (
		req {{.Name}}.UpdateRequest
		id  = ctx.Param("id")
	)

	if err := uuid.ValidateUUID(id); err != nil {
		u.logger.Error(err)
		u.badRequest(ctx, err)
		return
	}

	err := ctx.ShouldBind(&req)
	if err != nil {
		u.logger.Error(err)
		u.badRequest(ctx, err)
		return
	}

	if validateErrors := u.validate(req); validateErrors != nil {
		u.logger.Error(validateErrors)
		u.unprocessableEntity(ctx, validateErrors)
		return
	}

	db, err := u.getDB()
	if err != nil {
		u.serviceError(ctx, err)
		return
	}
	repository := {{.Name}}.NewRepo(db)
	response, err := {{.Name}}.NewService(repository).Update(id, req)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			u.logger.Error(err)
			u.notFound(ctx, err)
			return
		}
		u.logger.Error(err)
		u.serviceError(ctx, err)
		return
	}
	u.success(ctx, response)
}

// Delete{{upperFirst .Name}} delete {{.Name}}.
// @Summary Delete a {{.Name}}
// @Description delete string by ID
// @Tags {{.Name}}
// @ID {{.Name}}-delete
// @Accept  json
// @Produce  json
// @Param id path string true "{{.Name}} id"
// @Success 200 {string} string "ok"
// @Failure 400 {object} errorResponse
// @Failure 404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router /{{.Name}}s/{id} [delete]
func (u *{{upperFirst .Name}}Handler) Delete{{upperFirst .Name}}(ctx *gin.Context) {
	var (
		id = ctx.Param("id")
	)

	if err := uuid.ValidateUUID(id); err != nil {
		u.logger.Error(err)
		u.badRequest(ctx, err)
		return
	}

	db, err := u.getDB()
	if err != nil {
		u.serviceError(ctx, err)
		return
	}
	repository := {{.Name}}.NewRepo(db)
	err = {{.Name}}.NewService(repository).Delete(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			u.logger.Error(err)
			u.notFound(ctx, err)
			return
		}
		u.logger.Error(err)
		u.serviceError(ctx, err)
		return
	}
	u.success(ctx, nil)
}
`
)
