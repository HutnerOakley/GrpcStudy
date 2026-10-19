package main

import (
	"GrpcStudy/stream_proto"
	"context"
	"fmt"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {

	// 使用grpc.NewClient创建一个到指定地址的gRPC连接。
	conn, err := grpc.NewClient("127.0.0.1:8080", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to create gRPC client: %v", err)
	}
	defer conn.Close()

	// 初始化客户端
	streamClient := stream_proto.NewServiceStreamClient(conn)
	stream, err := streamClient.Fun(context.Background(), &stream_proto.Request{
		Name: "张三111",
	})
	if err != nil {
		log.Fatalf("failed to create gRPC stream: %v", err)
	}

	for i := 0; i < 10; i++ {
		res, err := stream.Recv()
		fmt.Println(res, err)

	}
}
