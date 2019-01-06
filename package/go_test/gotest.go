package gotest

import "errors"

func Division(a, b float64) (float64, error) {
	if 0 == b {
		return 0, errors.New("除数不能为0")
	}
	return a / b, nil
}
