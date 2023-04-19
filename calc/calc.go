package calc

import (
	"math"
)

func CalculateIsArmstrong(n int) bool {
	a := n / 100
	b := n % 100 / 10
	c := n % 10
	return n == int(math.Pow(float64(a), 3)+math.Pow(float64(b), 3)+math.Pow(float64(c), 3))
}

// Random function
func RandomFunction(n int) bool {
	if n == 7 {
		return true
	} else {
		return false
	}
}
