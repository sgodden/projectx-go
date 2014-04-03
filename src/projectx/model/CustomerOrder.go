package model

type CustomerOrder struct {
	Id                int
	OrderNumber       string
	CustomerReference string
}

func (o *CustomerOrder) Validate() {

}
