package persistence

import (
	"labix.org/v2/mgo/bson"
	"sgo/projectx/util"
	"sgo/projectx/model"
)

type CustomerOrderConverter struct{}

func (CustomerOrderConverter) ToBson(obj interface {}) bson.M {
	o := obj.(*model.CustomerOrder)
	m := IdentifiableConverter{}.ToBson(o)
	util.MapUtil{}.AddAll(m, bson.M{
		"orderNumber":       o.OrderNumber(),
		"customerReference": o.CustomerReference(),
	})
	return m
}

func (CustomerOrderConverter) FromMap(m map[string]interface{}, obj interface {}) {
	o := obj.(*model.CustomerOrder)
	IdentifiableConverter{}.FromMap(m, o)
	o.SetOrderNumber(m["orderNumber"].(string))
	o.SetCustomerReference(m["customerReference"].(string))
}
