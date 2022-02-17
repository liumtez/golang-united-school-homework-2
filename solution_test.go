package square

import "testing"

func TestSqare(t * testing.T) {
	println(CalcSquare(10.0, SidesTriangle))
	println(CalcSquare(10.0, SidesSquare))
	println(CalcSquare(10.0, SidesCircle))
}