package main

import (
	"context"
	"fmt"
	services "grpcstudy/demo03/client/service"
	"log"

	"google.golang.org/grpc"
)

func main() {
	//创建grpc, grpc.WithInsecure()表示不使用证书
	conn, err := grpc.Dial(":8081", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	//创建客户端
	prodClient := services.NewProdServiceClient(conn)

	//获得商品库存
	prodRes, err := prodClient.GetProdStock(context.Background(),
		&services.ProdRequest{ProdId: 20})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(prodRes.ProdStock)
}
