package main

import (
	"fmt"
	"strings"
)

type Dias []string

func Uppercase(diasDaSemana []string) {
	for i := range diasDaSemana {
		diasDaSemana[i] = strings.ToUpper(diasDaSemana[i])
	}
}

func (d Dias) Uppercase() {
	for i := range d {
		d[i] = strings.ToUpper(d[i])
	}
}

func main() {

	diasDaSemana := Dias{
		"Segunda",
		"Terça",
		"Quarta",
		"Quinta",
		"Sexta",
		"Sábado",
		"Domingo",
	}
	fmt.Println(diasDaSemana)

	diasDaSemana.Uppercase()
	fmt.Println(diasDaSemana)
}
