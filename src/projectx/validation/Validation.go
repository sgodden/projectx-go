package validation

type ValidationError struct {
    Path string
    Message string
}

type Validatable interface {
  Validate() []ValidationError
}
