package main

import "fmt"

type Pessoa struct {
	idade int
}

func (p *Pessoa) aniversario() {
	if p != nil {
		p.idade++
	}
}

func main() {
	var soma func(a, b int) int
	fmt.Println(soma == nil)

	var notas []float64
	fmt.Println(notas == nil)
	fmt.Println(len(notas))
	for _, nota := range notas {
		fmt.Println(nota)
	}
	notas = append(notas, 5.6)
	fmt.Println(notas)

	popPorUF := make(map[string]int, 5)
	fmt.Println(popPorUF == nil)
	fmt.Println(popPorUF["SP"])
	popPorUF["SP"] = 5000000
	if v, ok := popPorUF["SP"]; ok {
		fmt.Println(v)
	} else {
		fmt.Println("Não encontrado")
	}

}
