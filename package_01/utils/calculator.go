package utils

func Plus(number1 int, number2 int) (result int) {
	result = number1 + number2
	return result
}

func Minus(number1 int, number2 int) (result int) {
	result = number1 - number2
	return result
}

func Multiply(number1 int, number2 int) (result int) {
	result = number1 * number2
	return result
}

func Divide(number1 int, number2 int) (result int) {
	if number2 == 0 {
		result = 0
	} else {
		result = number1 / number2
	}
	return result
}
