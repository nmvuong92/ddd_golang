package services

import (
	"context"
	"ddd_golang/aggregates"
	"ddd_golang/domain/customer"
	"ddd_golang/domain/customer/memory"
	"ddd_golang/domain/customer/mongo"
	"ddd_golang/domain/product"
	product_mem "ddd_golang/domain/product/memory"
	"github.com/google/uuid"
	"log"
)

type OrderConfiguration func(os *OrderService) error

type OrderService struct {
	customers customer.CustomerRepository
	products  product.ProductRepository

	// billing billing.Service   <---- infuture subservice inside service
}

func NewOrderService(cfgs ...OrderConfiguration) (*OrderService, error) {
	os := &OrderService{}

	// loop through all the Cfgs and apply them
	for _, cfg := range cfgs {
		err := cfg(os)

		if err != nil {
			return nil, err
		}
	}

	return os, nil
}

// WithCustomerRepository applies a customer repository to the OrderService
func WithCustomerRepository(cr customer.CustomerRepository) OrderConfiguration {
	// return a function that matches the OrderConfiguration alias
	return func(os *OrderService) error {
		os.customers = cr
		return nil
	}
}

func WithMemoryCustomerRepository() OrderConfiguration {
	cr := memory.New()
	return WithCustomerRepository(cr)
}

func WithMongoCustomerRepository(ctx context.Context, conString string) OrderConfiguration {
	return func(os *OrderService) error {
		cr, err := mongo.New(ctx, conString)
		if err != nil {
			return err
		}
		os.customers = cr
		return nil
	}
}

func WithMemoryProductRepository(products []aggregates.Product) OrderConfiguration {
	return func(os *OrderService) error {
		pr := product_mem.New()

		for _, p := range products {
			if err := pr.Add(p); err != nil {
				return err
			}
		}

		os.products = pr
		return nil
	}
}

func (o *OrderService) CreateOrder(CustomerID uuid.UUID, productIDS []uuid.UUID) (float64, error) {
	// Fetch customer
	c, err := o.customers.Get(CustomerID)
	if err != nil {
		return 0.0, err
	}
	// Get each Product
	log.Println(c)
	var products []aggregates.Product
	var total float64

	for _, id := range productIDS {
		p, err := o.products.GetByID(id)
		if err != nil {
			return 0.0, err
		}
		products = append(products, p)
		total += p.GetPrice()
	}
	log.Printf("Customer: %s has ordered %d products", c.GetID(), len(products))
	return total, nil
}
