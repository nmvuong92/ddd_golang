package memory

import (
	"ddd_golang/aggregates"
	"ddd_golang/domain/customer"
	"errors"
	"github.com/google/uuid"
	"testing"
)

func TestMemory_GetCustomer(t *testing.T) {
	type testCase struct {
		name        string
		id          uuid.UUID
		expectedErr error
	}

	cus, err := aggregates.NewCustomer("percy")
	if err != nil {
		t.Fatal(err)
	}

	id := cus.GetID()

	repo := MemoryRepository{
		customers: map[uuid.UUID]aggregates.Customer{
			id: cus,
		},
	}

	testCases := []testCase{
		{
			name:        "no customer by id",
			id:          uuid.MustParse("2673d384-bbd1-4086-b1b9-742a921137b3"),
			expectedErr: customer.ErrCustomerNotFound,
		},
		{
			name:        "customer by id",
			id:          id,
			expectedErr: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			_, err := repo.Get(tc.id)
			if !errors.Is(err, tc.expectedErr) {
				t.Errorf("unexpected err %v, got %v", tc.expectedErr, err)
			}
		})
	}

}
