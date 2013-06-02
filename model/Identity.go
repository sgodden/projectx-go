package model

type Identity struct {
	id string
}

func (i *Identity) Id() string {
	return i.id
}

func (i *Identity) SetId(s string) {
	i.id = s
}
