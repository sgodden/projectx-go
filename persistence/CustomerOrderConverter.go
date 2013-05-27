package persistence

import (
	"sgo/projectx/model"
	"sgo/projectx/modelapi"
	"sgo/projectx/util"
	"labix.org/v2/mgo/bson"
)

type CustomerOrderConverter struct {}

func (CustomerOrderConverter) ToBson(o modelapi.ICustomerOrder) (bson.M) {
	m := IdentifiableConverter{}.ToBson(o)
	util.MapUtil{}.AddAll(m, bson.M {
		"orderNumber": o.OrderNumber(),
		"customerReference": o.CustomerReference(),
		})
	return m;
}

func (CustomerOrderConverter) FromMap(m map[string] interface {}) modelapi.ICustomerOrder {
	o := model.CustomerOrder{}
	o.SetId(m["_id"].(bson.ObjectId).Hex())
	o.SetOrderNumber(m["orderNumber"].(string))
	o.SetCustomerReference(m["customerReference"].(string))
	return &o
}


