package order

import (
	"github.com/google/uuid"
	"tavern/domain/customer"
	"tavern/domain/product"
	"testing"
)

// initialize test easier (create simple array product for us)
func Init_products(t *testing.T) []product.Product {
	beer, err := product.NewProduct("beer", "healthy rice", 10)
	if err != nil {
		t.Fatal(err)
	}
	peenuts, err := product.NewProduct("peanuts", "snacks", 0.99)
	if err != nil {
		t.Fatal(err)
	}

	wine, err := product.NewProduct("wine", "nasty drink", 0.99)
	if err != nil {
		t.Fatal(err)
	}
	return []product.Product{
		beer, peenuts, wine,
	}
}

func TestOrder_NewOrderService(t *testing.T) {
	products := Init_products(t)

	// how cool configuration pattern,
	// make realy easy in the future
	os, err := NewOrderService(
		WithMemoryCustomerRepository(),
		WithMemoryProductRepository(products),
	)

	if err != nil {
		t.Fatal(err)
	}

	cust, err := customer.NewCustomer("percy")
	if err != nil {
		t.Error(err)
	}

	err = os.customers.Add(cust)
	if err != nil {
		t.Error(err)
	}

	order := []uuid.UUID{
		products[0].GetID(),
	}

	_, err = os.CreateOrder(cust.GetID(), order)
	if err != nil {
		t.Error(err)
	}
}
