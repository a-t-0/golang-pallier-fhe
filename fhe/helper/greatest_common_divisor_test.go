package helper

import (
	"math/big"
	"testing"
)

// Test for the AddTwo function
func TestGcd(t *testing.T) {
	// t.SkipNow()
	// Test case 1: Adding 2 to 3 should result in 5
	resultOne := Gcd(big.NewInt(8), big.NewInt(12))
	expectedOne := big.NewInt(4)
	if resultOne.Cmp(expectedOne) != 0 {
		t.Errorf("Expected %d, but got %d", expectedOne, resultOne)
	}

	// Test case 1: Adding 2 to 3 should result in 5
	resultTwo := Gcd(big.NewInt(12), big.NewInt(8))
	expectedTwo := big.NewInt(4)
	if resultTwo.Cmp(expectedTwo) != 0 {
		t.Errorf("Expected %d, but got %d", expectedTwo, resultTwo)
	}

	// Test case 2: Adding 2 to -1 should result in 1
	result_three := Gcd(big.NewInt(15), big.NewInt(20))
	expected_three := big.NewInt(5)
	if result_three.Cmp(expected_three) != 0 {
		t.Errorf("Expected %d, but got %d", expected_three, result_three)
	}

	// Test case 2: Adding 2 to -1 should result in 1
	result_four := Gcd(big.NewInt(20), big.NewInt(15))
	expected_four := big.NewInt(5)
	if result_four.Cmp(expected_four) != 0 {
		t.Errorf("Expected %d, but got %d", expected_four, result_four)
	}
}
