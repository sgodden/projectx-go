package resources

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

func Configure(r *mux.Router) {
	
	model.SetLogger(&logger{})
	
	s := r.PathPrefix("/orders").Subrouter()
	
	s.HandleFunc("/", query).Methods("GET")
	s.HandleFunc("/", post).Methods("POST")
}

func query(w http.ResponseWriter, r *http.Request) {
	fmt.Println("QUERY")
	response, err := json.Marshal(repo.Query())
	if err != nil {
		panic(err)
	}
	w.Write(response)
}

func post(w http.ResponseWriter, r *http.Request) {
	fmt.Println("POST")
	decoder := json.NewDecoder(r.Body)
	var order model.CustomerOrder
	err := decoder.Decode(&order)
	if err != nil {
		panic(err)
	}
	repo.Save(&order)
}
