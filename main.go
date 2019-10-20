package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

func main() {
	ricardo := person{firstName: "Ricardo", lastName: "Oliveira"}

	var nome person
	nome.firstName = "Marco"
	nome.lastName = "Oliveira"

	ricardoComercial := person{
		firstName: "RICARDO",
		lastName:  "OLIVEIRA",
		contact: contactInfo{
			email:   "teste@empresa.com",
			zipCode: 01007,
		},
	}

	fmt.Println("Dados Pessoal:")
	fmt.Printf("Meu nome é %s %s", ricardo.firstName, ricardo.lastName)
	fmt.Println("")
	fmt.Printf("%+v", nome)

	fmt.Println("")
	fmt.Println("")
	fmt.Println("Dados Comercial:")
	fmt.Printf("%+v", ricardoComercial)

}
