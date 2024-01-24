package decrypt

import (
	"math/big"
	"testing"
)

// Test for the AddTwo function
func TestDecryptNumber(t *testing.T) {

	// Test case 1: Adding 2 to 3 should result in 5
	var encryptedMessage *big.Int = big.NewInt(25889)

	var lambda *big.Int = big.NewInt(48)
	var n *big.Int = big.NewInt(221)
	var mu *big.Int = big.NewInt(159)

	var resultOne *big.Int = DecryptNumber(encryptedMessage, lambda, n, mu)
	var expectedOne *big.Int = big.NewInt(123)

	if resultOne.Cmp(expectedOne) != 0 {
		t.Errorf("Expected %d, but got %d", expectedOne, resultOne)
	}

}
