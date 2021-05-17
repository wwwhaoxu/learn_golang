package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gorm"
	"helloworld/internal/biz"
)

type greeterRepo struct {
	data *Data
	log  *log.Helper
}

var _ biz.GreeterRepo = (*greeterRepo)(nil)


type Greeter struct {
	gorm.Model
	Hello string
}

// NewGreeterRepo .
func NewGreeterRepo(data *Data, logger log.Logger) biz.GreeterRepo {
	return &greeterRepo{
		data: data,
		log: log.NewHelper("data/order", logger),
	}
}

func (r *greeterRepo) CreateGreeter(ctx context.Context, g *biz.Greeter) (*biz.Greeter ,error) {
	o := Greeter{Hello: g.Hello}
	result := r.data.db.WithContext(ctx).Create(o)
	return &biz.Greeter{
		Hello: o.Hello,
	}, result.Error

}

