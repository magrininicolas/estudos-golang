package main

import "fmt"

func hello(name string) string {
	return "Hello, " + name
}

func soma(a, b int) (int, string) {
	defer fmt.Println("Fim da execução da Soma")
	res := a + b
	resFmt := fmt.Sprintf("%d + %d = %d", a, b, res)

	return res, resFmt
}

func subtrai(a, b int) (res int, resFmt string) {
	res = a - b
	resFmt = fmt.Sprintf("%d - %d = %d", a, b, res)
	return
}

func main() {
	resSoma, formatSoma := soma(5, 10)
	resSubtrai, formatSubtrai := subtrai(5, 10)

	fmt.Println(hello("Nicolas"))
	fmt.Println(resSoma)
	fmt.Println(formatSoma)
	fmt.Println(resSubtrai)
	fmt.Println(formatSubtrai)
}
