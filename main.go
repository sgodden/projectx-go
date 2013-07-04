package main

import (
	"fmt"
	"sgo/projectx/model"
	"sgo/projectx/service"
	"sgo/projectx/persistence/base"
	"encoding/json"
	"sgo/projectx/modelapi"
	"bytes"
	"net/http"
)

type Foo interface {
	CustomerReference() string
}

type Order struct {
	Id string
	OrderNumber string
	CustomerReference string
}

func fromEntity(entity modelapi.ICustomerOrder) Order {
	o := Order{}
	o.Id = entity.Id()
	o.OrderNumber = entity.OrderNumber()
	o.CustomerReference = entity.CustomerReference()
	return o
}

func main() {
	base.InitSession()

	o := model.CustomerOrder{}
	o.SetOrderNumber("ord001")
	o.SetCustomerReference("cr001")

	s := service.NewCustomerOrderService()
	s.Save(&o)

	p := s.FindById(o.Id())
	if p != nil {
		fmt.Println(p.Id())
		fmt.Println(p.CustomerReference())
	}

	b, _ := json.Marshal(fromEntity(p))
	buffer := bytes.NewBuffer(b)

	fmt.Println(buffer.String())

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, buffer.String())
	})

	http.ListenAndServe(":8080", nil)
}
