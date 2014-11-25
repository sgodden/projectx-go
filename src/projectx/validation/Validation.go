package validation

type ValidationError struct {
    Path string
    Message string
}

interface Validatable {
  Validate() []ValidationError
}
