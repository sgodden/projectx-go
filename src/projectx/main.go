package main

import (
    "projectx/persistence"
	"fmt"
)

func main() {
	repo := persistence.NewCustomerOrderRepository()
	
	fmt.Println(repo.FindById(1))
}
