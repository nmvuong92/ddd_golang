package services

import (
	"github.com/google/uuid"
	"log"
)

// note: tavern like a bar, shop

// function signature
type TavernConfiguration func(os *Tavern) error

type Tavern struct {
	//orderservice to takes orders (hold subservices)
	OrderService *OrderService

	// billing service
	BillingService interface{} // <--- can implement then
}

func NewTavern(cfgs ...TavernConfiguration) (*Tavern, error) {
	t := &Tavern{}

	for _, cfg := range cfgs {
		if err := cfg(t); err != nil {
			return nil, err
		}
	}
	return t, nil
}

func WithOrderService(os *OrderService) TavernConfiguration {
	return func(t *Tavern) error {
		t.OrderService = os
		return nil
	}
}

func (t *Tavern) Order(customer uuid.UUID, products []uuid.UUID) error {
	price, err := t.OrderService.CreateOrder(customer, products)
	if err != nil {
		return err
	}

	// implement mongodb repository customer
	log.Printf("\nBill the customer: %0.0f", price)
	// billing service
	return nil
}
