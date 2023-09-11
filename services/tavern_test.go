package services

import (
	"context"
	"ddd_golang/aggregates"
	"github.com/google/uuid"
	"testing"
)

func Test_Tavern(t *testing.T) {
	products := init_products(t)

	os, err := NewOrderService(
		//WithMemoryCustomerRepository(),
		//	WithMongoCustomerRepository(), //<--- can impl and replace it in the future
		WithMongoCustomerRepository(context.Background(), "mongodb://localhost:27017"),
		WithMemoryProductRepository(products),
	)
	if err != nil {
		t.Fatal(err)
	}

	tavern, err := NewTavern(WithOrderService(os))
	if err != nil {
		t.Fatal(err)
	}

	cust, err := aggregates.NewCustomer("percy")
	if err != nil {
		t.Fatal(err)
	}

	if err := os.customers.Add(cust); err != nil {
		t.Fatal(err)
	}

	order := []uuid.UUID{
		products[0].GetID(),
	}

	err = tavern.Order(cust.GetID(), order)

	if err != nil {
		t.Fatal(err)
	}
}
