package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Client struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	CPF   int    `json:"cpf"`
}

func (c Client) Show() {
	fmt.Println("Showing client by the method:", c.Name)
}

type ForeignerClient struct {
	Client
	Country string `json:"country"`
}

func main() {

	client := Client{
		Name:  "Nicolas",
		Email: "nicolas@email.com",
		CPF:   47159866800,
	}

	client2 := ForeignerClient{
		Client:  Client{Name: "Rodson", Email: "rodson@email.com", CPF: 15838626844},
		Country: "Canada",
	}

	fmt.Println(client)
	fmt.Println(client2)

	client.Show()
	client2.Show()

	clientJson, err := json.Marshal(client2)

	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println(string(clientJson))

	jsonClient3 := `{"name":"Rodson","email":"rodson@email.com","cpf":15838626844,"country":"Canada"}`
	client3 := ForeignerClient{}

	json.Unmarshal([]byte(jsonClient3), &client3)

	fmt.Println(client3)
	client3.Show()
}
