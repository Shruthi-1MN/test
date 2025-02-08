package mathutil

import "math"

type MathOperations interface{
	Add(a, b int) int
	Subtract(a, b int)int
}

type Calculator struct{
	Name string
}

func (c Calculator) Add(a, b int)int{
	return a+b
}

func (c Calculator) Subtract(a, b int)int{
	return a-b
}

func (c Calculator) SquareRoot(x float64) float64{
	return math.Sqrt(x)
}