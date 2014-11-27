package main

import (
	"fmt"
	"projectx/model"
	"projectx/persistencemongo"
)

func main() {
	
	repo := persistencemongo.CustomerOrderRepository{}
	
	order := model.CustomerOrder{}
	order.OrderNumber = "ASDASD"
	order.CustomerReference = "WERWER"

	repo.Save(&order)
	
	fmt.Println(order.Id)
	
	reloadedOrder := repo.Get(order.Id)
	fmt.Println(reloadedOrder.Id)
}
