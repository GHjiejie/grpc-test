package main

import (
	"context"
	"log"
	"net"

	pb "github.com/GHjiejie/grpc-test/user" // 导入生成的用户包
	"google.golang.org/grpc"
)

// server 结构体实现 UserServiceServer 接口
type server struct {
	pb.UnimplementedUserServiceServer
}

// GetUser 方法的实现
func (s *server) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	// 模拟获取用户信息，这里硬编码一个用户
	return &pb.GetUserResponse{
		Id:   req.Id,
		Name: "John Doe", // 假设返回用户名称
	}, nil
}

func main() {
	// 创建一个 TCP 监听器
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// 创建 gRPC 服务器
	grpcServer := grpc.NewServer()
	pb.RegisterUserServiceServer(grpcServer, &server{}) // 注册服务

	log.Println("Server is running on port :50051")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
