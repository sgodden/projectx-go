package validation

type Constrained interface {
	GetConstraints() map[string][]Constraint
}

type ConstraintViolation struct {
	Message string
	Path    string
}

type Constraint interface {
	Validate(value interface{}, path string) ConstraintViolation
}

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

