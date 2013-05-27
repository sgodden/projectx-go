package modelapi

type Identifiable interface {
	Id() string
	SetId(value string)
}

