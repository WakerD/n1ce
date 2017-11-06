package models

import (
	"gopkg.in/mgo.v2"
)

var session mgo.Session
var DB *mog.Database

func SetDB(url, database string) error {
	session, err := mgo.Dial(url)
	if err != nil {
		log.Fatalf("set db wrong")
		session.Close()
		return err
	}
	DB = session.DB(database)
}

func GetDB() *mgo.Database {
	return DB
}
