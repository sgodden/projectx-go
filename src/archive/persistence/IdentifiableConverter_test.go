package persistence

import (
	"fmt"
	"labix.org/v2/mgo/bson"
	"reflect"
	"sgo/projectx/model"
	"testing"
)

func TestIdentifiableConverterShouldAllocateId(t *testing.T) {
	o := model.CustomerOrder{}
	m := IdentifiableConverter{}.ToBson(&o)

	value, present := m["_id"]
	if !present {
		t.Error("_id is not present")
	} else if reflect.TypeOf(value) != reflect.TypeOf(bson.NewObjectId()) {
		t.Error(fmt.Sprintf("_id is wrong type: %T", value))
	}

}
