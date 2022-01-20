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

func (s *SillyNephewError) Error() string {
	return fmt.Sprintf("silly nephew, there cannot be %d cows", s.cows)
}

// DivideFood computes the fodder amount per cow for the given cows.
func DivideFood(weightFodder WeightFodder, cows int) (float64, error) {
	fodderAmount, err := weightFodder.FodderAmount()

	if cows < 0 {
		return 0, &SillyNephewError{cows}
	}

	if fodderAmount < 0 && (err == nil || err == ErrScaleMalfunction) {
		return 0, errors.New("negative fodder")
	}

	if cows == 0 {
		return 0, errors.New("division by zero")
	}

	if err == ErrScaleMalfunction && fodderAmount > 0 {
		return (fodderAmount * 2) / float64(cows), nil
	}

	if err != nil {
		return 0, err
	}

	return fodderAmount / float64(cows), nil
}
