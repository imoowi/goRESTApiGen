/*
Copyright © 2023 yuanjun <imoowi@qq.com>

*/
package template

// import "go.mongodb.org/mongo-driver/bson/primitive"

type TemplateService struct {
	ModuleName        string
	ServiceName       string
	ModelName         string
	ModelInstanceName string
}

func (t *TemplateService) PreModel() string {
	return `package services

import (
	"` + t.ModuleName + `/models"
	"github.com/imoowi/goRESTApiGen/util/response"
)
var ` + t.ModelInstanceName + ` *models.` + t.ModelName + `
type ` + t.ServiceName + ` struct {
	
}
`
}

func (t *TemplateService) PreList() string {
	return `
// 列表
func (s *` + t.ServiceName + `) List(searchKey string, page int64, pageSize int64) (pages *response.Pages, res []*models.` + t.ModelName + `) {
	pages, res = ` + t.ModelInstanceName + `.List(searchKey, page, pageSize)
	return
}
`
}

func (t *TemplateService) PreAdd() string {
	return `
// 添加
func (s *` + t.ServiceName + `) Add(lightModel *models.` + t.ModelName + `) (newId string, err error) {
	newId, err = ` + t.ModelInstanceName + `.Add(lightModel)
	return
}
`
}

func (t *TemplateService) PreUpdate() string {
	return `

// 修改
func (s *` + t.ServiceName + `) Update(lightModel *models.` + t.ModelName + `) (updated bool, err error) {
	updated, err = ` + t.ModelInstanceName + `.Update(lightModel)
	return
}

`
}

func (t *TemplateService) PreDelete() string {
	return `
// 删除
func (s *` + t.ServiceName + `) Delete(id string) (deleted bool, err error) {
	deleted, err = ` + t.ModelInstanceName + `.Delete(id)
	return
}
`
}

func (t *TemplateService) PreGetOne() string {
	return `
// 查询一个
func (s *` + t.ServiceName + `) GetOne(id string) (lightModel *models.` + t.ModelName + `, err error) {
	lightModel, err = ` + t.ModelInstanceName + `.GetOne(id)
	return
}

`
}
