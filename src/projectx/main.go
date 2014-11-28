package main

import (
	"fmt"
	"projectx/model"
	"projectx/persistencemongo"
)

type logger struct {}
func (l *logger) Log(msg string) {
	fmt.Println(msg)
}

func main() {
	
	model.SetLogger(&logger{})
	
	repo := persistencemongo.CustomerOrderRepository{}
	
	order := model.CustomerOrder{}
	order.OrderNumber = "ASDASD"
	order.CustomerReference = "WERWER"
	order.Status = "CREATED"

	repo.Save(&order)
	
	fmt.Println(order.Id)
	
	reloadedOrder := repo.Get(order.Id)
	fmt.Println(reloadedOrder.Id)
	fmt.Println(reloadedOrder.Status)
	
	orders := repo.Query()
	
	for _, order := range orders {
		fmt.Println(order.Id)
	}
	
	/* let's start a web server */
	
}
