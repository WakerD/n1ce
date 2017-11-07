package models

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	//	"github.com/davecgh/go-spew/spew"
)

var (
	userC         *mgo.Collection
	UserIns       = new(User)
	paramsInvalid = errors.New("bad params")
)

type User struct {
	ID       bson.ObjectId `bson:"_id,omitempty"`
	Account  string        `json:"account" form:"account"`
	Password string        `json:"password" form:"password"`
}

func (*User) Init() {
	userC = DB.C("users")
}

func (*User) Create(data User) error {
	data.HashPassword()
	err := userC.Insert(data)
	return err
}

func (*User) Update(selector, update map[string]interface{}) error {
	err := userC.Update(selector, update)
	return err
}

func (*User) ReadOne(query map[string]interface{}) (*User, error) {
	data := new(User)
	err := userC.Find(query).One(data)
	return data, err
}

func (*User) ReadMany(query map[string]interface{}, offset, limit int) ([]*User, error) {
	data := make([]*User, 0)
	err := userC.Find(query).Skip(offset).Limit(limit).All(data)
	return data, err
}

func (*User) Delete(selector map[string]interface{}, all bool) (int, error) {
	if all {
		err := userC.Remove(selector)
		return 1, err
	}
	info, err := userC.RemoveAll(selector)
	return info.Removed, err
}

/*
func (u *User) Validate() error {
	switch {
	case len(u.Phone) != 11:
		return paramsInvalid
	case len(u.Password) > 20 || len(u.Password) < 6:
		return paramsInvalid
	default:
		return nil
	}
}*/
func (user *User) HashPassword() error {
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hash)
	return nil
}

func (user *User) VerifyPassword(password string) bool {
	if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)) != nil {
		return false
	}
	return true
}
