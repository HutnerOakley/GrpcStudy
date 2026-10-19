package main

import (
	"GrpcStudy/hello_grpc"
	"context"
	"fmt"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
)

// 定义一个结构体,去实现这个服务的全部方法
type HelloServer struct {
	hello_grpc.UnimplementedHelloServiceServer
}

func (h *HelloServer) SayHello(content context.Context, request *hello_grpc.HelloRequest) (res *hello_grpc.HelloResponse, err error) {
	fmt.Println("入参：", request.Name, request.Message)
	pd := &hello_grpc.HelloResponse{
		Name:    "你好",
		Message: "ok",
	}
	return pd, nil
}

func main() {

	// 监听端口
	lister, err := net.Listen("tcp", ":8080")
	if err != nil {
		grpclog.Fatal("Failed to listen: %v", err)
	}

	// 创建一个grpc服务器实例
	s := grpc.NewServer()
	server := HelloServer{}
	// 将server结构体注册为gRPC服务
	hello_grpc.RegisterHelloServiceServer(s, &server)
	// 开始处理客户端请求
	err = s.Serve(lister)

}
