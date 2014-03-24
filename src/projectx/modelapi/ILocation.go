package modelapi

type ILocation interface {
	Identifiable
	ICoded

	Description() string
	SetDescription(value string)

	TimeZone() string
	SetTimeZone(value string)
}
