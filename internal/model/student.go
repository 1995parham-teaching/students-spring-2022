package model

type Student struct {
	ID        int64
	FirstName string
	LastName  string
	Average   float64
	Age       int
	Courses   []Course
}
