package main

import (
	"fmt"
	"sgo/projectx/model"
	"sgo/projectx/service"
	"sgo/projectx/persistence/base"
)

type Foo interface {
	CustomerReference() string
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

}
