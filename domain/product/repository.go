package product

import (
	"errors"
	"github.com/google/uuid"
)

var ErrProductNotFound = errors.New("product not found")
var ErrProductAlreadyExists = errors.New("there is already such an product")

// to manage and handle the product aggregate
type Repository interface {
	GetAll() ([]Product, error)
	GetByID(id uuid.UUID) (Product, error)
	Add(product Product) error
	Update(product Product) error
	Delete(id uuid.UUID) error
}
