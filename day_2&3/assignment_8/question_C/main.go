package main

import (
	"project/model"
	"project/person"
)

func main() {
	person1 := model.Person{
		Name: "Mohammed Afthab",
		Age:  22,
	}
	person2 := model.Person{
		Name: "Poorvi",
		Age:  22,
	}
	person.PrintPerson(person1)
	person.PrintPerson(person2)
}
