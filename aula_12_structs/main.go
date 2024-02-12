package main

import (
	"encoding/json"
	"fmt"
)

type Student struct {
	Name   string     `json:"name"`
	Grades [4]float32 `json:"grades"`
}

func changeName(s *Student) {
	s.Name = "Chaged name"
}

func main() {
	s1 := Student{Name: "Nicolas", Grades: [4]float32{5, 6, 8, 9}}
	s2 := Student{Name: "Lucas", Grades: [4]float32{10, 10, 10, 10}}

	fmt.Println(s1)
	fmt.Println(s2)

	s1Json, _ := json.Marshal(s1)

	fmt.Println(string(s1Json))
}
