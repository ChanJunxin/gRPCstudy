package main

import (
	"context"
	"grpcstudy/demo07/server/services"
	"log"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"
)

func main() {
	//创建ServeMux
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	gwmux := runtime.NewServeMux()
	defer cancel()

	//指定客户端进行请求时候使用的证书
	opt := []grpc.DialOption{grpc.WithInsecure()}
	//opt := []grpc.DialOption{grpc.WithTransportCredentials(creds)}

	//8099是对应的grpc的监听端口，必须与grpc开放的端口一致,所以启动前务必启动grpc服务
	err := services.RegisterProdServiceHandlerFromEndpoint(ctx, gwmux, "localhost:8099", opt)
	if err != nil {
		log.Fatal(err)
	}

	httpServer := &http.Server{
		Addr:    ":8080",
		Handler: gwmux,
	}
	httpServer.ListenAndServe()
}

/*第二版，B站视频原版
func main() {
	gwmux := runtime.NewServeMux()
	opt := []grpc.DialOption{grpc.WithTransportCredentials(helper.GetClientCreds())}

	err := services.RegisterProdServiceHandlerFromEndpoint(context.Background(), gwmux, "localhost:8099", opt)
	if err != nil {
		log.Fatal(err)
	}

	httpServer := &http.Server{
		Addr:    ":8080",
		Handler: gwmux,
	}
	httpServer.ListenAndServe()
}
*/
