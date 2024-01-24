package fhe

import "testing"

// Test for the AddTwo function
func TestAddTwo(t *testing.T) {
	// Test case 1: Adding 2 to 3 should result in 5
	result := Lcm(8, 12)
	expected := 24
	if result != expected {
		t.Errorf("Expected %d, but got %d", expected, result)
	}

	// Test case 2: Adding 2 to -1 should result in 1
	result = Lcm(15, 20)
	expected = 60
	if result != expected {
		t.Errorf("Expected %d, but got %d", expected, result)
	}

	// Add more test cases as needed
}
