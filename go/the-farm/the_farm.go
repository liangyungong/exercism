package thefarm

import (
	"errors"
	"fmt"
)

// See types.go for the types defined for this exercise.

// TODO: Define the SillyNephewError type here.
type SillyNephewError struct {
	cows int
}

func (e *SillyNephewError) Error() string {
	return fmt.Sprintf("silly nephew, there cannot be %d cows", e.cows)
}

// DivideFood computes the fodder amount per cow for the given cows.
func DivideFood(weightFodder WeightFodder, cows int) (float64, error) {
	w, err := weightFodder.FodderAmount()

	if err == ErrScaleMalfunction && w > 0 {

		return 2 * w / float64(cows), nil
	}

	if w < 0 && (err == ErrScaleMalfunction || err == nil) {
		return 0.0, errors.New("negative fodder")
	}

	if cows == 0 {
		return 0.0, errors.New("division by zero")
	}

	if cows < 0 {
		return 0.0, &SillyNephewError{cows}
	}

	if err != nil {
		return 0.0, err
	}

	return w / float64(cows), nil
}
