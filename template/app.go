package template

import (
	"github.com/imoowi/goRESTApiGen/util"
)

// import "go.mongodb.org/mongo-driver/bson/primitive"

type TemplateApp struct {
	// Id        primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	// CreatedAt int64              `json:"createdAt" bson:"createdAt"`
	// Deleted   bool               `json:"deleted" bson:"deleted"`
}

func (t *TemplateApp) PreModel(module string, appName string) string {
	appName = util.FirstUpper(appName)
	return `
package services
import "` + module + `/services"

import "github.com/gin-gonic/gin"
	`
}

func (t *TemplateApp) FuncSearch(appName string, serviceName string) string {
	appName = util.FirstUpper(appName)
	serviceName = util.FirstUpper(serviceName)
	return `
func Search(c *gin.Context) {
	fmt.println("search func")
	return 
}
`
}

func (t *TemplateApp) PreRouter(appName string, serviceName string) string {
	appName = util.FirstUpper(appName)
	serviceName = util.FirstUpper(serviceName)
	return `
func Search(c *gin.Context) {
	fmt.println("search func")
	return 
}
`
}

//
