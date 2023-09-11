package memory

import (
	"ddd_golang/aggregates"
	"ddd_golang/domain/product"
	"github.com/google/uuid"
	"sync"
)

type MemoryProductRepository struct {
	products map[uuid.UUID]aggregates.Product
	sync.Mutex
}

func New() *MemoryProductRepository {
	return &MemoryProductRepository{
		products: make(map[uuid.UUID]aggregates.Product),
	}
}

func (mpr *MemoryProductRepository) GetAll() ([]aggregates.Product, error) {
	var products []aggregates.Product

	for _, product := range mpr.products {
		products = append(products, product)
	}

	return products, nil
}

func (mpr *MemoryProductRepository) GetByID(id uuid.UUID) (aggregates.Product, error) {
	if product, ok := mpr.products[id]; ok {
		return product, nil
	}

	return aggregates.Product{}, product.ErrProductNotFound
}

func (mpr *MemoryProductRepository) Add(newprod aggregates.Product) error {
	mpr.Lock()
	defer mpr.Unlock()

	if _, ok := mpr.products[newprod.GetID()]; ok {
		return product.ErrProductAlreadyExists
	}

	mpr.products[newprod.GetID()] = newprod

	return nil
}

func (mpr *MemoryProductRepository) Update(update aggregates.Product) error {
	mpr.Lock()
	defer mpr.Unlock()
	if _, ok := mpr.products[update.GetID()]; !ok {
		return product.ErrProductNotFound
	}
	mpr.products[update.GetID()] = update
	return nil
}

func (mpr *MemoryProductRepository) Delete(id uuid.UUID) error {
	mpr.Lock()
	defer mpr.Unlock()

	if _, ok := mpr.products[id]; !ok {
		return product.ErrProductNotFound
	}
	delete(mpr.products, id)
	return nil
}
