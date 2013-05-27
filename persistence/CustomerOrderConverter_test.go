package persistence

import (
	"testing"
	"sgo/projectx/model"
)

func TestCustomerOrderConverterShouldAllocateId(t *testing.T) {
	o := model.CustomerOrder{}
	CustomerOrderToBson(&o)

	if len(o.Id()) == 0 {
		t.Error("No id allocated")
	}

}

