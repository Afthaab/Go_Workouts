package main

import "fmt"

type Student struct {
	Name string
	Age  int
}

func main() {
	s := Student{
		Name: "Mohammed Afthab",
		Age:  22,
	}
	fmt.Println("The struct before = ", s)
	s = UpdateTheStruct(&s, 23, "poorvi")
	fmt.Println("The updated struct is = ", s)

}

func UpdateTheStruct(s *Student, age int, name string) Student {
	s.Age = age
	s.Name = name
	return *s
}
