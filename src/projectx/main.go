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

	fmt.Println(repo.Save(&order))
}
