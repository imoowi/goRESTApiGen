package template

import (
	"github.com/imoowi/goRESTApiGen/util"
)

// import "go.mongodb.org/mongo-driver/bson/primitive"

type TemplateService struct {
	// Id        primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	// CreatedAt int64              `json:"createdAt" bson:"createdAt"`
	// Deleted   bool               `json:"deleted" bson:"deleted"`
}

func (t *TemplateService) PreModel(serviceName string) string {
	serviceName = util.FirstUpper(serviceName)
	return `
package services
import "go-monitor/models"
type ` + serviceName + ` struct {
	
}
	`
}

func (t *TemplateService) FuncSearch(serviceName string, modelName string) string {
	serviceName = util.FirstUpper(serviceName)
	modelName = util.FirstUpper(modelName)
	return `
func (s *` + serviceName + `) Search(searchKey string, page int64, pageSize int64) (res []*models.` + modelName + `, err error){
	return 
}
`
}
