package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

type Greeter struct {
	Hello string
}

type GreeterRepo interface {
	CreateGreeter(context.Context, *Greeter) (*Greeter, error)
}

type GreeterUsecase struct {
	repo GreeterRepo
	log  *log.Helper
}

func NewGreeterUsecase(repo GreeterRepo, logger log.Logger) *GreeterUsecase {
	return &GreeterUsecase{repo: repo, log: log.NewHelper("usecase/greeter", logger)}
}

func (uc *GreeterUsecase) Create(ctx context.Context,g *Greeter) (*Greeter, error) {
	return uc.repo.CreateGreeter(ctx,g)
}

