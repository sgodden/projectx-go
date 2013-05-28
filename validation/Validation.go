package validation

type ValidationError struct {
	Message string
	Path    string
}

type validator interface {
	Validate(value interface{}, path string) ValidationError
}

type NotNilValidator struct {}

func (NotNilValidator) Validate(value interface{}, path string) ValidationError {
	ret := ValidationError{}

	badRet := ValidationError{
		"May not be nil",
		path,
	}

	switch value := value.(type) {
	default:
		if value == nil {
			ret = badRet
		}
	case string:
		if len(value) == 0 {
			ret = badRet
		}
	}

	if value == 0 {
		ret = ValidationError{
			"May not be nil",
			path,
		}
	}

	return ret
}

