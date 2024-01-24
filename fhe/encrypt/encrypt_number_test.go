package encrypt

import (
	"math/big"
	"testing"
)

// Test for the AddTwo function
func TestEncryptNumber(t *testing.T) {
	// Test case 1: Adding 2 to 3 should result in 5
	var message *big.Int = big.NewInt(123)
	var rand_nr *big.Int = big.NewInt(666)
	var n *big.Int = big.NewInt(221)
	var g *big.Int = big.NewInt(4886)

	var resultOne *big.Int = EncryptNumber(message, rand_nr, n, g)
	var expectedOne *big.Int = big.NewInt(25889)
	// var expectedOne *big.Int = big.NewInt(39655)

	if resultOne.Cmp(expectedOne) != 0 {

		t.Errorf("Expected %d, but got %d", expectedOne, resultOne)
	}

}
