package student

import (
	"errors"

	"githuh.com/1995parham-teaching/students/internal/model"
)

var (
	ErrStudentNotFound  = errors.New("student not found")
	ErrStudentDuplicate = errors.New("student already exists")
)

type Student interface {
	GetAll() ([]model.Student, error)
	Get(id int64) (model.Student, error)
	Set(model.Student) error
}
