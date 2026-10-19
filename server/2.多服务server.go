package main

import (
	"GrpcStudy/duo_proto"
	"context"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
)

type OrderServiceServer struct {
	duo_proto.UnimplementedOrderServiceServer
}

func (o *OrderServiceServer) Buy(ctx context.Context, request *duo_proto.Request) (*duo_proto.Response, error) {

	return &duo_proto.Response{
		Name: "order11",
	}, nil

}

type VideoServiceServer struct {
	duo_proto.UnimplementedVideoServiceServer
}

func (v *VideoServiceServer) Look(ctx context.Context, request *duo_proto.Request) (*duo_proto.Response, error) {
	return &duo_proto.Response{
		Name: "video22",
	}, nil
}

func main() {

	// 监听端口
	listen, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}

	// 创建grpc服务
	server := grpc.NewServer()
	// 注册服务
	duo_proto.RegisterOrderServiceServer(server, &OrderServiceServer{})
	duo_proto.RegisterVideoServiceServer(server, &VideoServiceServer{})
	fmt.Println("grpc server 运行在:8080")
	// 处理客户端请求
	err = server.Serve(listen)

}
