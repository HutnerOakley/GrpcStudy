package main

import (
	"GrpcStudy/stream_proto"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
)

type BothStream struct {
	stream_proto.UnimplementedBothStreamServer
}

func (b *BothStream) Chat(stream stream_proto.BothStream_ChatServer) error {
	for i := 0; i < 10; i++ {
		request, _ := stream.Recv()
		fmt.Println(request)
		stream.Send(&stream_proto.Response{
			Text: "你好",
		})
	}

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
	stream_proto.RegisterBothStreamServer(server, &BothStream{})
	server.Serve(listen)

}
