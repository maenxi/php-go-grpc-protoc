# php-go-grpc-protoc
- php >= 7.2
- go >= 1.7
- go依赖 go get github.com/golang/protobuf
- php依赖 composer、grpc_php_plugin
- 原始目录结构
```
├── README.md
├── go.mod
├── go.sum
├── idl
│   └── Person.proto
├── main.go
├── php
│   ├── client.php
│   └── composer.json
└── services
```

> 生成php、go对应pb文件
```
protoc --proto_path=./idl --plugin=protoc-gen-grpc=/usr/local/bin/grpc_php_plugin --php_out=./php --grpc_out=./php --go_out=plugins=grpc:./services Person.proto 


生成后的目录结构

├── README.md
├── go.mod
├── go.sum
├── idl
│   └── Person.proto
├── main.go
├── php
│   ├── GPBMetadata
│   │   └── Person.php
│   ├── Services
│   │   ├── PersonInfoReq.php
│   │   ├── PersonInfoResp.php
│   │   ├── PersonListReq.php
│   │   ├── PersonListResp.php
│   │   ├── PersonRespData.php
│   │   └── PersonServiceClient.php
│   ├── client.php
│   └── composer.json
└── services
    └── Person.pb.go

```
> 启动golang服务端
```
go mod tidy // 检测模块依赖
go run main.go


Listen on 8002
```
> 测试php客户端
```
cd ./php && composer install // 安装依赖
php client.php


========GetPersonInfo==========
error = 0 | msg = ok | token = token | name = name=110 | login = 1
========GetPersonList==========
error = 0 | msg = ok 
0: token = token | name = name=wahaha | login = 1
1: token = token2 | name = name=wahaha | login = 2

```