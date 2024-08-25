package matlib

import (
    "testing"
)

// TestAdd tests the Add function.
func TestAdd(t *testing.T) {
    result := Add(2, 3)
    expected := 5
    if result != expected {
        t.Errorf("Add(2, 3) = %d; want %d", result, expected)
    }
}

// TestSubtract tests the Subtract function.
func TestSubtract(t *testing.T) {
    result := Subtract(5, 3)
    expected := 2
    if result != expected {
        t.Errorf("Subtract(5, 3) = %d; want %d", result, expected)
    }
}

// TestMultiply tests the Multiply function.
func TestMultiply(t *testing.T) {
    result := Multiply(2, 3)
    expected := 6
    if result != expected {
        t.Errorf("Multiply(2, 3) = %d; want %d", result, expected)
    }
}

// TestDivide tests the Divide function.
func TestDivide(t *testing.T) {
    // Test case: Normal division
    result, err := Divide(6, 3)
    expected := 2
    if err != nil || result != expected {
        t.Errorf("Divide(6, 3) = %d, %v; want %d, nil", result, err, expected)
    }

    // Test case: Division by zero
    _, err = Divide(6, 0)
    if err == nil {
        t.Error("Expected an error for division by zero, but got nil")
    }
}

