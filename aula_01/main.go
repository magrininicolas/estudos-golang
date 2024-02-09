package main

import (
	"fmt"
	"math/rand"
)

const a = 2

func main() {
	fmt.Println("Hello, World!")
	fmt.Printf("2 + 2 = %v\n", a+a)

	guessNumber := rand.Intn(10)
	fmt.Printf("Número gerado: %v", guessNumber)
}
