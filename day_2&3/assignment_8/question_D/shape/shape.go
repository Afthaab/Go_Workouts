package shape

import "project1/models"

func AreaOfCircle() float32 {
	c1 := models.Circle{
		Radius: 9.9,
	}
	return 3.14 * c1.Radius

}
