package validation

import (
	"testing"
)

func TestNotEmptyConstraintDoesNotAllowEmptyString(t *testing.T) {
	success, _ := NotEmptyConstraint{}.Validate("", "foo")
	if success {
		t.Error("Should have failed")
	}
}
func TestNotEmptyConstraintAllowsNonEmptyString(t *testing.T) {
	success, _ := NotEmptyConstraint{}.Validate("bar", "foo")
	if !success {
		t.Error("Should have passed")
	}
}

