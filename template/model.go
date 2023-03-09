package template

import (
	"github.com/imoowi/goRESTApiGen/util"
)

// import "go.mongodb.org/mongo-driver/bson/primitive"

type TemplateModel struct {
	// Id        primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	// CreatedAt int64              `json:"createdAt" bson:"createdAt"`
	// Deleted   bool               `json:"deleted" bson:"deleted"`
}

func (t *TemplateModel) PreModel(modelName string) string {
	modelName = util.FirstUpper(modelName)
	return `
package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type ` + modelName + ` struct {
	Id        primitive.ObjectID ` + "`" + `json:"id" bson:"_id,omitempty"` + "`" + `
	CreatedAt int64              ` + "`" + `json:"createdAt" bson:"createdAt"` + "`" + `
	Deleted   bool               ` + "`" + `json:"deleted" bson:"deleted"` + "`" + `
// add your code next
}
	`
}
