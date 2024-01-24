package helper

import "testing"

// Test for the AddTwo function
func TestAddTwo(t *testing.T) {
	// Test case 1: Adding 2 to 3 should result in 5
	result_one := Lcm(8, 12)
	expected_one := 24
	if result_one != expected_one {
		t.Errorf("Expected %d, but got %d", expected_one, result_one)
	}

	// Test case 2: Adding 2 to -1 should result in 1
	result_two := Lcm(15, 20)
	expected_two := 60
	if result_two != expected_two {
		t.Errorf("Expected %d, but got %d", expected_two, result_two)
	}

	// Add more test cases as needed
}
