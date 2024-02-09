package main

import "fmt"

func main() {

	idade := 25
	temCarteiraDeMotorista := false

	if idade >= 18 && temCarteiraDeMotorista {
		fmt.Println("Usuário pode dirigir")
	} else {
		fmt.Println("Usuário não tem permissão para dirigir")
	}

	switch idade {
	case 10:
		fmt.Println("10 anos")
	case 18:
		fmt.Println("18 anos")
	}

	switch {
	case idade <= 10:
		fmt.Println("Criança")
	case idade > 12 && idade < 18:
		fmt.Println("Adolescente")
	default:
		fmt.Println("Adulto")
	}
}
