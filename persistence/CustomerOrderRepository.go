package persistence

import (
	"labix.org/v2/mgo"
	"sgo/projectx/modelapi"
	"sgo/projectx/persistence/base"
	"sgo/projectx/model"
	"fmt"
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

func (r *customerOrderRepository) CollectionName() string {
	return "customerOrders"
}

func (r *customerOrderRepository) NewObject() interface {} {
	return &model.CustomerOrder{}
}

func (r *customerOrderRepository) getColl() *mgo.Collection {
	return base.GetColl("customerOrders")
}

func (r *customerOrderRepository) Save(order modelapi.ICustomerOrder) {
	r.getColl().Insert(CustomerOrderConverter{}.ToBson(order.(*model.CustomerOrder)))
}

func (r *customerOrderRepository) FindById(id string) modelapi.ICustomerOrder {
	fmt.Println("Id " + id)
	return base.FindById(id, r, CustomerOrderConverter{}).(modelapi.ICustomerOrder)
}
