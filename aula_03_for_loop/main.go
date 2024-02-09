package main

import (
	"fmt"
	"math/rand"
)

func main() {
	// i := 1
	// for i < 11 {
	// 	fmt.Println("Hello")
	// 	i++
	// 	time.Sleep(100 * time.Millisecond)
	// }

	numeroSecreto := 27
	i := 0

	for {
		sorteado := rand.Intn(99) + 1

		if sorteado > numeroSecreto {
			fmt.Printf("%v é maior que o número secreto\n", sorteado)
		} else if sorteado < numeroSecreto {
			fmt.Printf("%v é menor que o número secreto\n", sorteado)
		} else {
			fmt.Printf("O número secreto é %v\n", sorteado)
			fmt.Printf("Número de tentativas: %v\n", i)
			break
		}
		i++
	}

}
