package main

import (
	"grpcstudy/demo02/server/services"
	"net"

	"google.golang.org/grpc"
)

func main() {
	//创建grpc服务
	rpcServer := grpc.NewServer()

	//注册服务
	services.RegisterProdServiceServer(rpcServer, new(services.ProdService))

	//监听端口
	lis, _ := net.Listen("tcp", ":8081")
	rpcServer.Serve(lis)
}
