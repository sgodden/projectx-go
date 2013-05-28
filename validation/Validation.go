package validation

import (
	"fmt"
)

// Constrained indicates that the implementing object returns a list of constraints that should be used to validate it
type Constrained interface {
	Constraints() map[string][]Constraint
}

type ConstraintViolation struct {
	Message string
	Path    string
}

type Constraint interface {
	Validate(value interface{}, path string) (bool, ConstraintViolation)
}

// Validator validates a passed instance of Constrained and returns an array of constraint violations
type Validator interface {
	Validate(obj Constrained) []ConstraintViolation
}

type validator struct {
}

func (* validator) Validate(obj Constrained) []ConstraintViolation {
	constraints := obj.Constraints()
	for key, value := range constraints {
		// TODO - use reflect to get the property value
		fmt.Println(key)
		fmt.Println(value)

		for _, value2 := range value {
			fmt.Println(value2.Validate("", "foo"))
		}
	}
	return make([]ConstraintViolation, 0)
}

func NewValidator() Validator {
	return &validator{}
}

// A constraint which does not allow the corresponding value to be empty or nil.
type NotEmptyConstraint struct {
	Constraint
}
func (NotEmptyConstraint) Validate(value interface{}, path string) (bool, ConstraintViolation) {
	success := true

	ret := ConstraintViolation{}
	badRet := ConstraintViolation{
		"May not be nil",
		path,
	}

	switch value := value.(type) {
	default:
		if value == nil {
			success = false
		}
	case string:
		if len(value) == 0 {
			success = false
		}
	case int:
		if value == 0 {
			success = false
		}
	}

	if (!success) {
		ret = badRet
	}

	return success, ret

}

