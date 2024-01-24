package helper

import (
	"math/big"
	"testing"
)

// Test for the AddTwo function
func TestLcm(t *testing.T) {

	// Test case 1: Adding 2 to 3 should result in 5
	resultOne := Lcm(big.NewInt(8), big.NewInt(12))
	expectedOne := big.NewInt(24)

	if resultOne.Cmp(expectedOne) != 0 {
		t.Errorf("Expected %d, but got %d", expectedOne, resultOne)
	}

	// Test case 2: Adding 2 to -1 should result in 1
	resultTwo := Lcm(big.NewInt(15), big.NewInt(20))
	expectedTwo := big.NewInt(60)
	if resultTwo.Cmp(expectedTwo) != 0 {
		t.Errorf("Expected %d, but got %d", expectedTwo, resultTwo)
	}

	// Add more test cases as needed
}
