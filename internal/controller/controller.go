package controller

import order_creation "2GIS/internal/usecases/order-creation"

type orderCreation interface {
	Handle(input order_creation.Input) (*order_creation.Output, error)
}

type Controller struct {
	orderCreation orderCreation
}

func NewController(orderCreation orderCreation) *Controller {
	return &Controller{
		orderCreation: orderCreation,
	}
}
