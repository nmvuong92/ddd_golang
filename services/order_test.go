package services

import (
	"ddd_golang/aggregates"
	"github.com/google/uuid"
	"testing"
)

// initialize test easier (create simple array product for us)
func init_products(t *testing.T) []aggregates.Product {
	beer, err := aggregates.NewProduct("beer", "healthy rice", 10)
	if err != nil {
		t.Fatal(err)
	}
	peenuts, err := aggregates.NewProduct("peanuts", "snacks", 0.99)
	if err != nil {
		t.Fatal(err)
	}

	wine, err := aggregates.NewProduct("wine", "nasty drink", 0.99)
	if err != nil {
		t.Fatal(err)
	}
	return []aggregates.Product{
		beer, peenuts, wine,
	}
}

func TestOrder_NewOrderService(t *testing.T) {
	products := init_products(t)

	// how cool configuration pattern,
	// make realy easy in the future
	os, err := NewOrderService(
		WithMemoryCustomerRepository(),
		WithMemoryProductRepository(products),
	)

	if err != nil {
		t.Fatal(err)
	}

	cust, err := aggregates.NewCustomer("percy")
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
