package persistence

import (
	"sgo/projectx/model"
	"sgo/projectx/testutil"
	"testing"
)

func TestCustomerOrderConverterShouldAllocateId(t *testing.T) {

	o := model.CustomerOrder{}
	o.SetOrderNumber("ordNum")
	o.SetCustomerReference("custRef")

	m := CustomerOrderConverter{}.ToBson(&o)

	testutil.AssertLength(m["orderNumber"], "Order number not present", t)
	testutil.AssertLength(m["customerReference"], "Customer reference not present", t)
}
