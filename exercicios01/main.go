package main

import (
	"fmt"
	"os"
	"strings"
)

type Student struct {
	Name   string
	Grades [4]float32
}

var studentCounter int

func (s Student) calcAverage() float32 {
	var sum float32 = 0.0
	for _, v := range s.Grades {
		sum += v
	}

	return float32(sum / 4.0)
}

func menu() int {
	op := 0
	fmt.Println("1 - Add a student")
	fmt.Println("2 - Show report card")
	fmt.Println("3 - Exit")

	fmt.Println("Choose an option")
	fmt.Scanln(&op)
	return op
}

func addStudent(students map[string]Student) (string, [4]float32) {
	var name string
	var grades [4]float32

	fmt.Println("Student name: ")
	fmt.Scanln(&name)

	_, ok := students[name]

	name = verifiesSameName(ok, name)

	fmt.Println(name, "grades: ")
	for i := 0; i < 4; i++ {
		var grade float32
		fmt.Printf("Grade #%d: ", i+1)
		fmt.Scanln(&grade)
		grades[i] = verifiesGrade(grade)
	}

	return name, grades
}

func verifiesGrade(grade float32) float32 {
	for grade < 0 || grade > 10 {
		fmt.Println("Try again")
		fmt.Scanln(&grade)
	}

	return grade
}

func verifiesSameName(ok bool, name string) string {
	if ok {
		equalName := name
		for strings.EqualFold(name, equalName) {
			fmt.Println("Student name: ")
			fmt.Scanln(&name)
		}
	}

	return name
}

func main() {
	students := make(map[string]Student)
	op := 0

	for op != 3 {

		op = menu()

		switch op {
		case 1:
			name, grades := addStudent(students)

			students[name] = Student{Name: name, Grades: grades}
		case 2:
			var name string
			fmt.Println("What student? ")
			fmt.Scanln(&name)

			student, exists := students[name]

			if exists {
				fmt.Printf("%s's informations: \n", student.Name)
				fmt.Println("Grades:", student.Grades)
				fmt.Println("Average:", student.calcAverage())
			}
		case 3:
		default:
			os.Exit(0)
		}
	}

	fmt.Println("Final of the program!")
}
