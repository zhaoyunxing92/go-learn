//go:generate wire
//+build !wireinject
package service

type AccountService struct {
	orderService IOrderService
}

type IAccountService interface {
	GetOrderNum() (int, error)
}

// NewAccountService 账户服务
func NewAccountService(order IOrderService) IAccountService {
	return AccountService{order}
}

// GetOrderNum 获取全部订单个数
func (ac AccountService) GetOrderNum() (int, error) {
	return ac.orderService.GetCountOrderNum()
}
