package main

import "fmt"

type animal interface {
	MakeSound() string
}

type Dog struct {
	Name  string
	Breed string
}

type Cat struct {
	Name  string
	Breed string
}

func (D Dog) MakeSound() string {
	return D.Name + " of breed " + D.Breed + " maked woooo sound"
}

func (C Cat) MakeSound() string {
	return C.Name + " of breed " + C.Breed + " maked woooo sound"
}

func Sound(a animal) {
	fmt.Println(a.MakeSound())
}

func main() {
	D1 := Dog{
		Name:  "Rex",
		Breed: "Labrodog",
	}
	C1 := Cat{
		Name:  "Chinnu",
		Breed: "Pomeranian",
	}
	Sound(D1)
	Sound(C1)

}
