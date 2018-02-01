package rpc

import (
  "context"

	"google.golang.org/grpc"
)

func MyInterceptor(
	ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (
    rsp interface{}, err error) {

}
