package resources

import (
	"fmt"
	"projectx/model"
	"projectx/persistencemongo"
	"net/http"
	"encoding/json"
	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2/bson"
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

	r.HandleFunc("/order", query).Methods("GET")
	r.HandleFunc("/order/{id}", get).Methods("GET")
	r.HandleFunc("/order", post).Methods("POST")
}

func query(w http.ResponseWriter, r *http.Request) {
	fmt.Println("QUERY")
	response, err := json.Marshal(repo.Query())
	if err != nil {
		panic(err)
	}
	w.Write(response)
}

func get(w http.ResponseWriter, r *http.Request) {
	fmt.Println("GET")
	vars := mux.Vars(r)
	
	response, err := json.Marshal(bson.ObjectId{repo.Get(vars["id"])})
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
