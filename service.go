package main

import (
	"context"
	"fmt"

	"github.com/bestateless/example-service/proto"
)

type Service struct{}

func (s *Service) SayHello(ctx context.Context, req *proto.SayHelloRequest) (*proto.SayHelloResponse, error) {
	return &proto.SayHelloResponse{Message: fmt.Sprintf("Oh, hello there, %s", req.Name)}, nil
}
