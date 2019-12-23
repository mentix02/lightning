package main

import (
	"testing"
)

// TestInvalidKeyExtraction tests the extractKeyFromToken
// functions with test strings that result in an error return.
func TestInvalidKeyExtraction(t *testing.T) {
	// Test with string without space.
	_, err := extractKeyFromToken("Token1432502")
	if err != nil {
		if err.Error() != "Authentication credentials were not provided." {
			t.Error("Invalid error message returned.")
		}
	} else {
		t.Error("No error message returned.")
	}

	// Test with string with three spaces.
	_, err = extractKeyFromToken("")
	if err != nil {
		if err.Error() != "Authentication credentials were not provided." {
			t.Error("Invalid error message returned.")
		}
	} else {
		t.Error("No error message returned.")
	}

	// Test with first word (should be Token), wrong.
	_, err = extractKeyFromToken("Abcd 123545")
	if err != nil {
		if err.Error() != "Invalid token." {
			t.Error("Invalid error message returned.")
		}
	} else {
		t.Error("No error message returned.")
	}
}

// TestValidKeyExtraction tests similarly to
// TestInvalidKeyExtraction but with proper, valid
// Authorization headers for key extraction.
func TestValidKeyExtraction(t *testing.T) {
	const k string = "3196e58b7d61d159877b68766d8da61ee3f61776"
	key, err := extractKeyFromToken("Token " + k)
	if err != nil {
		t.Error("Illegal error message returned.")
	} else {
		if key != k {
			t.Error("Invalid key returned.")
		}
	}
}
