package main

import (
	"context"
	"fmt"
	"grpcstudy/demo09/client/helper"
	. "grpcstudy/demo09/client/service"
	"log"

	"google.golang.org/grpc"
)

func main() {

	conn, err := grpc.Dial(":8082", grpc.WithTransportCredentials(helper.GetClientCreds()))

	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	prodClient := NewProdServiceClient(conn)
	//获取ProductId=12的商品在区域C的库存数													ProdArea不传默认是0
	prodRes, err := prodClient.GetProdStock(context.Background(), &ProdRequest{ProdId: 12, ProdArea: ProdAreas_B})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("库存数量：", prodRes.ProdStock)
}
