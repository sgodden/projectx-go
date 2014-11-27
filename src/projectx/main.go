package main

import (
	"fmt"
	"projectx/model"
	"projectx/persistencemongo"
)

type Logger struct {}
func (l *Logger) Log(msg string) {
	fmt.Println(msg)
}

func main() {
	
	model.SetLogger(&Logger{})
	
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
}
