package service

import (
	"sgo/projectx/modelapi"
	"sgo/projectx/persistence"
)

var (
	customerOrderRepo persistence.ICustomerOrderRepository
)

func init() {
	customerOrderRepo = persistence.NewCustomerOrderRepository()
}

type ICustomerOrderService interface {
	Save(order modelapi.ICustomerOrder)
	FindById(id string) modelapi.ICustomerOrder
}

type customerOrderService struct {
	ICustomerOrderService
}

func (* customerOrderService) Save(order modelapi.ICustomerOrder) {
	customerOrderRepo.Save(order)
}

func (* customerOrderService) FindById(id string) modelapi.ICustomerOrder {
	return customerOrderRepo.FindById(id)
}

func NewCustomerOrderService() ICustomerOrderService {
	return &customerOrderService{}
}


