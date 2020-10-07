package services

import (
	"context"
	"fmt"
	"log"

	"github.com/golang/protobuf/ptypes"
)

type OrdersService struct {
}

func (this *OrdersService) NewOrder(ctx context.Context, orderReq *OrderRequest) (*OrderResponse, error) {
	orderMain := orderReq.OrderMain

	//Validate使用此消息的proto定义中定义的规则检查OrderMain上的字段值。如果违反了任何规则，则返回错误。
	//在这里用于检查商品金额是否大于1
	if err := orderMain.Validate(); err != nil {
		return &OrderResponse{
				Status:  "error",
				Message: err.Error(),
			},
			nil
	}

	if orderMain.OrderTime != nil {
		orderTime, err := ptypes.Timestamp(orderMain.OrderTime)
		if err != nil {
			log.Fatal(err)
		}
		timeGo := orderTime.Format("2006-01-02 15:04:05")
		fmt.Println("goTime:", timeGo)
		fmt.Println("orderTime:", orderTime)
		fmt.Printf("server收到订单信息:订单号:%v, 订单金额:%v,订单时间:%v, 子订单:%v\n", orderMain.OrderNo, orderMain.OrderMoney, timeGo, orderMain.OrderDetails)
	} else {
		fmt.Printf("server收到订单信息:订单号:%v, 订单金额:%v, 子订单:%v\n", orderMain.OrderNo, orderMain.OrderMoney, orderMain.OrderDetails)
	}
	return &OrderResponse{Status: "OK", Message: "success"}, nil
}
