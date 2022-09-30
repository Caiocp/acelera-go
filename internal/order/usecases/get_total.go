package usecases

import "github.com/caiocp/go-intensivo/internal/order/entity"

type GetTotalOutputDto struct {
	Total int
}

type GetTotalUseCase struct {
	OrderRepository entity.OrderRepositoryInterface
}

func NewGetTotalUseCase(orderRepository entity.OrderRepositoryInterface) *GetTotalUseCase {
	return &GetTotalUseCase{OrderRepository: orderRepository}
}

func (u *GetTotalUseCase) Execute() (*GetTotalOutputDto, error) {
	total, err := u.OrderRepository.GetTotal()
	if err != nil {
		return nil, err
	}

	return &GetTotalOutputDto{Total: total}, nil
}