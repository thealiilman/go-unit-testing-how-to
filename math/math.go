package math

func Add(x int, y int) int {
	return x + y
}

func Subtract(x int, y int) int {
	return x - y
}

func Divide(x int, y int) float64 {
	if x == 0 || y == 0 {
		return 0.0
	}

	return float64(x / y)
}

func Multiply(x int, y int) float64 {
	return float64(x * y)
}
