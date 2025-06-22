package calculator

import "testing"

// Name convention for testing:
// Correct test: Should...Correct
// Incorrect test: Should...Incorrect

// Triple A standard:
// A - Arrange: what's the test?
// A - Act: what's the result?
// A - Assert: verify

// Testing mill command:
// go test -v in terminal

func TestSumCorrect(t *testing.T) {

	// Arrange
	test := sum(3, 2, 1)

	// Act
	result := 6

	// Assert
	if test != result {
		t.Error("Expected value:", result, "Returned value:", test)
	}
}

func TestSumIncorrect(t *testing.T) {
	test := sum(3, 2, 1)
	result := 7

	if test != result {
		t.Error("Expected value:", result, "Returned value:", test)
	}
}

func TestMultCorrect(t *testing.T) {
	test := multiply(10, 20)
	result := 200

	if test != result {
		t.Error("Expected value:", result, "Returned value:", test)
	}
}

func TestMultIncorrect(t *testing.T) {
	test := multiply(10, 20)
	result := 258

	if test != result {
		t.Error("Expected value:", result, "Returned value:", test)
	}
}

func TestSubtractCorrect(t *testing.T) {
	test := subtract(20, 10)
	result := 10

	if test != result {
		t.Error("Expected value:", result, "Returned value:", test)
	}
}

func TestSubtractIncorrect(t *testing.T) {
	test := subtract(20, 10)
	result := 11

	if test != result {
		t.Error("Expected value:", result, "Returned value:", test)
	}
}

func TestDivideCorrect(t *testing.T) {
	test := divide(20, 10)
	result := 2

	if test != result {
		t.Error("Expected value:", result, "Returned value:", test)
	}
}

func TestDivideIncorrect(t *testing.T) {
	test := divide(20, 10)
	result := 1

	if test != result {
		t.Error("Expected value:", result, "Returned value:", test)
	}
}
