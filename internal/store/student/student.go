package student

import (
	"errors"

	"githuh.com/cng-by-example/students/internal/model"
)

var (
	StudentNotFoundErr  = errors.New("student not found")
	StudentDuplicateErr = errors.New("student already exists")
)

type Student interface {
	GetAll() ([]model.Student, error)
	Get(id int64) (model.Student, error)
	Set(model.Student) error
}
