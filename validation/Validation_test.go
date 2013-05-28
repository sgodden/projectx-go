package validation

import (
	"testing"
)

var (
	capturedValue string
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

func TestValidatorValidatesConstrained(t *testing.T) {
	v := NewValidator()
	c := foo{"myvalue"}
	violations := v.Validate(&c)

	// our mock validator always returns a violation
	if len(violations) != 1 {
		t.Error("There should have been one violation")
	}

	violation := violations[0]
	if violation.Message != "FOO" {
		t.Error("Wrong message on returned violation")
	}

	if capturedValue != "myvalue" {
		t.Error("Wrong value was passed to validator")
	}
}


// A simmple type for our validation tests
type foo struct {
	someProperty string
}
func (f *foo) SomeProperty() string {
	return f.someProperty
}
func (* foo) Constraints() map[string][]Constraint {
	return map[string][]Constraint {
		"SomeProperty": []Constraint{myValidator{}},
	}
}

// A mock validator for our validation tests
type myValidator struct {}
func (myValidator) Validate(value interface {}, path string) (bool, ConstraintViolation) {
	// should be passed 'myvalue' as a string
	capturedValue = value.(string)

	return false, ConstraintViolation{
		"FOO",
		"BAR",
	}
}

