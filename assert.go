package validation

import "errors"

type Assert struct {
}

func NewAssertValidator() Assert {
	return Assert{}
}

func (Assert) StringNotEmpty(value string, message string) error {
	if len(value) == 0 {
		return errors.New(message)
	}

	return nil
}

func (Assert) Int32Gt(value int32, valueToCompare int32, message string) error {
	if value <= valueToCompare {
		return errors.New(message)
	}

	return nil
}
