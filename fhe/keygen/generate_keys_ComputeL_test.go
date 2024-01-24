package keygen

import (
	"math/big"
	"testing"
)

// Test for the AddTwo function
func TestComputeL0(t *testing.T) {
	// Test case 1: Adding 2 to 3 should result in 5
	var top *big.Int = big.NewInt(4)
	var bottom *big.Int = big.NewInt(3)
	var result *big.Int = ComputeL(top, bottom)
	var expected *big.Int = big.NewInt(1)
	if result.Cmp(expected) != 0 {
		t.Errorf("Expected %d, but got %d", expected, result)
	}
}

func TestComputeL1(t *testing.T) {
	// Test case 1: Adding 2 to 3 should result in 5
	var top *big.Int = big.NewInt(7)
	var bottom *big.Int = big.NewInt(3)
	var result *big.Int = ComputeL(top, bottom)
	var expected *big.Int = big.NewInt(2)
	if result.Cmp(expected) != 0 {
		t.Errorf("Expected %d, but got %d", expected, result)
	}
}

func TestComputeL2(t *testing.T) {
	// Test case 1: Adding 2 to 3 should result in 5
	var top *big.Int = big.NewInt(3)
	var bottom *big.Int = big.NewInt(4)
	var result *big.Int = ComputeL(top, bottom)
	var expected *big.Int = big.NewInt(0)
	if result.Cmp(expected) != 0 {
		t.Errorf("Expected %d, but got %d", expected, result)
	}
}

func TestComputeL3(t *testing.T) {
	// Test case 1: Adding 2 to 3 should result in 5
	var top *big.Int = big.NewInt(99)
	var bottom *big.Int = big.NewInt(4)
	var result *big.Int = ComputeL(top, bottom)
	var expected *big.Int = big.NewInt(24)
	if result.Cmp(expected) != 0 {
		t.Errorf("Expected %d, but got %d", expected, result)
	}
}
