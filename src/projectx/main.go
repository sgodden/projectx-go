package main

import (
	"fmt"
	"projectx/model"
	"projectx/persistencemongo"
	"net/http"
	"encoding/json"
	"github.com/gorilla/mux"
)

type logger struct {}
func (l *logger) Log(msg string) {
	fmt.Println(msg)
}

var (
	repo = persistencemongo.CustomerOrderRepository{}
)

func main() {
	
	model.SetLogger(&logger{})
	
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
	
	r.HandleFunc("/", query).Methods("GET");
	http.ListenAndServe(":8080", r)
}

func query(w http.ResponseWriter, r *http.Request) {
	response, err := json.Marshal(repo.Query())
	if err != nil {
		panic(err)
	}
	w.Write(response)
}

func save(w http.ResponseWriter, r *http.Request) {
	
}