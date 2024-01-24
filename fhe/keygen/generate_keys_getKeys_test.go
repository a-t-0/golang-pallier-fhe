package keygen

import (
	"math/big"
	"testing"

	"github.com/a-t-0/golang-pallier-fhe/fhe/decrypt"
	"github.com/a-t-0/golang-pallier-fhe/fhe/encrypt"
)

// Test for the AddTwo function
func TestGetKeys(t *testing.T) {
	// Test case 1: Adding 2 to 3 should result in 5
	var message *big.Int = big.NewInt(123)

	n, g, lambda, mu := getKeys()

	var rand_nr *big.Int = big.NewInt(666)
	var encryptedMessage *big.Int = encrypt.EncryptNumber(message, rand_nr, n, g)
	var expectedOne *big.Int = big.NewInt(25889)

	if encryptedMessage.Cmp(expectedOne) != 0 {
		t.Errorf("Expected %d, but got encryptedMessage %d", expectedOne, encryptedMessage)
	}

	var decryptedMessage *big.Int = decrypt.DecryptNumber(encryptedMessage, lambda, n, mu)
	if decryptedMessage.Cmp(message) != 0 {
		t.Errorf("Expected message %d, but got decryptedMessage %d", message, decryptedMessage)
	}

}
