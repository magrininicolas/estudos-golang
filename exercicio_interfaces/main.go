package main

import (
	"fmt"
	"learningInterfaces/shapes"
	"os"
)

func showArea(s shapes.Shape) {
	fmt.Printf("The area is equal to %.4f\n", s.Area())
}

func menu() int {
	op := 0

	fmt.Println("Choice an shape to show the area")
	fmt.Println("1 - Square")
	fmt.Println("2 - Rectangle")
	fmt.Println("3 - Circle")
	fmt.Println("4 - Triangle")
	fmt.Println("5 - Equilateral Triangle")
	fmt.Println("6 - Exit")
	fmt.Scan(&op)

	return op
}

func getInput(prompt string) float64 {
	var value float64
	fmt.Println(prompt)
	fmt.Scan(&value)
	return value
}

func registerShape(op int) shapes.Shape {
	switch op {
	case 1:
		side := getInput("Enter the square's side:")
		return shapes.Square{Side: side}
	case 2:
		width := getInput("Enter the rectangle's width: ")
		height := getInput("Enter the rectangle's height:")
		return shapes.Rectangle{Width: width, Height: height}
	case 3:
		radius := getInput("Enter the circle's radius: ")
		return shapes.Circle{Radius: radius}
	case 4:
		base := getInput("Enter the triangle's base: ")
		height := getInput("Enter the triangle's height: ")
		return shapes.Triangle{Base: base, Height: height}
	case 5:
		side := getInput("Enter the equilateral triangle's side: ")
		return shapes.EqtTriangle{Triangle: shapes.Triangle{Height: side}}
	case 6:
		os.Exit(0)
	default:
		fmt.Println("Enter a valid option!")
	}

	return nil
}

func main() {
	op := 0
	for op != 6 {
		op = menu()
		shape := registerShape(op)
		if op < 6 {
			showArea(shape)
		}
	}
}
