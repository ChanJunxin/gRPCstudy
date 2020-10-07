# 17.服务端流模式

场景：分批发送查询结果

上节课是`return &UserScoreResponse{Users: users}`一次性发送。在数据量大并且查询速度慢的情况下，如果一次性查完在返回，时间太慢——>使用服务端流

* 启动服务端
```bash
cd server
go run server.go
```
* 启动客户端
```bash
cd client
go run main.go
```
