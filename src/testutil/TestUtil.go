package testutil

import "testing"

func AssertLength(s interface{}, msg string, t *testing.T) {
	if len(s.(string)) == 0 {
		t.Error(msg)
	}
}
