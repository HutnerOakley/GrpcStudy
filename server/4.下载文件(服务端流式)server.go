package main

import (
	"GrpcStudy/stream_proto"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"path/filepath"

	"google.golang.org/grpc"
)

type ServiceFileStreamServer struct {
	stream_proto.UnimplementedServiceFileStreamServer
}

func (s *ServiceFileStreamServer) DownloadFile(request *stream_proto.Request, stream stream_proto.ServiceFileStream_DownloadFileServer) error {
	fmt.Println(request)
	// 打开文件
	wd, _ := os.Getwd() // 获取当前工作目录
	filePath := filepath.Join(wd, "static", "不背成语600词【竖版】.pdf")
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	buf := make([]byte, 4096)
	for {
		_, err := file.Read(buf)
		if err == io.EOF {
			break
		}
		if err != nil {
			break
		}
		// 发送
		sendErr := stream.Send(&stream_proto.FileResponse{
			Content: buf,
		})
		if sendErr != nil {
			break
		}

	}
	return nil
}

func main() {
	// 监听端口
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}
	server := grpc.NewServer()
	stream_proto.RegisterServiceFileStreamServer(server, &ServiceFileStreamServer{})
	server.Serve(listener)

}
