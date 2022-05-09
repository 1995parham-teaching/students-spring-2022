package request

import (
	"fmt"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

const MaxAge = 100

type Student struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Age       int    `json:"age"`
}

func (s Student) Validate() error {
	if err := validation.ValidateStruct(&s,
		validation.Field(&s.FirstName, validation.Required),
		validation.Field(&s.LastName, validation.Required),
		validation.Field(&s.Age, validation.Required, validation.Min(0), validation.Max(MaxAge)),
	); err != nil {
		return fmt.Errorf("student validation failed %w", err)
	}

	return nil
}
