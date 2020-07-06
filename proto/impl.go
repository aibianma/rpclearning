package proto

import (
	"context"
	"github.com/sirupsen/logrus"
)

type HelloServiceImpl struct{}

func (p *HelloServiceImpl) Hello(
	ctx context.Context, args *String,
) (*String, error) {
	logrus.Info(args.GetValue())
	reply := &String{Value: "hello:" + args.GetValue()}
	return reply, nil
}
