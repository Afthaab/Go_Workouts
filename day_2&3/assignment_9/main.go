package main

import "fmt"

type Shape interface {
	Area() float32
}

type Circle struct {
	Radius float32
}

type Rectangle struct {
	Length float32
	width  float32
}

func (C Circle) Area() float32 {
	return 3.14 * C.Radius * C.Radius
}

func (R Rectangle) Area() float32 {
	return R.Length * R.width
}

func Opreration(S Shape) {
	fmt.Println("The Area is : ", S.Area())
}

func main() {
	c := Circle{
		Radius: 7.7,
	}
	r := Rectangle{
		width:  10,
		Length: 11,
	}
	Opreration(c)
	Opreration(r)

}
