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

// Initialises the mongo session
func InitSession() {
	_createSession()
}

// Converter converts a map to an object and vice versa
type Converter interface {

	// Takes a map and returns an object
	FromMap(m map[string]interface {}, o interface {})

	// Takes an object and returns a BSON map
	ToBson(o interface {}) bson.M
}

type Repository interface {
	CollectionName() string
	NewObject() interface {}
	Converter() Converter
}

func getColl(collectionName string) *mgo.Collection {
	if session == nil {
		panic("No mongo session - ensure you have called InitSession on startup")
	}
	return session.DB("go-projectx").C(collectionName)
}

// Finds an object by its identity
func FindById(id string, r Repository) interface {} {
	m := make(map[string]interface{})
	o := r.NewObject()
	getColl(r.CollectionName()).FindId(bson.ObjectIdHex(id)).One(&m)
	r.Converter().FromMap(m, o)
	return o
}

func Save(o interface {}, r Repository) {
	getColl(r.CollectionName()).Insert(r.Converter().ToBson(o))
}

