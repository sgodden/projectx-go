package main

import (
	"fmt"
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
	
	fs := http.FileServer(http.Dir("static"))
	redirectToIndex := http.RedirectHandler("/index.htm", 302)
	
	r.Handle("", redirectToIndex)
	r.Handle("/", redirectToIndex)
	r.PathPrefix("/index.htm/").Handler(redirectToIndex)
	
	r.Handle("/index.htm", fs)
	r.PathPrefix("/app/").Handler(fs)
	
	s := r.PathPrefix("/rest").Subrouter()
	resources.Configure(s)
	
	fmt.Println("Listening on port 8080...")
	http.ListenAndServe(":8080", r)
}
