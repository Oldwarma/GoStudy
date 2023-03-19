package main

import (
	"context"
	"gostudy/study3/7-Grpc/test/rpc"
)

type Server struct {
	rpc.UnimplementedTestServer
}

func (s *Server) Hello(ctx context.Context, empty *rpc.Empty) (*rpc.HelloResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s *Server) Register(ctx context.Context, request *rpc.RegisterRequest) (*rpc.RegisterResponse, error) {
	//TODO implement me
	panic("implement me")
}
