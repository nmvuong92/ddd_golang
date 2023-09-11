package customer

import (
	"ddd_golang/aggregates"
	"errors"
	"github.com/google/uuid"
)

var (
	ErrCustomerNotFound    = errors.New("the customer was not found")
	ErrFailedToAddCustomer = errors.New("failed to add the customer")
	ErrUpdateCustomer      = errors.New("failed to update the customer")
)

type CustomerRepository interface {
	Get(uuid uuid.UUID) (aggregates.Customer, error) // can impl it via mongodb, inmemory, mysql...
	Add(customer aggregates.Customer) error
	Update(customer aggregates.Customer) error
}
