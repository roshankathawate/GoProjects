package numberutil

import (
	"errors"
	"math"
)

func Add(val1 float64, val2 float64) (float64, error) {
	if math.IsNaN(val1) || math.IsNaN(val2){
		return 0.0, errors.New("Invalid numbers.")
	}
	return val1 + val2, nil
}

func Subtract(val1 float64, val2 float64) (float64, error) {
	 if math.IsNaN(val1) || math.IsNaN(val2){
		return 0.0, errors.New("Invalid numbers.")
	}
	return val1 - val2, nil
}

func Multiply(val1 float64, val2 float64) (float64, error) {
	 if math.IsNaN(val1) || math.IsNaN(val2){
		return 0.0, errors.New("Invalid numbers.")
	}
	return val1 * val2, nil
}




