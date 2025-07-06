package main

import "errors"

func divide(divident, divisor int) (int, error) {
	if divisor == 0 {
		return 0, errors.New("Can't divide by zero")
	}
	return 0,nil
}