/*
Copyright © 2023 yuanjun <imoowi@qq.com>

*/
package template

import (
	"strings"
)

// import "go.mongodb.org/mongo-driver/bson/primitive"

type TemplateModel struct {
	ModuleName        string
	ServiceName       string
	ModelName         string
	ModelCollName     string
	ModelInstanceName string
}

func (t *TemplateModel) PreModel() string {
	return `
package models

import (
	"context"
	"log"
	"time"
	"` + t.ModuleName + `/global"
	"github.com/imoowi/goRESTApiGen/util/response"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)
const TABLE_NAME_` + strings.ToUpper(t.ModelName) + ` = "` + strings.ToLower(t.ModelCollName) + `"

type ` + t.ModelName + ` struct {
	Id        primitive.ObjectID ` + "`" + `json:"id" bson:"_id,omitempty"` + "`" + `
	Name      string             ` + "`" + `json:"name" bson:"name" binding:"required"` + "`" + `
	CreatedAt int64              ` + "`" + `json:"createdAt" bson:"createdAt"` + "`" + `
	Deleted   bool               ` + "`" + `json:"deleted" bson:"deleted"` + "`" + `
// add your code next
}
	`
}

func (t *TemplateModel) PreList() string {

	tableName := `TABLE_NAME_` + strings.ToUpper(t.ModelName)
	return `
// 列表
func (m *` + t.ModelName + `) List(searchKey string, page int64, pageSize int64) (pages response.Pages, res []*` + t.ModelName + `) {
		coll := global.Mongo.Collection(` + tableName + `)
		filter := bson.M{}
		filter["deleted"] = false
		if searchKey != "" {
			filter["name"] = bson.M{"$regex": primitive.Regex{Pattern: searchKey, Options: "i"}}
		}

		count, err := coll.CountDocuments(context.TODO(), filter)
		if err != nil {
			log.Fatal(err)
		}
		cur, err := coll.Find(context.TODO(),
			filter,
			options.Find().SetLimit(pageSize),
			options.Find().SetSkip(pageSize*(page-1)),
			options.Find().SetSort(bson.M{
				"createdAt": -1,
			}),
		)
		if err != nil {
			log.Fatal(err)
		}
		cur.All(context.TODO(), &res)
		if err := cur.Err(); err != nil {
			log.Fatal(err)
		}
		cur.Close(context.TODO())
		pages = response.MakePages(count, page, pageSize)
		return
	}	
`
}

func (t *TemplateModel) PreAdd() string {
	tableName := `TABLE_NAME_` + strings.ToUpper(t.ModelName)
	return `
// 添加
func (m *` + t.ModelName + `) Add(` + t.ModelInstanceName + ` *` + t.ModelName + `) (newId string, err error) {
	` + t.ModelInstanceName + `.CreatedAt = time.Now().Unix()
	coll := global.Mongo.Collection(` + tableName + `)
	res, err := coll.InsertOne(context.TODO(), ` + t.ModelInstanceName + `)
	insertedId := res.InsertedID
	newId = insertedId.(primitive.ObjectID).Hex()
	return
}
`
}

func (t *TemplateModel) PreUpdate() string {
	tableName := `TABLE_NAME_` + strings.ToUpper(t.ModelName)
	return `

// 修改
func (m *` + t.ModelName + `) Update(` + t.ModelInstanceName + ` *` + t.ModelName + `) (updated bool, err error) {
	coll := global.Mongo.Collection(` + tableName + `)
	_id, _ := primitive.ObjectIDFromHex(` + t.ModelInstanceName + `.Id.Hex())
	wareByte, _ := bson.Marshal(` + t.ModelInstanceName + `)
	updateFields := bson.M{}
	bson.Unmarshal(wareByte, &updateFields)
	update := bson.M{
		"$set": updateFields,
	}
	res, err := coll.UpdateByID(context.TODO(), _id, update)
	return res.ModifiedCount > 0, err
}

`
}

func (t *TemplateModel) PreDelete() string {
	tableName := `TABLE_NAME_` + strings.ToUpper(t.ModelName)
	return `

// 软删除
func (m *` + t.ModelName + `) Delete(id string) (deleted bool, err error) {
	coll := global.Mongo.Collection(` + tableName + `)
	_id, _ := primitive.ObjectIDFromHex(id)
	updateFields := bson.M{}
	updateFields["deleted"] = true
	update := bson.M{
		"$set": updateFields,
	}
	res, err := coll.UpdateByID(context.TODO(), _id, update)
	return res.ModifiedCount > 0, err
}

`
}

func (t *TemplateModel) PreGetOne() string {
	tableName := `TABLE_NAME_` + strings.ToUpper(t.ModelName)
	return `
// 查询一个
func (m *` + t.ModelName + `) GetOne(id string) (` + t.ModelInstanceName + ` *` + t.ModelName + `, err error) {
	coll := global.Mongo.Collection(` + tableName + `)
	_id, _ := primitive.ObjectIDFromHex(id)
	filter := bson.M{"_id": _id, "deleted": false}
	err = coll.FindOne(context.TODO(), filter).Decode(&` + t.ModelInstanceName + `)
	return
}

`
}
