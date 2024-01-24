package fhe

import (
	"errors"
	"fmt"
)

// gcd returns the greatest common divisor of two integers x and y, using
// a Euclidean algorithm, which is different than Euclids algorithm.
// Source: https://en.wikipedia.org/wiki/Euclidean_algorithm#Procedure
// It starts with 2 input numbers, and then computes 2 new numbers until one
// of the two has reached 0. In that case the other/second to last number is
// the greatest common divisor. (E.g. for 8 and 12, gcd=4). The function to
// compute the two new numbers is:
//
//	r_{k-2}=q_k * r_{k-1}+r_k
//
// with:
//
//	r_{k-1} > r_{k} >= 0
//
// In this algorithm, q_k and r_k need to be searched iteratively using Euclids
// division lemma: https://en.wikipedia.org/wiki/Euclidean_division.
func Gcd(a int, b int) int {
	var rk2 int
	var rk1 int
	// var rk0 int
	if a > b {
		rk2 = a
		rk1 = b
	} else if a < b {
		rk2 = b
		rk1 = a
	} else if a == b {
		//TODO: return shortcut value.
		err := errors.New("TODO: implement edge case of two computing gcd for 2 identical integers.")
		fmt.Println("Error:", err.Error())
	} else {
		err := errors.New("If statements were incomplete.")
		fmt.Println("Error:", err.Error())
	}

	// Calling findQkAndRk until the remainder (r0) becomes 0.
	for rk1 != 0 {
		q, rk0 := findQkAndRk(rk2, rk1)
		fmt.Printf("rk2 = %d, rk1 = %d, q = %d, r0 = %d\n", rk2, rk1, q, rk0)
		rk2, rk1 = rk1, rk0
	}
	return rk2

}

// Finds the values of q_k and r_k for a given pair of r_{k-2} and r_{k-1}.
func findQkAndRk(rk2 int, rk1 int) (int, int) {
	// TODO: assert rk2>rk1.
	// Integer division (equivalent to // in Python).
	q := rk2 / rk1

	// Modulo (equivalent to % in Python)
	rk0 := rk2 % rk1

	fmt.Printf("Integer division: %d\n", q)
	fmt.Printf("Modulo: %d\n", rk0)
	return q, rk0
}
