package model

type CustomerOrder struct {
	id                int
	orderNumber       string
	customerReference string
}

func (o *CustomerOrder) Validate() {

}

func (o *CustomerOrder) Id() int {
	return o.id
}

func (o *CustomerOrder) OrderNumber() string {
	return o.orderNumber
}

func (o *CustomerOrder) CustomerReference() string {
	return o.customerReference
}

func (o *CustomerOrder) SetId(value int) {
	o.id = value;
}

func (o *CustomerOrder) SetOrderNumber(value string) {
	o.orderNumber = value
}

func (o *CustomerOrder) SetCustomerReference(value string) {
	o.customerReference = value
}
