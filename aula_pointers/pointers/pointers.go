package main

import "fmt"

func doublePointer(a *int) {
	*a = *a * 2
}

func double(b int) int {
	return b * 2
}

func main() {
	a := 10
	b := 15
	fmt.Println(a)
	fmt.Println(b)

	doublePointer(&a)
	fmt.Println(a)
	b = double(b)
	fmt.Println(b)
}
