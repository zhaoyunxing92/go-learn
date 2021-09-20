//go:generate wire
//+build !wireinject
package example

import (
	"github.com/google/wire"
	"github.com/zhaoyunxing/wire/service"
)

func GetUserService() service.IUserService {
	panic(wire.Build(
		service.NewOrderService,
		service.NewAccountService,
		service.NewUserService,
	))
}
