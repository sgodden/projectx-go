package persistence

import (
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
	"sgo/projectx/modelapi"
	"sgo/projectx/persistence/base"
	"sgo/projectx/model"
)

type ICustomerOrderRepository interface {
	Save(modelapi.ICustomerOrder)
	FindById(id string) modelapi.ICustomerOrder
}

// Factory method to return a customer order repository instance.
func NewCustomerOrderRepository() ICustomerOrderRepository {
	return &customerOrderRepository{}
}

// Customer order repository implementation.
type customerOrderRepository struct {
}

func (r *customerOrderRepository) getColl() *mgo.Collection {
	return base.GetColl("customerOrders")
}

func (r *customerOrderRepository) Save(order modelapi.ICustomerOrder) {
	r.getColl().Insert(CustomerOrderConverter{}.ToBson(order))
}

func (r *customerOrderRepository) FindById(id string) modelapi.ICustomerOrder {
	m := make(map[string]interface{})
	r.getColl().FindId(bson.ObjectIdHex(id)).One(&m)
	ret := model.CustomerOrder{}
	CustomerOrderConverter{}.FromMap(m, &ret)
	return &ret
}
