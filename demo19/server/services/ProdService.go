package services

import (
	"context"
)

type ProdService struct {
}

//服务具体实现
func (this *ProdService) GetProdStock(ctx context.Context, req *ProdRequest) (*ProdResponse, error) {
	var stock int32 = 0
	switch req.ProdArea {
	case ProdAreas_B:
		stock = 31 //B区库存31个
	case ProdAreas_C:
		stock = 50 //B区库存31个
	default:
		stock = 102 //B区库存31个
	}
	return &ProdResponse{ProdStock: stock}, nil
}

func (this *ProdService) GetProdStocks(ctx context.Context, size *QuerySize) (*ProdResponseList, error) {
	myProdres := []*ProdResponse{
		&ProdResponse{ProdStock: 12},
		&ProdResponse{ProdStock: 13},
		&ProdResponse{ProdStock: 14},
	}
	return &ProdResponseList{Prodres: myProdres}, nil
}

func (this *ProdService) GetProdInfo(ctx context.Context, in *ProdRequest) (*ProdModel, error) {
	return &ProdModel{ProdId: 101, ProdName: "测试商品", ProdPrice: 20.5}, nil
}
