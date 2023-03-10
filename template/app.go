/*
Copyright © 2023 yuanjun <imoowi@qq.com>

*/
package template

import (
	"github.com/imoowi/goRESTApiGen/util"
)

// import "go.mongodb.org/mongo-driver/bson/primitive"

type TemplateApp struct {
	AppName           string
	ModuleName        string
	ServiceName       string
	ModelName         string
	ModelInstanceName string
}

func (a *TemplateApp) PreModel() string {
	return `
package ` + a.AppName + `
import (
	"net/http"
	"github.com/gin-gonic/gin"
	"` + a.ModuleName + `/models"
	"` + a.ModuleName + `/services"
	"github.com/imoowi/goRESTApiGen/util/response"
	"github.com/spf13/cast"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var ` + a.AppName + `Service *services.` + util.FirstUpper(a.AppName) + `Service
	`
}

func (a *TemplateApp) PreList() string {
	return `

//	@Summary	列表
//	@Tags		` + a.AppName + `
//	@Accept		application/json
//	@Produce	application/json
//	@Param		Authorization	header	string	true	"Bearer 用户令牌"
//	@Param		page			query	int		true	"页码 (1)"
//	@Param		pageSize		query	int		false	"页数"
//	@Success	200
//	@Failure	400
//	@Failure	500
//	@Router		/api/` + a.AppName + ` [get]
func List(c *gin.Context) {
	searchKey := c.DefaultQuery("searchKey", "")
	page := cast.ToInt64(c.DefaultQuery("page", "1"))
	pageSize := cast.ToInt64(c.DefaultQuery("pageSize", "20"))
	pages, list := ` + a.AppName + `Service.List(searchKey, page, pageSize)
	res := gin.H{
		"pages": pages,
		"list":  list,
	}
	response.OK(res, c)
}

`
}

func (a *TemplateApp) PreAdd() string {
	return `

//	@Summary	添加
//	@Tags		` + a.AppName + `
//	@Accept		application/json
//	@Produce	application/json
//	@Param		Authorization	header	string				true	"Bearer 用户令牌"
//	@Param		body			body	models.` + a.ModelName + `	true	"models.` + a.ModelName + `"
//	@Success	200
//	@Failure	400
//	@Failure	500
//	@Router		/api/` + a.AppName + ` [post]
func Add(c *gin.Context) {
	var ` + a.ModelInstanceName + ` *models.` + a.ModelName + `
	err := c.ShouldBindJSON(&` + a.ModelInstanceName + `)
	if err != nil {
		response.Error(err.Error(), http.StatusBadRequest, c)
		return
	}
	id, err := ` + a.AppName + `Service.Add(` + a.ModelInstanceName + `)
	if err != nil {
		response.Error(err.Error(), http.StatusBadRequest, c)
		return
	}
	response.OK(id, c)
}

`
}

func (a *TemplateApp) PreUpdate() string {
	return `
//	@Summary	修改
//	@Tags		` + a.AppName + `
//	@Accept		application/json
//	@Produce	application/json
//	@Param		Authorization	header	string				true	"Bearer 用户令牌"
//	@Param		id				query	string				true	"id"
//	@Param		body			body	models.` + a.ModelName + `	true	"models.` + a.ModelName + `"
//	@Success	200
//	@Failure	400
//	@Failure	500
//	@Router		/api/` + a.AppName + `/:id [put]
func Update(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		response.Error("pls input id", http.StatusBadRequest, c)
		return
	}
	var ` + a.ModelInstanceName + ` *models.` + a.ModelName + `
	err := c.ShouldBindJSON(&` + a.ModelInstanceName + `)
	if err != nil {
		response.Error(err.Error(), http.StatusBadRequest, c)
		return
	}
	` + a.ModelInstanceName + `.Id, err = primitive.ObjectIDFromHex(id)
	if err != nil {
		response.Error(err.Error(), http.StatusBadRequest, c)
		return
	}
	updated, err := ` + a.AppName + `Service.Update(` + a.ModelInstanceName + `)
	if err != nil {
		response.Error(err.Error(), http.StatusBadRequest, c)
		return
	}
	response.OK(updated, c)
}

`
}

func (a *TemplateApp) PreDelete() string {
	return `
//	@Summary	删除
//	@Tags		` + a.AppName + `
//	@Accept		application/json
//	@Produce	application/json
//	@Param		Authorization	header	string	true	"Bearer 用户令牌"
//	@Param		id				query	string	true	"id"
//	@Success	200
//	@Failure	400
//	@Failure	500
//	@Router		/api/` + a.AppName + `/:id [delete]
func Delete(c *gin.Context) {
	id := c.Param("id")
	if id == " "{
		response.Error("pls input id", http.StatusBadRequest, c)
		return
	}
	deleted, err := ` + a.AppName + `Service.Delete(id)
	if err != nil {
		response.Error(err.Error(), http.StatusBadRequest, c)
		return
	}
	response.OK(deleted, c)
}
`
}

func (a *TemplateApp) PreGetOne() string {
	return `
//	@Summary	单个信息
//	@Tags		` + a.AppName + `
//	@Accept		application/json
//	@Produce	application/json
//	@Param		Authorization	header	string	true	"Bearer 用户令牌"
//	@Param		id				query	string	true	"id"
//	@Success	200
//	@Failure	400
//	@Failure	500
//	@Router		/api/` + a.AppName + `/:id [get]
func GetOne(c *gin.Context) {
	id := c.Param("id")
	if id == " "{
		response.Error("pls input id", http.StatusBadRequest, c)
		return
	}
	info, err := ` + a.AppName + `Service.GetOne(id)
	if err != nil {
		response.Error(err.Error(), http.StatusBadRequest, c)
		return
	}
	response.OK(info, c)
}

`
}

func (a *TemplateApp) PreRouter() string {
	return `
package ` + a.AppName + `
import (
	"` + a.ModuleName + `/middleware"
	"` + a.ModuleName + `/router"

	"github.com/gin-gonic/gin"
)

func init() {
	router.RegisterRoute(Routers)
}


func Routers(e *gin.Engine) {

	e.Use(middleware.RuntimeMiddleware())
	_` + a.AppName + ` := e.Group("/api/` + a.AppName + `")
	{
		//验证登录
		_` + a.AppName + `.Use(middleware.JWTAuthMiddleware())
		//验证权限
		_` + a.AppName + `.Use(middleware.CasbinMiddleware())

		_` + a.AppName + `.GET("", List)
		_` + a.AppName + `.POST("", Add)
		_` + a.AppName + `.PUT("/:id", Update)
		_` + a.AppName + `.DELETE("/:id", Delete)
		_` + a.AppName + `.GET("/:id", GetOne)
	}
}

`
}
