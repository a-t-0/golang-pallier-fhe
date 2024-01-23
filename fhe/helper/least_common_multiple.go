package fhe

import "math"

// AddTwo adds 2 to an incoming integer x.
// Source: https://stackoverflow.com/questions/5611751/gcd-and-lcm-relation
// Source: https://en.wikipedia.org/wiki/Least_common_multiple#Using_the_greatest_common_divisor
func Lcm(a int, b int) int {
	// TODO: assert p is a prime.
	// TODO: assert p is a prime.
	// TODO: assert lcm return value is an int.
	// TODO: implement.
	absoluteValue := math.Abs(float64(a * b))
	return absoluteValue / gcd(a, b)
}
