package service

import "go-learn/service/basic"

var ServiceGroupApp = new(ServiceGroup)

type ServiceGroup struct {
	BasicService basic.BasicService
}
