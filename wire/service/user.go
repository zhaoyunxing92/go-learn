//go:generate wire
//+build !wireinject
package service

type UserService struct {
	accountService IAccountService
}

type IUserService interface {
	GetOrderCount() (int, error)
}

// NewUserService 用户服务
func NewUserService(accountService IAccountService) IUserService {
	return UserService{accountService: accountService}
}

func (us UserService) GetOrderCount() (int, error) {
	return us.accountService.GetOrderNum()
}
