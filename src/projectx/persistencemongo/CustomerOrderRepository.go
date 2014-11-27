package persistencemongo

import (
	"projectx/model"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func collection() (*mgo.Session, *mgo.Collection){
	session, err := mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}
	c := session.DB("go-projectx").C("customerOrders")
	
	return session, c
}

type CustomerOrderRepository struct {}

func (o *CustomerOrderRepository) Save(order *model.CustomerOrder) bson.ObjectId {
	session, c := collection()
	defer session.Close()
	
	order.Id = bson.NewObjectId()
	c.UpsertId(order.Id, order)
	return order.Id
}

func (o *CustomerOrderRepository) Get(id bson.ObjectId) *model.CustomerOrder {
	session, collection := collection()
	defer session.Close()
	
	ret := model.CustomerOrder{}
	
	err := collection.FindId(id).One(&ret)
	if err != nil {
		panic(err)
	}
	
	return &ret
}