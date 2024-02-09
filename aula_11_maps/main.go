package main

import "fmt"

func main() {
	var idadeAlunos map[string]int
	idadeAlunos = map[string]int{
		"Nicolas": 21,
	}

	idadeAlunos["Rodson"] = 51 //Add an element, if the element already exists, itll be removed
	idadeAlunos["Lucas"] = 21
	fmt.Println(idadeAlunos)

	nicolas, ok := idadeAlunos["Nicolas"] //Get an element (v1 = value, v2 = true if exists/false otherwise)
	if ok {
		fmt.Printf("Idade de Nicolas: %d\n", nicolas)
	}

	idadeAlunos["Nicolas"] = 20
	fmt.Println(idadeAlunos)
	idadeAlunos["Nicolas"] = 21
	fmt.Println(idadeAlunos)

	delete(idadeAlunos, "Lucas")
	fmt.Println(idadeAlunos)

	populacaoPorUF := make(map[string]int, 27) //Initialize a map using make()
	populacaoPorUF["SP"] = 44_400_000
	populacaoPorUF["RJ"] = 16_460_000
	fmt.Println(populacaoPorUF)

	fmt.Println(len(idadeAlunos)) //len()
	fmt.Println(len(populacaoPorUF))

}
