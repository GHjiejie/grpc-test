package main

import (
	"context"
	"log"
	"net"
	"net/http"

	pb "grpc-gateway-example/repository/userpb" // 替换为你的包名

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

// server 结构体实现了 UserServiceServer 接口
type server struct {
	pb.UnimplementedUserServiceServer // 确保实现所有服务接口
}

// 实现 RegisterUser 方法
func (s *server) RegisterUser(ctx context.Context, req *pb.RegisterUserRequest) (*pb.RegisterUserResponse, error) {
	log.Printf("Received RegisterUser request: %v", req)
	return &pb.RegisterUserResponse{Id: "12345", Username: req.Username}, nil
}

// 实现 LoginUser 方法
func (s *server) LoginUser(ctx context.Context, req *pb.LoginUserRequest) (*pb.LoginUserResponse, error) {
	log.Printf("Received LoginUser request: %v", req)
	return &pb.LoginUserResponse{Id: "12345", Username: req.Username}, nil
}

func startGRPCServer() {
	lis, err := net.Listen("tcp", ":8000") // 启动 gRPC 服务器
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterUserServiceServer(s, &server{}) // 注册服务
	log.Printf("gRPC server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func startHTTPServer() {
	mux := runtime.NewServeMux()                   // 创建 gRPC-Gateway 多路复用器
	opts := []grpc.DialOption{grpc.WithInsecure()} // 不安全连接选项

	// 注册 gRPC 服务到 HTTP 网关
	err := pb.RegisterUserServiceHandlerFromEndpoint(context.Background(), mux, "localhost:8000", opts)
	if err != nil {
		log.Fatalf("Failed to register gateway: %v", err)
	}

	log.Println("HTTP server listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", mux)) // 启动 HTTP 服务器
}

func main() {
	go startGRPCServer() // 启动 gRPC 服务器
	startHTTPServer()    // 启动 HTTP 服务器
}
