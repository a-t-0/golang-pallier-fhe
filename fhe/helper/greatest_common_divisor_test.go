package helper

import "testing"

// Test for the AddTwo function
func TestGcd(t *testing.T) {
	// Test case 1: Adding 2 to 3 should result in 5
	result_one := Gcd(8, 12)
	expected_one := 4
	if result_one != expected_one {
		t.Errorf("Expected %d, but got %d", expected_one, result_one)
	}

	// Test case 1: Adding 2 to 3 should result in 5
	result_two := Gcd(12, 8)
	expected_two := 4
	if result_two != expected_two {
		t.Errorf("Expected %d, but got %d", expected_two, result_two)
	}

	// Test case 2: Adding 2 to -1 should result in 1
	result_three := Gcd(15, 20)
	expected_three := 5
	if result_three != expected_three {
		t.Errorf("Expected %d, but got %d", expected_three, result_three)
	}

	// Test case 2: Adding 2 to -1 should result in 1
	result_four := Gcd(20, 15)
	expected_four := 5
	if result_four != expected_four {
		t.Errorf("Expected %d, but got %d", expected_four, result_four)
	}
}
