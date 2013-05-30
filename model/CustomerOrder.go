package model

import (
	"fmt"
)

var (
	x = 4
)

func init() {
	x = 5
	fmt.Println(x);
}

type CustomerOrder struct {
	Identity
	orderNumber string
	customerReference string
}

func (o *CustomerOrder) Validate() {

}

func (o *CustomerOrder) OrderNumber() string {
	fmt.Printf("get order number %v\n", x)
	return o.orderNumber
}

func (o *CustomerOrder) CustomerReference() string {
	return o.customerReference
}

func (o *CustomerOrder) SetOrderNumber(value string) {
	o.orderNumber = value
}

func (o *CustomerOrder) SetCustomerReference(value string) {
	o.customerReference = value
}

