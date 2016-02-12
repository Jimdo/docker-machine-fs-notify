package kmgMath

import (
	"math"
	"strconv"
)

//使用相对精度输出浮点
func Float64RoundToRelativePrec(f float64, prec int) float64 {
	s := strconv.FormatFloat(f, 'e', prec, 64)
	o, err := strconv.ParseFloat(s, 64)
	if err != nil {
		panic(err) //应该不可能
	}
	return o
}

//以一个基准数据,输出相对浮点精度
func Float64RoundToRelativePrecOnOneFloat(f float64, prec int, precBaseF float64) float64 {
	if precBaseF == 0 { //避免输出NaN
		return 0
	}
	absPrec := math.Floor(math.Log10(precBaseF)) - float64(prec)
	return f - math.Mod(f, math.Pow10(int(absPrec)))
}
