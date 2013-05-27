package modelapi

type ICustomerOrder interface {
	Identifiable

	OrderNumber() string
	SetOrderNumber(value string)

	CustomerReference() string
	SetCustomerReference(value string)
}
