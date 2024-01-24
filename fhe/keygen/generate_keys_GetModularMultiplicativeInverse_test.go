package keygen

import (
	"fmt"
	"math/big"
	"testing"
)

// Test for the AddTwo function
func TestGetModularMultiplicativeInverse(t *testing.T) {

	// Test case 1: Adding 2 to 3 should result in 5
	var n *big.Int = big.NewInt(221)
	var g *big.Int = big.NewInt(4886)
	var lambda *big.Int = big.NewInt(48)

	result := getModularMultiplicativeInverse(g, lambda, n)
	var expected *big.Int = big.NewInt(159)
	fmt.Printf("result=%d", result)

	if expected.Cmp(result) != 0 {
		t.Errorf("Expected %d, but got %d", expected, result)
	}
}
