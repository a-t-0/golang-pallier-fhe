package encrypt

import (
	"errors"
	"fmt"
	"math/big"
)

func EncryptNumber(message *big.Int, rand_nr *big.Int, n *big.Int, g *big.Int) *big.Int {

	if message.Cmp(big.NewInt(0)) < 0 {
		err := errors.New("the message integer should be 0 or larger")
		fmt.Println("Error:", err.Error())
	}
	if message.Cmp(n) >= 0 {
		err := errors.New("the message integer should be smaller than n")
		fmt.Println("Error:", err.Error())
	}

	if rand_nr.Cmp(big.NewInt(0)) <= 0 {
		err := errors.New("the random integer should be larger than 0")
		fmt.Println("Error:", err.Error())
	}
	if rand_nr.Cmp(n) >= 0 {
		err := errors.New("the random integer should be smaller than n")
		fmt.Println("Error:", err.Error())
	}

	// var firstExponent big.Int = int(math.Pow(float64(g), float64(message)))
	firstExponent := new(big.Int).Exp(g, message, nil)

	// var secondExponent int = int(math.Pow(float64(rand_nr), float64(n)))
	secondExponent := new(big.Int).Exp(rand_nr, n, nil)
	fmt.Printf("firstExponent %d\n", firstExponent)
	fmt.Printf("secondExponent %d\n", secondExponent)

	// var encrypted_message int = (firstExponent * secondExponent) % (n * n)
	var encryptedMessage big.Int
	encryptedMessage.Mul(firstExponent, secondExponent)
	encryptedMessage.Mod(&encryptedMessage, new(big.Int).Mul(n, n))

	return &encryptedMessage
}
