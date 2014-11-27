package model

import ( 
	"projectx/validation"
	"gopkg.in/mgo.v2/bson"
)

type CustomerOrder struct {
	Id                bson.ObjectId
	OrderNumber       string
	CustomerReference string
}

func (o *CustomerOrder) Validate() []validation.ValidationError {
    return nil
}
