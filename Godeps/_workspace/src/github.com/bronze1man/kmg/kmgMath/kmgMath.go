package kmgMath

import (
	"math"
	"strconv"
)

func FloorToInt(x float64) int {
	return int(math.Floor(x))
}

func CeilToInt(x float64) int {
	return int(math.Ceil(x))
}

func MustParseFloat(s string) float64 {
	f, err := strconv.ParseFloat(s, 64)
	if err != nil {
		panic(err)
	}
	return f
}
