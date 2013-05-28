package validation

import (
	"testing"
)

func TestEmptyValidatorDoesNotAllowEmptyString(t *testing.T) {
	success, _ := NotEmptyConstraint{}.Validate("", "foo")
	if success {
		t.Error("Should have failed")
	}
}
func TestEmptyValidatorAllowsNonEmptyString(t *testing.T) {
	success, _ := NotEmptyConstraint{}.Validate("bar", "foo")
	if !success {
		t.Error("Should have passed")
	}
}

