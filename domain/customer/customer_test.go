package customer

import (
	"errors"
	"testing"
)

func TestNewCustomer(t *testing.T) {
	type TestCase struct {
		test      string
		name      string
		expectErr error
	}

	//
	testCases := []TestCase{
		{
			test:      "Empty name validation",
			name:      "",
			expectErr: ErrInvalidPerson,
		},
		{
			test:      "Valid name",
			name:      "Adam",
			expectErr: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			_, err := NewCustomer(tc.name)

			if !errors.Is(err, tc.expectErr) {
				t.Errorf("expected error: %v, got %v", tc.expectErr, err)
			}
		})
	}
}
