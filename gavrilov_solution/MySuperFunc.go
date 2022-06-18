package gavrilov_solution

import (
	"math"
)

// MySuperFuncImpl MySuperFunc - реализация
// 1. `n==0` -> `x1`
// 2. `n==1` -> `x1 * x2`
// 3. `n>1` -> `f(x1, x2, n-2) * f(x1, x2, n-1)`
func MySuperFuncImpl(x1 float64, x2 float64, n uint8) float64 {
	switch n {
	case 0:
		return x1
	default:
		if x1 == 0 || x2 == 0 {
			return 0
		}
		if x1 == 1 && x2 == 1 {
			return 1
		}
		var acc float64 = x1
		var res float64 = x1 * x2
		for i := uint8(2); i <= n; i++ {
			acc, res = res, acc*res
		}
		return res
	}
}

func Fib(n uint8) uint32 {
	switch n {
	case 0:
		return 0
	default:
		var prev uint32 = 0
		var res uint32 = 1
		for i := uint8(2); i <= n; i++ {
			res, prev = res+prev, res
		}
		return uint32(res)
	}
}

func FillFibs() [50]uint32 {
	var res [50]uint32
	for i := uint8(0); i < 50; i++ {
		res[i] = Fib(i)
	}
	return res
}

var fibs [50]uint32 = FillFibs()

func Power(x float64, p uint32) float64 {
	switch p {
	case 0:
		return 1
	case 1:
		return x
	default:
		var bits uint8 = 0
		for i := p; i > 0; i >>= 1 {
			bits++
		}
		var acc = x
		var total float64 = 1
		for i := uint8(0); i < bits; i++ {
			if i > 0 {
				acc *= acc
			}
			if p&(1<<i) > 0 {
				total *= acc
			}
		}
		return total
	}
}

func MySuperFuncImpl2(x1 float64, x2 float64, n uint8) float64 {
	return math.Pow(x1, float64(fibs[n+1])) * math.Pow(x2, float64(fibs[n]))
}

func MySuperFuncImpl3(x1 float64, x2 float64, n uint8) float64 {
	return math.Pow(x1*x2, float64(fibs[n])) * math.Pow(x1, float64(fibs[n+1]-fibs[n]))
}

func MySuperFuncImpl4(x1 float64, x2 float64, n uint8) float64 {
	return Power(x1, fibs[n+1]) * Power(x2, fibs[n])
}

func MySuperFuncImpl5(x1 float64, x2 float64, n uint8) float64 {
	return Power(x1*x2, fibs[n]) * Power(x1, fibs[n+1]-fibs[n])
}
