package model

type CustomerOrder struct {
	Identity
	orderNumber       string
	customerReference string
}

func (o *CustomerOrder) Validate() {

}

func (o *CustomerOrder) OrderNumber() string {
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
