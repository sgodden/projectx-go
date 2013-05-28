package validation

import (
	"testing"
)

func TestEmptyValidatorDoesNotAllowEmptyString(t *testing.T) {
	success, _ := NotEmptyValidator{}.Validate("", "foo")
	if success {
		t.Error("Should have failed")
	}
}
func TestEmptyValidatorAllowsNonEmptyString(t *testing.T) {
	success, _ := NotEmptyValidator{}.Validate("bar", "foo")
	if !success {
		t.Error("Should have passed")
	}
}

