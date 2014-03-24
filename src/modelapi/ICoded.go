package modelapi

type ICoded interface {
	Code() string
	SetCode(value string)
}
