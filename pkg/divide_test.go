package pkg_test

import (
	"errors"
	"testing"

	"github.com/nathaponb/go-cicd/pkg"
)

func TestDivide(t *testing.T) {
	tests := []struct {
		a, b        int
		expected    int
		expectedErr error
	}{
		{10, 2, 5, nil}, // Normal case
		{9, 3, 3, nil},  // Another normal case
		{1, 0, 0, errors.New("cannot divide by zero")}, // Divide by zero case
	}

	for _, tt := range tests {
		result, err := pkg.Divide(tt.a, tt.b)

		// Check if the result matches the expected value
		if result != tt.expected {
			t.Errorf("divide(%d, %d) = %d; want %d", tt.a, tt.b, result, tt.expected)
		}

		// Check if the error matches the expected error
		if err != nil && err.Error() != tt.expectedErr.Error() {
			t.Errorf("Divide(%d, %d) error = %v; want %v", tt.a, tt.b, err, tt.expectedErr)
		}
	}
}
