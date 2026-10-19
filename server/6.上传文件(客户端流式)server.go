package main

import (
	"GrpcStudy/stream_proto"
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"os"

	"google.golang.org/grpc"
)

type ClientFileStream struct {
	stream_proto.UnimplementedClientFileStreamServer
}

func (c *ClientFileStream) UploadFile(stream stream_proto.ClientFileStream_UploadFileServer) error {

	file, err := os.OpenFile("static/x.png", os.O_CREATE|os.O_WRONLY, 0600)
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	var index int
	for {
		index++
		response, err := stream.Recv()
		if err == io.EOF {
			break
		}
		writer.Write(response.Content)
		fmt.Printf("第%d次", index)
	}

	writer.Flush()
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
	stream_proto.RegisterClientFileStreamServer(server, &ClientFileStream{})
	// 处理客户端请求
	server.Serve(listen)
}
