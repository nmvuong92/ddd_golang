package customer

import (
	"errors"
	"github.com/google/uuid"
)

var (
	ErrCustomerNotFound    = errors.New("the customer was not found")
	ErrFailedToAddCustomer = errors.New("failed to add the customer")
	ErrUpdateCustomer      = errors.New("failed to update the customer")
)

// // to manage and handle the customer aggregate
type Repository interface {
	Get(uuid uuid.UUID) (Customer, error) // can impl it via mongodb, inmemory, mysql...
	Add(customer Customer) error
	Update(customer Customer) error
}
