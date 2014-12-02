package main

import (
	"net/http"
	"github.com/gorilla/mux"
	"projectx/resources"
)

func main() {
	
	// order := model.CustomerOrder{}
	// order.OrderNumber = "ASDASD"
	// order.CustomerReference = "WERWER"
	// order.Status = "CREATED"
	//
	// repo.Save(&order)
	//
	// fmt.Println(order.Id)
	//
	// reloadedOrder := repo.Get(order.Id)
	// fmt.Println(reloadedOrder.Id)
	// fmt.Println(reloadedOrder.Status)
	//
	// orders := repo.Query()
	//
	// for _, order := range orders {
	// 	fmt.Println(order.Id)
	// }
	
	/* let's start a web server */
	
	r := mux.NewRouter()
	
	resources.Configure(r)
	
	http.ListenAndServe(":8080", r)
}
