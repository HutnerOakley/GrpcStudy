package main

import (
	"GrpcStudy/stream_proto"
	"bufio"
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"

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
	fileCilent := stream_proto.NewServiceFileStreamClient(conn)
	// 调用服务端的下载文件方法
	stream, err := fileCilent.DownloadFile(context.Background(), &stream_proto.Request{
		Name: "下载文件",
	})
	if err != nil {
		log.Fatalf("发起下载请求失败: %v", err)
	}

	// 接收文件并且保存到本地
	wd, _ := os.Getwd()
	filePath := filepath.Join(wd, "static", "1.test.pdf")
	file, err := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0666)
	if err != nil {
		log.Fatalln("创建本地文件失败:", err)
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	defer writer.Flush() // 使用 defer 确保即使中途报错，缓冲区的数据也能刷入磁盘

	var index int
	for {
		index++
		res, err := stream.Recv()
		if err == io.EOF {
			fmt.Println("文件接收完毕！")
			break
		}
		if err != nil {
			log.Printf("接收流数据出错: %v\n", err)
			break
		}

		if _, err := writer.Write(res.Content); err != nil {
			log.Printf("写入本地文件出错: %v\n", err)
			break
		}
		fmt.Printf("第%d 次， 写入 %d 数据\n", index, len(res.Content))
	}

}
