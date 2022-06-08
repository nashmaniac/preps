package tree

import "math"

func Pow(a int, b int) int {
	return int(math.Pow(float64(a), float64(b)))
}
func Max(a int, b int) int {
	return int(math.Max(float64(a), float64(b)))
}
