package main

import (
	"GrpcStudy/duo_proto"
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
	orderClient := duo_proto.NewOrderServiceClient(conn)
	// 调用方法
	res1, err := orderClient.Buy(
		context.Background(), &duo_proto.Request{
			Name: "buy",
		})
	fmt.Println(res1, err)

	videoClient := duo_proto.NewVideoServiceClient(conn)
	res2, err := videoClient.Look(context.Background(), &duo_proto.Request{
		Name: "look",
	})
	fmt.Println(res2, err)
}
