package person

import (
	"fmt"
	"project/model"
)

func PrintPerson(person model.Person) {
	fmt.Println("The details : ")
	fmt.Println(person)
}
