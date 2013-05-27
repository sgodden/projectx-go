package persistence

import (
	"sgo/projectx/model"
	"sgo/projectx/modelapi"
	"labix.org/v2/mgo/bson"
)

func CustomerOrderToBson(o modelapi.ICustomerOrder) (bson.M) {
	m := bson.M{
		"orderNumber": o.OrderNumber(),
		"customerReference": o.CustomerReference(),
	}
	if len(o.Id()) == 0 {
		oid := bson.NewObjectId()
		m["_id"] = oid
		o.SetId(oid.Hex())
	} else {
		m["_id"] = bson.ObjectIdHex(o.Id())
	}
	return m
}

func CustomOrderFromMap(m map[string] interface {}) modelapi.ICustomerOrder {
	o := model.CustomerOrder{}
	o.SetId(m["_id"].(bson.ObjectId).Hex())
	o.SetOrderNumber(m["orderNumber"].(string))
	o.SetCustomerReference(m["customerReference"].(string))
	return &o
}


