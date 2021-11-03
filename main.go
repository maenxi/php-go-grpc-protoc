package main

import (
	"fmt"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/reflection"
	"net"
	"php-go-grpc-protoc/services"
)

//实现接口
type personService struct{}

func (p personService) GetPersonInfo(ctx context.Context, in *services.PersonInfoReq) (*services.PersonInfoResp, error) {
	resp := new(services.PersonInfoResp)
	resp.Msg = "ok"
	resp.Error = 0

	personData := services.PersonRespData{
		Token: "token",
		Name:  fmt.Sprintf("name=%d", in.GetId()),
		Login: 1,
	}
	resp.Data = &personData
	return resp, nil
}
func (p personService) GetPersonList(ctx context.Context, in *services.PersonListReq) (*services.PersonListResp, error) {
	resp := new(services.PersonListResp)
	resp.Msg = "ok"
	resp.Error = 0
	personListData := []*services.PersonRespData{
		{
			Token: "token",
			Name:  fmt.Sprintf("name=%s", in.GetName()),
			Login: 1,
		},
		{
			Token: "token2",
			Name:  fmt.Sprintf("name=%s", in.GetName()),
			Login: 2,
		},
	}
	resp.Data = personListData
	return resp, nil
}

func main() {
	listen, err := net.Listen("tcp", ":8002")
	if err != nil {
		grpclog.Fatalf("Failed to listen: %v", err)
	}
	fmt.Println("Listen on 8002")
	// 实例化
	s := grpc.NewServer()
	// grpc反射出去 ui界面可访问
	reflection.Register(s)
	// 注册
	services.RegisterPersonServiceServer(s, personService{})
	s.Serve(listen)
}
