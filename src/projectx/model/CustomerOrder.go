package model

import (
	"gopkg.in/mgo.v2/bson"
	"projectx/validation"
)

type CustomerOrder struct {
	Id                bson.ObjectId
	OrderNumber       string
	CustomerReference string
	Status            string
}

func (o *CustomerOrder) Validate() []validation.ValidationError {
	return nil
}
