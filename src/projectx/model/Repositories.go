package model

import (
	"gopkg.in/mgo.v2/bson"
)

type CustomerOrderRepository interface {
	Save(order *CustomerOrder) bson.ObjectId
	Get(id string) *CustomerOrder
}
