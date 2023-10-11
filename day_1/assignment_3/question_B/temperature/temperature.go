package temperature

func TempratureConverter(temp float64) float64 {
	return ((temp - 32) * 5) / 9
}
