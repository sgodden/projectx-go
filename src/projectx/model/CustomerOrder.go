package model

import ( 
  "projectx/validation"
)

type CustomerOrder struct {
	Id                uint64
	OrderNumber       string
	CustomerReference string
}

func (o *CustomerOrder) Validate() []validation.ValidationError {
    return nil
}

func (o *CustomerOrder) PrePersist() {
  if o.Id == 0 {
    panic("No id on customer order")
  }
}
