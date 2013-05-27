package persistence

import (
	"sgo/projectx/modelapi"
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
	"fmt"
)

var (
	coll *mgo.Collection
)

func init() {
	session, err := mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}
	coll = session.DB("go-projectx").C("customerOrders")
}

type ICustomerOrderRepository interface {
	Save(modelapi.ICustomerOrder)
	FindById(id string) modelapi.ICustomerOrder
}

func NewCustomerOrderRepository() ICustomerOrderRepository {
	return &customerOrderRepository{}
}

type customerOrderRepository struct {
}

func (* customerOrderRepository) Save(order modelapi.ICustomerOrder) {
	fmt.Println(order)
	coll.Insert(CustomerOrderToBson(order))
}

func (* customerOrderRepository) FindById(id string) modelapi.ICustomerOrder {
	m := make(map[string] interface {})
	coll.FindId(bson.ObjectIdHex(id)).One(&m)
	return CustomOrderFromMap(m)
}

