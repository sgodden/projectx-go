package base

import "labix.org/v2/mgo"

var (
	session *mgo.Session
)

func GetColl(collectionName string) *mgo.Collection {
	// FIXME - use mutexes to make this thread-safe
	if session == nil {
		s, err := mgo.Dial("localhost")
		if err != nil {
			panic(err)
		}
		session = s
	}
	return session.DB("go-projectx").C(collectionName)
}

