package main

import (
	"context"
	"grpc-gateway-demo/api/demopb"
	"testing"
)

func TestMockJavaServe(t *testing.T) {
	serve(":9090", ":9091", &javaDemo{})
}

type javaDemo struct {
	demopb.UnimplementedDemoServer
}

func (d *javaDemo) Add(ctx context.Context, data *demopb.DemoData) (*demopb.AddResp, error) {
	//TODO implement me
	panic("implement me")
}

func (d *javaDemo) Delete(ctx context.Context, req *demopb.DeleteReq) (*demopb.DeleteResp, error) {
	//TODO implement me
	panic("implement me")
}

func (d *javaDemo) Update(ctx context.Context, data *demopb.DemoData) (*demopb.UpdateResp, error) {
	//TODO implement me
	panic("implement me")
}

func (d *javaDemo) Get(ctx context.Context, req *demopb.GetReq) (*demopb.GetResp, error) {
	return &demopb.GetResp{
		Result: &demopb.Result{
			Code: 0,
			Msg:  "success",
		},
		Data: &demopb.DemoData{
			Id:          req.Id,
			Name:        "java demo",
			Description: "description",
		},
	}, nil
}
