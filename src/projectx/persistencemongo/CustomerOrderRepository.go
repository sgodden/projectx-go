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

func (o *CustomerOrderRepository) Query() []model.CustomerOrder {
	session, collection := collection()
	defer session.Close()
	
	ret := make([]model.CustomerOrder, 0)
	
	collection.Find(bson.M{}).All(&ret)

	return ret;
}

func (o *CustomerOrderRepository) Save(order *model.CustomerOrder) bson.ObjectId {

	order.PrePersist()

	session, collection := collection()
	defer session.Close()

	order.Id = bson.NewObjectId()
	collection.UpsertId(order.Id, order)

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