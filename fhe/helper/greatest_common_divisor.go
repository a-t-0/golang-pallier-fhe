package helper

import (
	"errors"
	"fmt"
	"math/big"
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
func Gcd(a *big.Int, b *big.Int) *big.Int {
	var rk2 *big.Int
	var rk1 *big.Int
	// var rk0 int
	// if a > b {
	if a.Cmp(b) > 0 {
		rk2 = a
		rk1 = b
		// } else if a < b {
	} else if a.Cmp(b) < 0 {
		rk2 = b
		rk1 = a
		// } else if a == b {
	} else if a.Cmp(b) == 0 {
		//TODO: return shortcut value.
		err := errors.New("todo: implement edge case of two computing gcd for 2 identical integers")
		fmt.Println("Error:", err.Error())
	} else {
		err := errors.New("if statements were incomplete")
		fmt.Println("Error:", err.Error())
	}

	// Calling findQkAndRk until the remainder (r0) becomes 0.
	q := new(big.Int)
	rk0 := new(big.Int)
	for rk1.Cmp(big.NewInt(0)) != 0 {
		q, rk0 := findQkAndRk(rk2, rk1)
		fmt.Printf("rk2 = %d, rk1 = %d, q = %d, r0 = %d\n", rk2, rk1, q, rk0)
		rk2, rk1 = rk1, rk0
	}
	fmt.Printf("rk1=%d", rk1)
	fmt.Printf("q=%d", q)
	fmt.Printf("rk0=%d", rk0)
	return rk2

}

// Finds the values of q_k and r_k for a given pair of r_{k-2} and r_{k-1}.
func findQkAndRk(rk2 *big.Int, rk1 *big.Int) (*big.Int, *big.Int) {
	// TODO: assert rk2>rk1.
	if rk1.Cmp(big.NewInt(0)) <= 0 {
		err := errors.New("division by 0")
		fmt.Println("Error:", err.Error())
	}
	if rk2.Cmp(big.NewInt(0)) <= 0 {
		err := errors.New("division by 0")
		fmt.Println("Error:", err.Error())
	}
	// Integer division (equivalent to // in Python).
	// q := rk2 / rk1
	q := new(big.Int)
	q.Div(rk2, rk1)

	// Modulo (equivalent to % in Python)
	// rk0 := rk2 % rk1
	rk0 := new(big.Int)
	rk0.Mod(rk2, rk1)

	fmt.Printf("Integer division: %d\n", q)
	fmt.Printf("Modulo: %d\n", rk0)
	return q, rk0
}
