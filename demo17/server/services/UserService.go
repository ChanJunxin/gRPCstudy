package services

import (
	"context"
	"time"
)

type UserService struct {
}

//普通方法
func (*UserService) GetUserScore(ctx context.Context, in *UserScoreRequest) (*UserScoreResponse, error) {
	var score int32 = 101
	users := make([]*UserInfo, 0)
	for _, user := range in.Users {
		user.UserScore = score
		score++
		users = append(users, user)
	}
	return &UserScoreResponse{Users: users}, nil
}

//服务端流式响应，UserService_GetUserScoreByServerStreamServer类型可以发送内容
func (*UserService) GetUserScoreByServerStream(in *UserScoreRequest, stream UserService_GetUserScoreByServerStreamServer) error {
	var score int32 = 101
	users := make([]*UserInfo, 0)
	for index, user := range in.Users {
		user.UserScore = score
		score++
		users = append(users, user)

		//每隔两条发送
		if (index+1)%2 == 0 && index > 0 {
			err := stream.Send(&UserScoreResponse{Users: users})
			if err != nil {
				return err
			}
			//清空切片
			users = (users)[0:0]
		}
		time.Sleep(1 * time.Second)
	}
	if len(users) > 0 {
		return stream.Send(&UserScoreResponse{Users: users})
	}
	return nil
}
