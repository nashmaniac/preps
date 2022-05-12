package utils

import "math"

func Max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func Abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func Pow(a, b int) int {
	return int(math.Pow(float64(a), float64(b)))
}
