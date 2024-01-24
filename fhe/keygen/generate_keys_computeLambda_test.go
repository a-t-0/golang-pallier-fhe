package keygen

import (
	"math/big"
	"testing"
)

// Test for the AddTwo function
func TestComputeLambda(t *testing.T) {
	t.SkipNow()
	// Test case 1: Adding 2 to 3 should result in 5
	var left *big.Int = big.NewInt(13)
	var right *big.Int = big.NewInt(17)
	var result *big.Int = ComputeLambda(left, right)
	var expected *big.Int = big.NewInt(48)
	if result.Cmp(expected) != 0 {
		t.Errorf("Expected %d, but got %d", expected, result)
	}
}
