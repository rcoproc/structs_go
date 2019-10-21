package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contactInfo
}

func main() {
	ricardo := person{firstName: "Ricardo", lastName: "Oliveira"}

	var nome person

	nome.updateName("Marco Antonio")
	nome.updateLastname("Oliveira")

	ricardoComercial := person{
		firstName: "RICARDO",
		lastName:  "OLIVEIRA",
		contactInfo: contactInfo{
			email:   "teste@empresa.com",
			zipCode: 01007,
		},
	}

	fmt.Println("Dados variável nome:")
	nome.print()
	fmt.Println("")
	fmt.Println("")
	fmt.Println("Dados Pessoal:")
	fmt.Printf("Meu nome é %s %s", ricardo.firstName, ricardo.lastName)
	fmt.Println("")
	nome.print()

	fmt.Println("")
	fmt.Println("")
	fmt.Println("Dados Comercial:")
	fmt.Printf("%+v", ricardoComercial)
	ricardoComercial.print()

}

func (p person) print() {
	fmt.Printf("%+v", p)
}

// values parameters default by value(copy of data types) in go
// func (p person) updateName(newFirstName string) {
func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}

func (pointerToPerson *person) updateLastname(newLastName string) {
	(*pointerToPerson).lastName = newLastName
}
