package usecases

import "github.com/caiocp/go-intensivo/internal/order/entity"

type CreateOrderInputDto struct {
	ID string
	Price float64
	Tax float64
}

type CreateOrderOutputDto struct {
	ID string
	Price float64
	Tax float64
	FinalPrice float64
}

type CalculateFinalPriceUseCase struct {
	OrderRepository entity.OrderRepositoryInterface
}

func NewCalculateFinalPriceUseCase(orderRepository entity.OrderRepositoryInterface) *CalculateFinalPriceUseCase {
	return &CalculateFinalPriceUseCase{OrderRepository: orderRepository}
}

func (c *CalculateFinalPriceUseCase) Execute(input CreateOrderInputDto) (*CreateOrderOutputDto, error) {
	order, err := entity.NewOrder(input.ID, input.Price, input.Tax)
	if err != nil {
		return nil, err
	}

	err = order.CalculateFinalPrice()
	if err != nil {
		return nil, err
	}

	err = c.OrderRepository.Save(order)
	if err != nil {
		return nil, err
	}

	return &CreateOrderOutputDto{
		ID: order.ID,
		Price: order.Price,
		Tax: order.Tax,
		FinalPrice: order.FinalPrice,
	}, nil
}