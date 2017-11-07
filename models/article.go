package models

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	//	"github.com/davecgh/go-spew/spew"
)

var (
	articleC   *mgo.Collection
	ArticleIns = new(Article)
)

type Article struct {
	ID      bson.ObjectId `bson:"_id,omitempty"`
	UserID  bson.ObjectId
	Title   string   `json:"account"`
	Content string   `json:"password"`
	Pics    []string `json:"pics"`
}

func (*Article) Init() {
	articleC = DB.C("articles")
}

func (*Article) Create(data Article) error {
	err := articleC.Insert(data)
	return err
}

func (*Article) Update(selector, update map[string]interface{}) error {
	err := articleC.Update(selector, update)
	return err
}

func (*Article) ReadOne(query map[string]interface{}) (*Article, error) {
	data := new(Article)
	err := articleC.Find(query).One(data)
	return data, err
}

func (*Article) ReadMany(query map[string]interface{}, offset, limit int) ([]*Article, error) {
	data := make([]*Article, 0)
	err := articleC.Find(query).Skip(offset).Limit(limit).All(data)
	return data, err
}

func (*Article) Delete(selector map[string]interface{}, all bool) (int, error) {
	if all {
		err := articleC.Remove(selector)
		return 1, err
	}
	info, err := articleC.RemoveAll(selector)
	return info.Removed, err
}
