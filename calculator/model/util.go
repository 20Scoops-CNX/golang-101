package model

import (
	"strconv"
)

func ConvertStringToInt(x string, y string) (n1 int64, n2 int64, err error) {
	n1, err = strconv.ParseInt(x, 10, 64)
	n2, err = strconv.ParseInt(y, 10, 64)
	if err != nil {
		return 0, 0, err
	}
	return n1, n2, nil
}

func GetNumberOperator(operator string) (number int64, err error) {
	choice, err := strconv.ParseInt(operator, 10, 0)
	if err != nil {
		return 0, err
	}
	return choice, nil
}

func GetResult(choice int64, n1 int64, n2 int64) (result int64) {
	if choice == 1 {
		return Plus(n1, n2)
	} else if choice == 2 {
		return Minus(n1, n2)
	} else if choice == 3 {
		return Multiply(n1, n2)
	}
	return Divide(n1, n2)
}
