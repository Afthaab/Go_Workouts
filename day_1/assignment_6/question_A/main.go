package main

import "fmt"

func main() {
	studentGrades := make(map[string]float32)
	studentGrades["Afthab"] = 10
	studentGrades["Jeevan"] = 9.8
	studentGrades["Poorvi"] = 9.9

	fmt.Println("The Grades are : ", studentGrades)
	delete(studentGrades, "Afthab")
	fmt.Println("The Grades are : ", studentGrades)
}
