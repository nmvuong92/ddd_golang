// package aggregate holds our aggrets that combines many entities in to a full objects
package aggregates

import (
	"errors"

	"ddd_golang/entity"
	"ddd_golang/valueobject"
	"github.com/google/uuid"
)

type Customer struct {
	// person is the root entity of customer
	// which means person.ID is the main identifier for the customer
	person   *entity.Person
	products []*entity.Item

	transactions []*valueobject.Transaction
}

var ErrInvalidPerson = errors.New("a customer has to have a valid name")

// NewCustomer is a factory to create a new customer aggregate
// it will validate that the name is not empty
func NewCustomer(name string) (Customer, error) {
	if name == "" {
		return Customer{}, ErrInvalidPerson
	}

	person := &entity.Person{
		Name: name,
		ID:   uuid.New(),
	}

	return Customer{
		person:       person,
		products:     make([]*entity.Item, 0),
		transactions: make([]*valueobject.Transaction, 0),
	}, nil
}

func (c *Customer) GetID() uuid.UUID {
	return c.person.ID
}

func (c *Customer) SetID(id uuid.UUID) {
	if c.person != nil {
		c.person = &entity.Person{}
	}
	c.person.ID = id
}

func (c *Customer) SetName(name string) {
	if c.person == nil {
		c.person = &entity.Person{}
	}
	c.person.Name = name
}

func (c *Customer) GetName() string {
	return c.person.Name
}
