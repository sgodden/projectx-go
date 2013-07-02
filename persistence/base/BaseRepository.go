package base

import (
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
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

type Converter interface {
	FromMap(m map[string]interface {}, o interface {})
	ToBson(o interface {}) bson.M
}

type Repository interface {
	CollectionName() string
	NewObject() interface {}
	Converter() Converter
}

func GetColl(collectionName string) *mgo.Collection {
	if session == nil {
		panic("No mongo session - ensure you have called InitSession on startup")
	}
	return session.DB("go-projectx").C(collectionName)
}

func FindById(id string, r Repository) interface {} {
	m := make(map[string]interface{})
	o := r.NewObject()
	GetColl(r.CollectionName()).FindId(bson.ObjectIdHex(id)).One(&m)
	r.Converter().FromMap(m, o)
	return o
}

func Save(o interface {}, r Repository) {
	GetColl(r.CollectionName()).Insert(r.Converter().ToBson(o))
}

