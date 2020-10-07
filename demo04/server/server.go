package main

import (
	"grpcstudy/demo04/server/services"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func main() {
	//创建并加载服务端证书文件与key文件
	creds, err := credentials.NewServerTLSFromFile("keys/server.crt", "keys/server_no_passwd.key")
	if err != nil {
		log.Fatal(err)
	}

	rpcServer := grpc.NewServer(grpc.Creds(creds))
	services.RegisterProdServiceServer(rpcServer, new(services.ProdService))
	lis, _ := net.Listen("tcp", ":8081")
	rpcServer.Serve(lis)
}

//此时服务端有证书验证
