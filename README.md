# GrpcStudy

一个使用 Go 学习 RPC 和 gRPC 的示例项目。项目包含普通 RPC、gRPC 单次请求、同一服务端注册多个服务，以及服务端流、客户端流和双向流示例。

## 环境要求

- Go 1.25.0（见 `go.mod`）
- 首次运行时需要下载 `go.mod` 中的依赖

`.proto` 文件和已生成的 Go 代码都包含在项目中，直接运行示例不需要安装 `protoc`。

## 快速开始

在项目根目录打开两个终端。先启动最基础的 gRPC 服务端：

```powershell
go run server/main.go
```

再在另一个终端运行对应客户端：

```powershell
go run client/main.go
```

客户端会向 `127.0.0.1:8080` 调用 `HelloService.SayHello` 并打印响应。按 `Ctrl+C` 停止服务端。

## 示例列表

下面每一行的服务端和客户端是一组。仍然使用两个终端，先运行服务端，再运行客户端；切换示例前先停止上一个服务端，因为它们都使用 `8080` 端口。

| 示例 | 服务端命令 | 客户端命令 | 内容 |
| --- | --- | --- | --- |
| 普通 RPC | `go run "server/1.原生rpc.go"` | `go run "client/1.原生rpc.go"` | 使用 Go 标准库 `net/rpc` 完成加法调用 |
| 多服务 | `go run "server/2.多服务server.go"` | `go run "client/2.多服务client.go"` | 一个 gRPC 服务端注册订单和视频两个服务 |
| 服务端流 | `go run "server/3.服务端流式server.go"` | `go run "client/3.服务端流式client.go"` | 一次请求接收多条响应 |
| 文件下载 | `go run "server/4.下载文件(服务端流式)server.go"` | `go run "client/4.下载文件(服务端流式)client.go"` | 通过服务端流传输 PDF |
| 客户端流 | `go run "server/5.客户端流式server.go"` | `go run "client/5.客户端流式client.go"` | 连续发送多条请求后接收一次响应 |
| 文件上传 | `go run "server/6.上传文件(客户端流式)server.go"` | `go run "client/6.上传文件(客户端流式)client.go"` | 通过客户端流传输 PNG |
| 双向流 | `go run "server/7.双向流式server.go"` | `go run "client/7.双向流式client.go"` | 客户端和服务端在同一条流中交替收发消息 |

文件下载读取 `static/不背成语600词【竖版】.pdf`，并写入 `static/1.test.pdf`；文件上传读取 `static/1.png`，并写入 `static/x.png`。这些路径相对于**运行命令时的工作目录**，因此请从项目根目录执行命令。两个输出文件已加入 `.gitignore`。

## 目录说明

| 路径 | 用途 |
| --- | --- |
| `client/`、`server/` | 各示例的客户端和服务端入口 |
| `proto/` | Protocol Buffers 服务与消息定义 |
| `hello_grpc/`、`duo_proto/`、`stream_proto/`、`type_grpc/`、`service_proto/proto/` | 从 `.proto` 生成的 Go 代码 |
| `static/` | 文件传输示例使用的输入文件及运行时输出文件 |

`client/` 和 `server/` 中有多个独立的 `main()`；请按上表指定单个文件运行，不要对整个目录执行 `go run ./client` 或 `go run ./server`。当前示例使用明文的本地连接，适合学习和本机演示。
