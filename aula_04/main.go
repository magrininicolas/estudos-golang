package main

import "fmt"

func main() {
	altura1 := 1.75
	altura2 := 1.68
	var altura3 float32
	var altura4 float64 = 1.82
	fmt.Printf("%2.3f\n", altura1)
	fmt.Printf("%2.3f\n", altura2)
	fmt.Printf("%2.3f\n", altura3)
	fmt.Printf("%2.3f\n", altura4)
	fmt.Printf("%T\n", altura3)
}
