package main

import (
	"context"
	"github.com/google/uuid"
	"tavern/domain/product"
	order2 "tavern/services/order"
	"tavern/services/tavern"
)

func main() {
	products := productInventory()

	os, err := order2.NewOrderService(
		//order2.WithMemoryCustomerRepository(),
		order2.WithMongoCustomerRepository(context.Background(), "mongodb://localhost:27017"),
		order2.WithMemoryProductRepository(products),
	)
	if err != nil {
		panic(err)
	}

	tarvenObject, err := tavern.NewTavern(
		tavern.WithOrderService(os),
	)
	if err != nil {
		panic(err)
	}
	uid, err := os.AddCustomer("percy")
	if err != nil {
		panic(err)
	}

	order := []uuid.UUID{
		products[0].GetID(),
	}

	err = tarvenObject.Order(uid, order)
	if err != nil {
		panic(err)
	}
}

func productInventory() []product.Product {
	beer, err := product.NewProduct("a", "b", 1.00)
	if err != nil {
		panic(err)
	}
	beer2, err := product.NewProduct("a2", "b", 1.00)
	if err != nil {
		panic(err)
	}
	beer3, err := product.NewProduct("a2", "b", 1.00)
	if err != nil {
		panic(err)
	}

	return []product.Product{beer, beer2, beer3}
}
