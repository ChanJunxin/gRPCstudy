package main

import (
	"context"
	"fmt"
	"grpcstudy/demo17/client/helper"
	. "grpcstudy/demo17/client/service"
	"io"
	"log"
	"google.golang.org/grpc"
)

func main() {

	conn, err := grpc.Dial(":8082", grpc.WithTransportCredentials(helper.GetClientCreds()))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	userClient := NewUserServiceClient(conn)
	ctx := context.Background()

	var i int32
	req := UserScoreRequest{}
	req.Users = make([]*UserInfo, 0)
	for i = 1; i < 6; i++ {
		req.Users = append(req.Users, &UserInfo{UserId: i})
	}
	stream, err := userClient.GetUserScoreByServerStream(ctx, &req)
	if err != nil {
		log.Fatal(err)
	}

	//需要循环读取，服务器停止才退出
	for {
		res, err := stream.Recv()
		//服务端发送结束
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(res.Users)
	}
}
