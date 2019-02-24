package powxn

import "math"

func myPow(x float64, n int) float64 {
	if n == 0 {
		return 1
	}

	if n == 1 {
		return x
	}

	if n < 0 {
		return 1 / myPow(x, -n)
	}

	if n%2 == 1 {
		return math.Pow(myPow(x, n/2), 2) * x
	} else {
		return math.Pow(myPow(x, n/2), 2)
	}
}
