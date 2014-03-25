package persistence

import (
	"database/sql"
	_ "github.com/lib/pq"
	"projectx/model"
)

type ICustomerOrderRepository interface {
	Save(model.CustomerOrder)
	FindById(id int) *model.CustomerOrder
}

func NewCustomerOrderRepository() ICustomerOrderRepository {
	return &customerOrderRepository{}
}

type customerOrderRepository struct {
}

func (r *customerOrderRepository) FindById(id int) *model.CustomerOrder {
	db, err := sql.Open("postgres", "user=simon dbname=simon")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	var orderNumber string
	var customerReference string
	err = db.QueryRow("SELECT orderNumber, customerReference FROM tcustomerorder where id=", id).
		Scan(&orderNumber, &customerReference)

	if err != nil {
		panic(err)
	}

	var customerOrder = model.CustomerOrder{}
	customerOrder.SetId(id)
	customerOrder.SetOrderNumber(orderNumber)
	customerOrder.SetCustomerReference(customerReference)

	return &customerOrder
}

func (r *customerOrderRepository) Save(model.CustomerOrder) {
}
