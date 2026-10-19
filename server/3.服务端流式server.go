package main

import (
	"GrpcStudy/stream_proto"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
)

type ServiceStreamServer struct {
	stream_proto.UnimplementedServiceStreamServer
}

func (s *ServiceStreamServer) Fun(request *stream_proto.Request, stream stream_proto.ServiceStream_FunServer) error {
	fmt.Println(request)

	for i := 0; i < 10; i++ {
		err := stream.SendMsg(&stream_proto.Response{
			Text: fmt.Sprintf("第%d轮数据", i),
		})
		if err != nil {
			return err
		}
	}

	return nil
}

func main() {

	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}

	server := grpc.NewServer()
	stream_proto.RegisterServiceStreamServer(server, &ServiceStreamServer{})
	server.Serve(listener)
}
