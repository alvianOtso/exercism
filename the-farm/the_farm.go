package thefarm

import (
	"errors"
	"fmt"
)

// TODO: define the 'DivideFood' function
func DivideFood(fodder FodderCalculator, cow int) (float64, error) {
	fodderAmount, err := fodder.FodderAmount(cow)
	if err != nil {
		return 0, err
	}

	fattening, err := fodder.FatteningFactor()
	if err != nil {
		return 0, err
	}

	totalFood := fodderAmount * fattening
	return totalFood / float64(cow), nil
}

// TODO: define the 'ValidateInputAndDivideFood' function
func ValidateInputAndDivideFood(fodder FodderCalculator, cow int) (float64, error) {
	if cow > 0 {
		return DivideFood(fodder, cow)
	}
	return 0, errors.New("invalid number of cows")
}

type InvalidCowsError struct {
	cow     int
	message string
}

func (e *InvalidCowsError) Error() string {
	return fmt.Sprintf("%d cows are invalid: %s", e.cow, e.message)
}

// TODO: define the 'ValidateNumberOfCows' function
func ValidateNumberOfCows(cows int) error {
	cowsErr := &InvalidCowsError{
		cow: cows,
	}
	if cows < 0 {
		cowsErr.message = "there are no negative cows"
		return cowsErr
	} else if cows == 0 {
		cowsErr.message = "no cows don't need food"
		return cowsErr
	}

	return nil
}

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.
