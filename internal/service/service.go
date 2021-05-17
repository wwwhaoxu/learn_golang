package service

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	v1 "helloworld/api/helloworld/v1"
	"helloworld/internal/biz"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(NewGreeterService)

type GreeterService struct {
	v1.UnimplementedGreeterServer
	oc *biz.GreeterUsecase
	log *log.Helper
}

func NewGreeterService(oc *biz.GreeterUsecase, logger log.Logger) *GreeterService  {
	return &GreeterService{
		oc: oc,
		log: log.NewHelper("service/greeter", logger),
	}
}