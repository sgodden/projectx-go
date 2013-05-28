package validation

import (
	"testing"
	"fmt"
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

type foo struct {
	someProperty string
}
func (f *foo) SomeProperty() string {
	return f.someProperty
}
func (* foo) Constraints() map[string][]Constraint {
	return map[string][]Constraint {
		"SomeProperty": []Constraint{NotEmptyConstraint{}},
	}
}

func TestValidatorValidatesConstrained(t *testing.T) {
	v := NewValidator()
	c := foo{"bar"}
	violations := v.Validate(&c)
	fmt.Println("asd")
	fmt.Println(len(violations))

	if len(violations) == 0 {
		t.Error("ASD")
	}

}

