package main

import (
	"GrpcStudy/hello_grpc"
	"context"
	"fmt"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	// 使用grpc.NewClient创建一个到指定地址的gRPC连接。
	// 此处使用不安全的证书来实现SSL/TLS连接

	conn, err := grpc.NewClient("127.0.0.1:8080", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to create gRPC client: %v", err)
	}
	defer conn.Close()

	// 初始化客户端
	client := hello_grpc.NewHelloServiceClient(conn)
	res, err := client.SayHello(context.Background(), &hello_grpc.HelloRequest{
		Name:    "hunter",
		Message: "ok",
	})
	fmt.Println(res, err)

}
