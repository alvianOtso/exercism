package grains

import (
	"errors"
)

func Square(number int) (uint64, error) {
	if number > 0 && number < 65 {
		return 1 << (number - 1), nil
	}
	return 0, errors.New("invalid")
}

func Total() uint64 {
	return 1<<64 - 1
}
