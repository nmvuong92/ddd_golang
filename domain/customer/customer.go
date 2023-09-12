// package aggregate holds our aggrets that combines many entities in to a full objects
package customer

import (
	"errors"
	"tavern"

	"github.com/google/uuid"
)

// root entity
type Customer struct {
	// person is the root entity of customer
	// which means person.ID is the main identifier for the customer
	person   *tavern.Person
	products []*tavern.Item

	transactions []*tavern.Transaction
}

var ErrInvalidPerson = errors.New("a customer has to have a valid name")

// NewCustomer is a factory to create a new customer aggregate
// it will validate that the name is not empty
func NewCustomer(name string) (Customer, error) {
	if name == "" {
		return Customer{}, ErrInvalidPerson
	}

	person := &tavern.Person{
		Name: name,
		ID:   uuid.New(),
	}

	return Customer{
		person:       person,
		products:     make([]*tavern.Item, 0),
		transactions: make([]*tavern.Transaction, 0),
	}, nil
}

func (c *Customer) GetID() uuid.UUID {
	return c.person.ID
}

func (c *Customer) SetID(id uuid.UUID) {
	if c.person == nil {
		c.person = &tavern.Person{}
	}
	c.person.ID = id
}

func (c *Customer) SetName(name string) {
	if c.person == nil {
		c.person = &tavern.Person{}
	}
	c.person.Name = name
}

func (c *Customer) GetName() string {
	return c.person.Name
}
