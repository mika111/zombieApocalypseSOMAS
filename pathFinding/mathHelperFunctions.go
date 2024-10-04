package pathfinding

import "math"

func chebyshev(xDiff, yDiff int) int {
	return min(abs(xDiff), abs(yDiff))
}

func abs(x int) int {
	if x < 0 {
		return -x
	} else {
		return x
	}
}

func min(x, y int) int {
	if x < y {
		return x
	} else {
		return y
	}
}

func sqrRootDistance(x, y int) float64 {
	var xSqr float64 = float64(x * x)
	var y_Sqr float64 = float64(y * y)
	return math.Sqrt(xSqr + y_Sqr)
}
