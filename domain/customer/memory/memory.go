// Package memtory is a in-memory implemenation of Customer repository
package memory

import (
	"ddd_golang/aggregates"
	"ddd_golang/domain/customer"
	"fmt"
	"github.com/google/uuid"
	"sync"
)

type MemoryRepository struct {
	customers map[uuid.UUID]aggregates.Customer
	sync.Mutex
}

func New() *MemoryRepository {
	return &MemoryRepository{
		customers: make(map[uuid.UUID]aggregates.Customer),
	}
}

func (mr *MemoryRepository) Get(id uuid.UUID) (aggregates.Customer, error) {
	if customer, ok := mr.customers[id]; ok {
		return customer, nil
	}
	return aggregates.Customer{}, customer.ErrCustomerNotFound
}

func (mr *MemoryRepository) Add(c aggregates.Customer) error {
	if mr.customers == nil {
		mr.Lock()
		mr.customers = make(map[uuid.UUID]aggregates.Customer)
		mr.Unlock()
	}

	// make sure customer is already in repo
	if _, ok := mr.customers[c.GetID()]; ok {
		return fmt.Errorf("customer already exists :%w", customer.ErrFailedToAddCustomer)
	}
	mr.Lock()
	mr.customers[c.GetID()] = c
	mr.Unlock()
	return nil
}

func (mr *MemoryRepository) Update(c aggregates.Customer) error {
	if _, ok := mr.customers[c.GetID()]; !ok {
		return fmt.Errorf("customer does not exist: %w", customer.ErrCustomerNotFound)
	}
	mr.Lock()
	mr.customers[c.GetID()] = c
	mr.Unlock()
	return nil
}
