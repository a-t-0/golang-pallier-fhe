package fhe

import "testing"

// Test for the AddTwo function
func TestGcd(t *testing.T) {
	// Test case 1: Adding 2 to 3 should result in 5
	result := gcd(3)
	expected := 5
	if result != expected {
		t.Errorf("Expected %d, but got %d", expected, result)
	}

	// Test case 2: Adding 2 to -1 should result in 1
	result = gcd(-1)
	expected = 1
	if result != expected {
		t.Errorf("Expected %d, but got %d", expected, result)
	}

	// Add more test cases as needed
}
