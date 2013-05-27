package persistence

import (
	"testing"
	"sgo/projectx/model"
)

func TestIdentifiableConverterShouldAllocateId(t *testing.T) {
	o := model.CustomerOrder{}
	IdentifiableConverter{}.ToBson(&o)

	if len(o.Id()) == 0 {
		t.Error("No id allocated")
	}

}

