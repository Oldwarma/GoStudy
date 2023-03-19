package main

import (
	"context"
	"gostudy/study3/7-Grpc/server/rpc"
)

type Server struct {
}

func (s Server) Hello(ctx context.Context, empty *rpc.Empty) (*rpc.HelloResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s Server) Register(ctx context.Context, request *rpc.RegisterRequest) (*rpc.RegisterResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s Server) mustEmbedUnimplementedServerServer() {
	//TODO implement me
	panic("implement me")
}

//// pb文件,ctx上下文
//func (s Server) Hello(req *rpc.Empty, ctx context.Context) (*rpc.HelloResponse, error) {
//	resp := rpc.HelloResponse{Hello: "hello client"}
//	return &resp, nil
//}
//
//func (s Server) Register(req *rpc.Empty, ctx context.Context) (*rpc.RegisterResponse, error) {
//	resp := rpc.RegisterResponse{}
//	resp.Uid = fmt.Sprintf("%s.%s", req.GetName(), req.GetPassword())
//	return &resp, nil
//}
