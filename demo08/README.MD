# 8.语法速学(1) 返回商品数组、repeated修饰符
### -- repeated 学习

需求：  
 上节课我们实现的是传入一个商品ID获取一个商品库存。
如果我们要获取 **一堆** 商品的库存列表呢？

* repeated 是一个修饰符，返回字段可以重复任意多次（包括0次）
可以认为就是一个数组（切片）


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