package v1

import (
	"context"
	"fmt"
	"github.com/sirupsen/logrus"
)

func init() {
	logrus.WithFields(logrus.Fields{
		"name": "dj",
		"age": 18,
	}).Info("info msg")
	fmt.Println("init hello package")
}

// HelloService hello service
//@DubboService(interfaceName = "helloService",interface="com.cloud.dubbo.service.HelloService", version = "2.0.0", group = "dubbo")
type HelloService struct {
	Say func(ctx context.Context, req []interface{}) error `dubbo:"hello"`
}

// OrderService order service
//@DubboService(interfaceName = "orderService",interface="com.cloud.dubbo.service.OrderService", version = "1.0.0", group = "dubbo")
type OrderService struct {
	GetOrder func(ctx context.Context, req []interface{}) error
}

