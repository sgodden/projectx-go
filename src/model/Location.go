package model

type Location struct {
	Identity
	Coded
	description       string
	timeZone          string
}

func (o *Location) Validate() {

}

func (o *Location) Description() string {
	return o.description
}

func (o *Location) SetDescription(v string) {
	o.description = v
}

func (o *Location) TimeZone() string {
	return o.timeZone
}

func (o *Location) SetTimeZone(v string) {
	o.timeZone = v
}
