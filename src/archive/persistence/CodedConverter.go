package persistence

import (
	"labix.org/v2/mgo/bson"
	"sgo/projectx/modelapi"
)

type CodedConverter struct {
}

func (CodedConverter) ToBson(o modelapi.ICoded) bson.M {
	m := bson.M{}
	m["code"] = o.Code()
	return m

}
