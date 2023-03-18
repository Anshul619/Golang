package main

import (
	"errors"
	"fmt"
)

// This file contains types used in the exercise but should not be modified.

// WeightFodder returns the amount of available fodder.
type WeightFodder interface {
	FodderAmount() (float64, error)
}

// ErrScaleMalfunction indicates an error with the scale.
var ErrScaleMalfunction = errors.New("sensor error")

// See types.go for the types defined for this exercise.

// TODO: Define the SillyNephewError type here.

type SillyNephewError struct {
	cows int
}

func (s SillyNephewError) Error() string {
	return fmt.Sprintf("silly nephew, there cannot be %d cows", s.cows)
}

// DivideFood computes the fodder amount per cow for the given cows.
func DivideFood(weightFodder WeightFodder, cows int) (float64, error) {

	if cows == 0 {
		return 0, errors.New("division by zero")
	}

	f, err := weightFodder.FodderAmount()

	if f < 0 && (err == ErrScaleMalfunction || err == nil) {
		return 0, errors.New("negative fodder")
	}

	if cows < 0 {
		return 0, SillyNephewError{cows}
	}

	if err == nil {
		return f / float64(cows), nil
	}

	if err == ErrScaleMalfunction && f > 0 {
		return (f * 2) / float64(cows), nil
	}

	return 0, err
}

// func main() {
// 	// twentyFodderNoError says there are 20.0 fodder
// 	fodder, err := DivideFood(twentyFodderNoError, 10)
// 	// fodder == 2.0
// 	// err == nil

// }
