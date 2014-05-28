package main

import (
    "projectx/persistence"
)

func main() {
	repo := persistence.NewCustomerOrderRepository()
	
	repo.FindById(1)
}
