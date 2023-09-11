package product

import (
	"errors"
	"github.com/google/uuid"
	"tavern"
)

// root entity
type Product struct {
	item     *tavern.Item
	price    float64
	quantity int
}

var ErrMissingValue = errors.New("missing important value")

func NewProduct(name, description string, price float64) (Product, error) {
	if name == "" || description == "" {
		return Product{}, ErrMissingValue
	}

	return Product{
		item: &tavern.Item{
			ID:          uuid.New(),
			Name:        name,
			Description: description,
		},
		price:    price,
		quantity: 0,
	}, nil
}

func (p *Product) GetID() uuid.UUID {
	return p.item.ID
}

func (p *Product) GetItem() *tavern.Item {
	return p.item
}

func (p *Product) GetPrice() float64 {
	return p.price
}
