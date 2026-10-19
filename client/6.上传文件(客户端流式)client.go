package main

import (
	"GrpcStudy/stream_proto"
	"context"
	"fmt"
	"io"
	"log"
	"os"

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
	client := stream_proto.NewClientFileStreamClient(conn)
	stream, err := client.UploadFile(context.Background())

	file, err := os.Open("static/1.png")
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	for {
		buf := make([]byte, 2048)
		_, err := file.Read(buf)
		if err == io.EOF {
			break
		}
		if err != nil {
			break
		}
		stream.Send(&stream_proto.FileRequest{
			FileName: "x.png",
			Content:  buf,
		})
	}

	res, err := stream.CloseAndRecv()
	fmt.Println(res, err)
}
