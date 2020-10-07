# 2.创建gRPC服务端并并运行

```bash
cd server/pbfiles  
protoc --go_out=plugins=grpc:../services Prod.proto
```

