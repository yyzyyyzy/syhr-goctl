// Code generated by goctl. DO NOT EDIT!
// Source: hi.proto

package server

import (
	"context"

	greetlogic "github.com/yyzyyyzy/goctls/example/rpc/hi/internal/logic/greet"
	"github.com/yyzyyyzy/goctls/example/rpc/hi/internal/svc"
	"github.com/yyzyyyzy/goctls/example/rpc/hi/pb/hi"
)

type GreetServer struct {
	svcCtx *svc.ServiceContext
	hi.UnimplementedGreetServer
}

func NewGreetServer(svcCtx *svc.ServiceContext) *GreetServer {
	return &GreetServer{
		svcCtx: svcCtx,
	}
}

func (s *GreetServer) SayHi(ctx context.Context, in *hi.HiReq) (*hi.HiResp, error) {
	l := greetlogic.NewSayHiLogic(ctx, s.svcCtx)
	return l.SayHi(in)
}

func (s *GreetServer) SayHello(ctx context.Context, in *hi.HelloReq) (*hi.HelloResp, error) {
	l := greetlogic.NewSayHelloLogic(ctx, s.svcCtx)
	return l.SayHello(in)
}
