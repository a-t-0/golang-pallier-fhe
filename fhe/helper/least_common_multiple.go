package helper

import (
	"math/big"
)

// AddTwo adds 2 to an incoming integer x.
// Source: https://stackoverflow.com/questions/5611751/gcd-and-lcm-relation
// Source: https://en.wikipedia.org/wiki/Least_common_multiple#Using_the_greatest_common_divisor
func Lcm(a *big.Int, b *big.Int) *big.Int {

	// TODO: assert p is a prime.
	// TODO: assert p > 0.
	// TODO: assert q is a prime.
	// TODO: assert q > 0.

	// TODO: assert lcm return value is an int.
	// TODO: implement.
	top := new(big.Int)
	top.Mul(a, b)

	result := new(big.Int)
	result.Div(top, Gcd(a, b))

	return result
}
