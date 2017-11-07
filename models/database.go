package models

import (
	"log"

	"gopkg.in/mgo.v2"
)

var session mgo.Session
var DB *mgo.Database

func SetDB(url, database string) error {
	session, err := mgo.Dial(url)
	if err != nil {
		log.Fatalf("set db wrong")
		session.Close()
		return err
	}
	DB = session.DB(database)
	UserIns.Init()
	ArticleIns.Init()
	return nil
}

func GetDB() *mgo.Database {
	return DB
}
