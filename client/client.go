package main

import (
	"context"
	"log"
	"time"

	pb "github.com/GHjiejie/grpc-test/user" // 导入生成的用户包
	"google.golang.org/grpc"
)

func main() {
	// 建立与 gRPC 服务器的连接
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure(), grpc.WithBlock(), grpc.WithTimeout(5*time.Second))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewUserServiceClient(conn)

	// 创建请求
	req := &pb.GetUserRequest{Id: "123"}
	res, err := client.GetUser(context.Background(), req)
	if err != nil {
		log.Fatalf("could not get user: %v", err)
	}

	log.Printf("User ID: %s, Name: %s", res.Id, res.Name)
}
