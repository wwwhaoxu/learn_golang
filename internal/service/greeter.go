package service

import (
	"context"
	"helloworld/api/helloworld/v1"
	"helloworld/internal/biz"
)





func (s *GreeterService) SayHello(ctx context.Context, req *v1.HelloRequest) (*v1.HelloReply, error) {
	x, err := s.oc.Create(ctx, &biz.Greeter{})
	return &v1.HelloReply{
		Message: x.Hello,
	}, err
}



