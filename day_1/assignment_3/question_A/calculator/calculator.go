package calculator

func SquareOfANumber(num int) int {
	return num * num
}

func SumOfNumbers(num1 int, num2 int) int {
	return num1 + num2
}

func DifferenceOfNumbers(num1 int, num2 int) int {
	return num1 - num2
}

func ProductOfNumbers(num1 int, num2 int) int {
	return num1 * num2
}

func DivisionOfNumbers(num1 int, num2 int) (int, int) {
	return num1 / num2, num2 % num2
}
