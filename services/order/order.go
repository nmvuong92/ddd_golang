package order

import (
	"context"
	"github.com/google/uuid"
	"log"
	"tavern/domain/customer"
	"tavern/domain/customer/memory"
	"tavern/domain/customer/mongo"
	"tavern/domain/product"
	product_mem "tavern/domain/product/memory"
)

type OrderConfiguration func(os *OrderService) error

type OrderService struct {
	customers customer.Repository
	products  product.Repository

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
func WithCustomerRepository(cr customer.Repository) OrderConfiguration {
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

func WithMemoryProductRepository(products []product.Product) OrderConfiguration {
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
	var products []product.Product
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

func (o *OrderService) AddCustomer(name string) (uuid.UUID, error) {
	c, err := customer.NewCustomer(name)
	if err != nil {
		return uuid.Nil, err
	}

	err = o.customers.Add(c)
	if err != nil {
		return uuid.Nil, err
	}

	return c.GetID(), nil
}
