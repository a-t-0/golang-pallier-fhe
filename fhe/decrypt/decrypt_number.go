package decrypt

import (
	"errors"
	"fmt"
	"math/big"

	"github.com/a-t-0/golang-pallier-fhe/fhe/keygen"
)

func DecryptNumber(encryptedMessage *big.Int, lambda *big.Int, n *big.Int, mu *big.Int) *big.Int {

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
	lValue := keygen.ComputeL(mod, n)
	thirdProduct := new(big.Int).Mul(lValue, mu)

	var decryptedMessage big.Int
	decryptedMessage.Mod(thirdProduct, n)

	return &decryptedMessage
}
