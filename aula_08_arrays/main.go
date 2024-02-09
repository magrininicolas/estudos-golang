package main

import "fmt"

func main() {
	notas := [5]float64{10, 9.4, 8.2}

	media := 0.0

	for i := 0; i < len(notas); i++ {
		media += notas[i]
	}
	fmt.Printf("Média: %.2f\n", media/float64(len(notas)))

	media = 0.0
	for _, nota := range notas {
		media += nota
	}

	fmt.Printf("Média: %.2f\n", media/float64(len(notas)))

}
