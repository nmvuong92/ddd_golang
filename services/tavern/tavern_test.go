package tavern

import (
	"context"
	"github.com/google/uuid"
	"tavern/domain/product"
	order2 "tavern/services/order"
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

func Test_Tavern(t *testing.T) {

	products := Init_products(t)

	os, err := order2.NewOrderService(
		//WithMemoryCustomerRepository(),
		//	WithMongoCustomerRepository(), //<--- can impl and replace it in the future
		order2.WithMongoCustomerRepository(context.Background(), "mongodb://localhost:27017"),
		order2.WithMemoryProductRepository(products),
	)
	if err != nil {
		t.Fatal(err)
	}

	tavern, err := NewTavern(WithOrderService(os))
	if err != nil {
		t.Fatal(err)
	}

	//cust, err := customer.NewCustomer("percy")
	//if err != nil {
	//	t.Fatal(err)
	//}
	//
	//if err := os.customers.Add(cust); err != nil {
	//	t.Fatal(err)
	//}

	uid, err := os.AddCustomer("percy")
	if err != nil {
		t.Fatal(err)
	}

	order := []uuid.UUID{
		products[0].GetID(),
	}

	err = tavern.Order(uid, order)

	if err != nil {
		t.Fatal(err)
	}
}
