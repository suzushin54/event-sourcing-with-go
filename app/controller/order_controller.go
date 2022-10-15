package controller

import (
	"suzushin54/event-sourcing-with-go/app/usecase"
)

type OrderController struct {
	orderUseCaseV1 usecase.OrderUseCaseV1
	orderUseCaseV2 usecase.OrderUseCaseV2
}

func NewOrderController(
	orderUseCaseV1 usecase.OrderUseCaseV1,
	orderUseCaseV2 usecase.OrderUseCaseV2,
) *OrderController {
	return &OrderController{
		orderUseCaseV1: orderUseCaseV1,
		orderUseCaseV2: orderUseCaseV2,
	}
}
