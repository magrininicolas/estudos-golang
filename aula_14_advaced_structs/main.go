package main

import "fmt"

type Enrollment struct {
	id      int
	class   string
	subject string
}

type SchoolTests struct {
	id           int
	gradesNumber uint8
	grades       []float32
}

type Student struct {
	name string
	Enrollment
	SchoolTests
}

func NewStudent(name, class, subject string, grades ...float32) Student {
	maxGrades := 4
	if len(grades) < 4 {
		maxGrades = len(grades)
	}

	return Student{name, Enrollment{0, class, subject}, SchoolTests{1, uint8(maxGrades), grades[:maxGrades]}}
}

func (s SchoolTests) gradeAverage() (soma float32) {

	for _, v := range s.grades {
		soma += v
	}
	soma /= float32(len(s.grades))
	return
}

func (s Student) compare(a Student) int {
	x := s.gradeAverage()
	x1 := a.gradeAverage()
	if x < x1 {
		return -1
	} else if x > x1 {
		return 1
	}

	return 0
}

func main() {
	s1 := NewStudent("Nicolas", "3ºA EM", "Mathematics", 5, 6, 8, 10, 3)
	s2 := NewStudent("Rodson", "3ºA EM", "Mathematics", 5, 6, 8, 10)

	fmt.Println(s1.gradeAverage())
	fmt.Println(s1.grades)
	fmt.Println(s2.gradeAverage())
	fmt.Println(s2.grades)
	fmt.Println(s1.compare(s2))

	fmt.Printf("\n%v", s1.subject)
	fmt.Printf("\n%v\n", s2.class)

	fmt.Println(s1.Enrollment.id)
	fmt.Println(s2.SchoolTests.id)
}
