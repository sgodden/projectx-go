package validation

import "testing"

func TestNilValidatorDoesNotAllowEmptyString(t *testing.T) {
	msg := NotNilValidator{}.Validate("", "foo")
	if len(msg.Message) == 0 {
		t.Error("Should have failed")
	}
}

