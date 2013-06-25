package persistence

import (
	"labix.org/v2/mgo/bson"
	"sgo/projectx/modelapi"
	"sgo/projectx/util"
)

type CustomerOrderConverter struct{}

func (CustomerOrderConverter) ToBson(o modelapi.ICustomerOrder) bson.M {
	m := IdentifiableConverter{}.ToBson(o)
	util.MapUtil{}.AddAll(m, bson.M{
		"orderNumber":       o.OrderNumber(),
		"customerReference": o.CustomerReference(),
	})
	return m
}

func (CustomerOrderConverter) FromMap(m map[string]interface{}, o modelapi.ICustomerOrder) {
	IdentifiableConverter{}.FromMap(m, o)
	o.SetOrderNumber(m["orderNumber"].(string))
	o.SetCustomerReference(m["customerReference"].(string))
}
