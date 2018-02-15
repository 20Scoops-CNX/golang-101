package model

func Plus(number1 int64, number2 int64) (result int64) {
	result = number1 + number2
	return result
}

func Minus(number1 int64, number2 int64) (result int64) {
	result = number1 - number2
	return result
}

func Multiply(number1 int64, number2 int64) (result int64) {
	result = number1 * number2
	return result
}

func Divide(number1 int64, number2 int64) (result int64) {
	if number2 == 0 {
		result = 0
	} else {
		result = number1 / number2
	}
	return result
}
