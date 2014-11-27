package model

import (
	"gopkg.in/mgo.v2/bson"
)

type CustomerOrder struct {
	Id                bson.ObjectId
	OrderNumber       string
	CustomerReference string
	Status            string
}

func (o *CustomerOrder) PrePersist() {
	thelogger.Log("CustomerOrder pre-persist")
}