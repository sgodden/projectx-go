package base

import (
	"labix.org/v2/mgo"
)

var (
	session *mgo.Session
)

func _createSession() {
	s, err := mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}
	session = s
}

func InitSession() {
	_createSession()
}

func GetColl(collectionName string) *mgo.Collection {
	if session == nil {
		panic("No mongo session - ensure you have called InitSession on startup")
	}
	return session.DB("go-projectx").C(collectionName)
}

