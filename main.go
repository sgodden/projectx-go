package main

import (
	"fmt"
	"sgo/projectx/model"
	"sgo/projectx/persistence"
)

func main() {

	o := model.CustomerOrder{}
	o.SetOrderNumber("ord001");
	o.SetCustomerReference("cr001");

	repo := persistence.NewCustomerOrderRepository()
	repo.Save(&o)

	p := repo.FindById(o.Id())
	fmt.Println(p.CustomerReference())

}
