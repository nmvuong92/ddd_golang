package product

import (
	"ddd_golang/aggregates"
	"errors"
	"github.com/google/uuid"
)

var ErrProductNotFound = errors.New("product not found")
var ErrProductAlreadyExists = errors.New("there is already such an product")

// to manage and handle the product aggregate
type ProductRepository interface {
	GetAll() ([]aggregates.Product, error)
	GetByID(id uuid.UUID) (aggregates.Product, error)
	Add(product aggregates.Product) error
	Update(product aggregates.Product) error
	Delete(id uuid.UUID) error
}
