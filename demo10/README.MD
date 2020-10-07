# 10.语法速学(3) 导入外部proto,获取商品详情


需求：
一般来讲 商品的信息包含
商品ID、商品名称、价格等等，于是就需要把这些专门放到一个proto文件中（就好像建立实体一样）
 


* 启动服务端
```bash
cd server
go run server.go
```
* 启动客户端连接并查询
```bash
cd client
go run main.go
```