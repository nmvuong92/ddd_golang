package main

import (
	"context"
	"github.com/google/uuid"
	"tavern/domain/product"
	"tavern/services/order"
	"tavern/services/tavern"
)

func main() {
	products := productInventory()

	orderService, err := order.NewOrderService(
		//order2.WithMemoryCustomerRepository(),
		order.WithMongoCustomerRepository(context.Background(), "mongodb://localhost:27017"),
		order.WithMemoryProductRepository(products),
	)
	if err != nil {
		panic(err)
	}

	tarvenObject, err := tavern.NewTavern(
		tavern.WithOrderService(orderService),
	)
	if err != nil {
		panic(err)
	}
	customerUUID, err := orderService.AddCustomer("percy")
	if err != nil {
		panic(err)
	}
	orderItem := []uuid.UUID{
		products[0].GetID(),
	}
	err = tarvenObject.Order(customerUUID, orderItem)
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
