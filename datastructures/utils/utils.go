package utils

import "math"

func Max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func Pow(a, b int) int {
	return int(math.Pow(float64(a), float64(b)))
}
