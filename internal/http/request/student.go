package request

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type Student struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Age       int    `json:"age"`
}

func (s Student) Validate() error {
	return validation.ValidateStruct(&s,
		validation.Field(&s.FirstName, validation.Required),
		validation.Field(&s.LastName, validation.Required),
		validation.Field(&s.Age, validation.Required, validation.Min(0), validation.Max(100)),
	)
}
