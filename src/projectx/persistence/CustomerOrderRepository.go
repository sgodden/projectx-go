package persistence

import (
	"projectx/modelapi"
	_ "github.com/lib/pq"
	"database/sql"
)

type ICustomerOrderRepository interface {
	Save(modelapi.ICustomerOrder)
	FindById(id string) modelapi.ICustomerOrder
}

func NewCustomerOrderRepository() ICustomerOrderRepository {
	return &customerOrderRepository{}
}

type customerOrderRepository struct {
}

func (r *customerOrderRepository) FindById(id string) modelapi.ICustomerOrder {
	db, err := sql.Open("postgres", "user=pqgotest dbname=pqgotest sslmode=verify-full")
	if err != nil {
		panic(err)
	}
	rows, err := db.Query("SELECT * FROM tcustomerorder")
	for rows.Next() {
		var orderNumber string
		err = rows.Scan(&orderNumber)
	}
	return nil
}

func (r *customerOrderRepository) Save(modelapi.ICustomerOrder) {
}
