package main

import "fmt"

type Rectangle struct {
	Width  float32
	Length float32
}

func (r Rectangle) PerimeterOfARectangle() float32 {
	return (2 + (r.Length + r.Width))
}
func main() {
	R1 := Rectangle{
		Width:  10,
		Length: 7.7,
	}
	perimeter := R1.PerimeterOfARectangle()
	fmt.Println("The perimeter is : ", perimeter)
}
