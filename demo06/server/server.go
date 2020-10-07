package main

import (
	"crypto/tls"
	"crypto/x509"
	"grpcstudy/demo06/server/services"
	"io/ioutil"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func main() {
	//通过读取的方式来生成一个证书池TLS，tls 是 ssl 的升级版
	cert, _ := tls.LoadX509KeyPair("cert/server.pem", "cert/server.key") //加载
	certPool := x509.NewCertPool()                                       //生成证书池
	ca, _ := ioutil.ReadFile("cert/ca.pem")                              //读取CA文件
	certPool.AppendCertsFromPEM(ca)                                      //将CA文件放入证书池

	//创建creds
	creds := credentials.NewTLS(&tls.Config{
		Certificates: []tls.Certificate{cert},        //服务端证书
		ClientAuth:   tls.RequireAndVerifyClientCert, //双向验证，需要验证客户端证书
		ClientCAs:    certPool,                       //指定客户端的证书
	})

	rpcServer := grpc.NewServer(grpc.Creds(creds))
	services.RegisterProdServiceServer(rpcServer, new(services.ProdService))
	lis, _ := net.Listen("tcp", ":8081")
	rpcServer.Serve(lis)
}
