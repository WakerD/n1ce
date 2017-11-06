package models

import (
	"errors"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

var (
	collection    *mgo.Collection
	UserIns       = new(User)
	paramsInvalid = errors.New("bad params")
)

func init() {
	collection = DB.C("users")
}

type User struct {
	ID       bson.ObjectId `bson:"_id,omitempty"`
	Phone    string
	Name     string
	Password string
}

func (*User) Create(data User) error {
	err := collection.Insert(data)
	return err
}

func (*User) Update(selector, update map[string]interface{}) error {
	err := collection.Update(selector, update)
	return err
}

func (*User) ReadOne(query map[string]interface{}) (*User, error) {
	data := new(User)
	err := collection.Find(query).One(data)
	return data, err
}

func (*User) ReadMany(query map[string]interface{}, offset, limit int) ([]*User, error) {
	data := make([]*User, 0)
	err := collection.Find(query).Skip(offset).Limit(limit).All(data)
	return data, err
}

func (*User) Delete(selector map[string]interface{}, all bool) (int, error) {
	if all {
		err := collection.Remove(selector)
		return 1, err
	}
	info, err := collection.RemoveAll(selector)
	return info.Removed, err
}

func (u *User) Validate() error {
	switch {
	case len(u.Phone) != 11:
		return paramsInvalid
	case len(u.Password) > 20 || len(u.Password) < 6:
		return paramsInvalid
	default:
		return nil
	}
}
