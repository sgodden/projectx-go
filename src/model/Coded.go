package model

type Coded struct {
	code string
}

func (i *Coded) Code() string {
	return i.code
}

func (i *Coded) SetCode(s string) {
	i.code = s
}
