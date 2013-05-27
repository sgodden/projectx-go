package persistence

import (
	"sgo/projectx/modelapi"
	"labix.org/v2/mgo/bson"
)

type IdentifiableConverter struct {
}

func (IdentifiableConverter) ToBson(o modelapi.Identifiable) bson.M {
	m := bson.M{}
	if len(o.Id()) == 0 {
		oid := bson.NewObjectId()
		m["_id"] = oid
		o.SetId(oid.Hex())
	} else {
		m["_id"] = bson.ObjectIdHex(o.Id())
	}
	return m

}
