//go:generate wire
//+build !wireinject
package service

type OrderService struct {
}

// IOrderService 订单服务
type IOrderService interface {
	GetCountOrderNum() (int, error)
}

// NewOrderService 创建订单服务
func NewOrderService() IOrderService {
	return new(OrderService)
}

// GetCountOrderNum 获取订单数
func (os OrderService) GetCountOrderNum() (int, error) {
	return 5, nil
}
