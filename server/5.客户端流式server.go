package main

import (
	"GrpcStudy/stream_proto"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
)

type ClientStream struct {
	stream_proto.UnimplementedClientStreamServer
}

func (c *ClientStream) UploadFile(stream stream_proto.ClientStream_UploadFileServer) error {
	for i := 0; i < 10; i++ {
		response, err := stream.Recv()
		fmt.Println(response, err)
	}

	stream.SendAndClose(&stream_proto.Response{
		Text: "完毕了",
	})

	return nil
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
	stream_proto.RegisterClientStreamServer(server, &ClientStream{})
	// 处理客户端请求
	server.Serve(listen)
}
