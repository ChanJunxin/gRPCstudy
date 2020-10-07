# 16.流模式入门（上）

前面的例子，仅仅传输比较小的数据
基本模式是客户端请求————服务端响应

如果是传输较大的数据，会带来：
1、数据包过大导致压力徒增
2、需要等待客户端包全部发送，才能处理以及相应

- 举例：假设需要从库里批量查询用户积分（几万到几十万）

先创建用户模型
```
message UserInfo{
    int32 uer_id = 1;
    int32 user_score = 2;
}
```



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

客户端响应：
```bash
[user_id:1 user_score:101  user_id:2 user_score:102  user_id:3 user_score:103  user_id:4 user_score:104  user_id:5 user_score:105 ]
```
