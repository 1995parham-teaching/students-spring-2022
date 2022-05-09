package student

import "githuh.com/cng-by-example/students/internal/model"

type InMemory struct {
	students map[int64]model.Student
}

func NewInMemory() *InMemory {
	return &InMemory{
		students: make(map[int64]model.Student),
	}
}

func (s *InMemory) GetAll() ([]model.Student, error) {
	students := make([]model.Student, 0)

	for _, student := range s.students {
		students = append(students, student)
	}

	return students, nil
}

func (s *InMemory) Get(id int64) (model.Student, error) {
	student, ok := s.students[id]
	if ok {
		return student, nil
	}

	return model.Student{}, ErrStudentNotFound
}

func (s *InMemory) Set(m model.Student) error {
	if _, ok := s.students[m.ID]; ok {
		return ErrStudentDuplicate
	}

	s.students[m.ID] = m

	return nil
}
