package main

import (
	"fmt"
	"projectx/persistence"
)

func main() {
	repo := persistence.NewCustomerOrderRepository()

	fmt.Println(repo.FindById(1))
}
