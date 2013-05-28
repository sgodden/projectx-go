package validation

type ValidationError struct {
	Message string
	Path    string
}

type validator interface {
	Validate(value interface{}, path string) ValidationError
}

type NotEmptyValidator struct {}

func (NotEmptyValidator) Validate(value interface{}, path string) (bool, ValidationError) {
	success := true

	ret := ValidationError{}
	badRet := ValidationError{
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

