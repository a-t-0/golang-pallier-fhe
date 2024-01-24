package decrypt

import (
	"errors"
	"fmt"
	"math/big"
)

func DecryptNumberBackup(encryptedMessage *big.Int, lambda *big.Int, n *big.Int, mu *big.Int) *big.Int {

	if encryptedMessage.Cmp(big.NewInt(0)) < 0 {
		err := errors.New("the encrypted message integer should be 0 or larger")
		fmt.Println("Error:", err.Error())
	}
	if encryptedMessage.Cmp(new(big.Int).Mul(n, n)) >= 0 {
		err := errors.New("the encrypted message integer should be smaller than n^2")
		fmt.Println("Error:", err.Error())
	}

	firstExponent := new(big.Int).Exp(encryptedMessage, lambda, nil)
	secondExponent := new(big.Int).Exp(n, big.NewInt(2), nil)
	mod := new(big.Int).Mul(firstExponent, secondExponent)

	fmt.Printf("firstExponent %d\n", firstExponent)
	fmt.Printf("secondExponent %d\n", secondExponent)
	fmt.Printf("mod %d\n", mod)

	// lValue := keygen.ComputeL(int(mod))
	lValue := ComputeL(mod, n)
	thirdProduct := new(big.Int).Mul(lValue, mu)

	var decryptedMessage big.Int
	decryptedMessage.Mod(thirdProduct, n)

	return &decryptedMessage
}

func DecryptNumber(encryptedMessage *big.Int, lambda *big.Int, n *big.Int, mu *big.Int) *big.Int {
	// Calculate c^lambda mod n^2
	power := new(big.Int).Exp(encryptedMessage, lambda, nil)
	cLambdaModN2 := new(big.Int).Mod(power, new(big.Int).Exp(n, big.NewInt(2), nil))

	// Calculate L(c^lambda mod n^2)
	l := ComputeL(cLambdaModN2, n)

	// Calculate m = L(c^lambda mod n^2) * μ mod n
	m := new(big.Int).Mul(l, mu)
	m.Mod(m, n)

	return m
}

// ComputeL computes the L function as named by Pascal Pallier in the FHE
// presentation. L(x) = [(x-1)/n], THIS IS NOT DIVIDE. This is: how often does
// the bottom fit in the top, like [2/3]=0, [5/3=1], [7/3=2]. A more
// mathematical description is:  the largest integer value v ≥ 0 to satisfy the
// relation a >= v*b. The Python equivalent is 11 // 4 (=2).
func ComputeL(x *big.Int, n *big.Int) *big.Int {
	// Ensure x is non-negative
	if x.Sign() < 0 {
		panic("x must be non-negative")
	}

	// Ensure n is positive
	if n.Sign() <= 0 {
		panic("n must be positive")
	}

	// Compute the result using big.Int's Div method
	// TODO: Important: verify you should also subtract 1 at the bottom.
	result := new(big.Int).Sub(x, big.NewInt(1)).Div(new(big.Int).Sub(x, big.NewInt(1)), n)

	// Convert the result to an int
	return result
}
