# WebPrj1 | 字节青训营web小项目

## 服务的运行和访问

### 1. 启动服务
```shell
go run server.go
```
### 2. 访问并返还页面信息
访问本地服务器并获取页面信息时，由于gin.default()初始化默认8080端口，因此要访问[http://localhost:8080/](http://localhost:8080/)。若不需要解析json数据可在cmd中输入以下命令
```shell
curl --location --request GET http://localhost:8080/community/page/2
```
若要解析json数据得到更直观的数据输出，需要下载jq，并通过通道解析并输出
```shell
curl --location --request GET http://localhost:8080/community/page/2 | sq
```