package student

import "githuh.com/cng-by-example/students/internal/model"

type StudentInMemory struct {
	students map[int64]model.Student
}

func NewInMemory() *StudentInMemory {
	return &StudentInMemory{
		students: make(map[int64]model.Student),
	}
}

func (s *StudentInMemory) GetAll() ([]model.Student, error) {
	students := make([]model.Student, 0)

	for _, student := range s.students {
		students = append(students, student)
	}

	return students, nil
}

func (s *StudentInMemory) Get(id int64) (model.Student, error) {
	student, ok := s.students[id]
	if ok {
		return student, nil
	}

	return model.Student{}, StudentNotFoundErr
}

func (s *StudentInMemory) Set(m model.Student) error {
	if _, ok := s.students[m.ID]; ok {
		return StudentDuplicateErr
	}

	s.students[m.ID] = m

	return nil
}
