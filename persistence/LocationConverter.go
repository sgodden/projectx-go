package persistence

import (
	"labix.org/v2/mgo/bson"
	"sgo/projectx/model"
	"sgo/projectx/modelapi"
	"sgo/projectx/util"
)

type LocationConverter struct{}

func (LocationConverter) ToBson(o modelapi.ILocation) bson.M {
	m := IdentifiableConverter{}.ToBson(o)
	util.MapUtil{}.AddAll(m, CodedConverter{}.ToBson(o))
	util.MapUtil{}.AddAll(m, bson.M{
			"timeZone":       o.TimeZone(),
		})
	return m
}

func (LocationConverter) FromMap(m map[string]interface{}) modelapi.ILocation {
	o := model.Location{}
	return &o
}
