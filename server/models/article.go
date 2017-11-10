package models

import (
	//	"github.com/davecgh/go-spew/spew"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

var (
	articleC   *mgo.Collection
	ArticleIns = new(Article)
)

type Article struct {
	ID      bson.ObjectId `bson:"_id,omitempty"`
	UserID  bson.ObjectId
	Title   string   `json:"title"`
	Content string   `json:"content"`
	Pics    []string `json:"pics"`
}

func (*Article) Init() {
	articleC = DB.C("articles")
}

func (*Article) Create(data Article) error {
	err := articleC.Insert(data)
	return err
}

func (*Article) Update(selector, update interface{}) error {
	err := articleC.Update(selector, update)
	return err
}

func (*Article) ReadOne(selector interface{}) (*Article, error) {
	data := new(Article)
	err := articleC.Find(selector).One(data)
	return data, err
}

func (*Article) ReadMany(query interface{}, offset, limit int) ([]*Article, error) {
	data := make([]*Article, 0)
	err := articleC.Find(query).Skip(offset).Limit(limit).All(&data)
	return data, err
}

func (*Article) Delete(selector interface{}, all bool) (int, error) {
	if all {
		err := articleC.Remove(selector)
		return 1, err
	}
	info, err := articleC.RemoveAll(selector)
	return info.Removed, err
}

/*
func (a Article) Valid() bool {
	if a.ID != nil || a.UserID != nil {
		return false
	}
	return true
}*/
