package main

import "fmt"

func soma(a ...int) (soma int) {
	for _, v := range a {
		soma += v
	}
	return
}

func main() {

	//Pointer
	a := []int{1}
	fmt.Printf("%p\n", a)

	//Len
	b := []int{1, 2, 3}
	fmt.Println(len(b))

	//Append
	b = append(b, 5, 6, 7)
	fmt.Println(b, len(b))

	//Cap
	fmt.Println(b, cap(b))
	b = append(b, 1)
	fmt.Println(b, cap(b))

	//Make
	c := make([]int, 0)
	c = append(c, 1)
	fmt.Println(c)

	//Variadic
	somaB := soma(b...)
	fmt.Println(b)
	fmt.Println(somaB)
}
