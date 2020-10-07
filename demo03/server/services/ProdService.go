package services

//服务的具体实现，客户端不需要这个文件，需要Prod.pb.go文件
import (
	"context"
	//"google.golang.org/grpc"
)

type ProdService struct {
}

//GetProdStock 服务具体实现
func (this *ProdService) GetProdStock(ctx context.Context, req *ProdRequest) (*ProdResponse, error) {
	return &ProdResponse{ProdStock: 20}, nil
}
