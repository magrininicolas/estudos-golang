package main

import "fmt"

type Student struct {
	name         string
	class        string
	subject      string
	gradesNumber uint8
	grades       []float32
}

func NewStudent(name, class, subject string, grades ...float32) Student {
	maxGrades := 4
	if len(grades) < 4 {
		maxGrades = len(grades)
	}
	return Student{name, class, subject, uint8(maxGrades), grades[:maxGrades]}
}

func (s Student) gradeAverage() (soma float32) {
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
	fmt.Println(s2.gradeAverage())
	fmt.Println(s1.compare(s2))
}
