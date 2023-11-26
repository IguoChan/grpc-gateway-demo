package handler

import (
	"context"
	"grpc-gateway-demo/api/demopb"
)

type DemoHandler struct {
	demopb.UnimplementedDemoServer
	client demopb.DemoClient
}

func NewDemoHandler(client demopb.DemoClient) *DemoHandler {
	return &DemoHandler{
		client: client,
	}
}

func (d *DemoHandler) Add(ctx context.Context, data *demopb.DemoData) (*demopb.AddResp, error) {
	//TODO implement me
	panic("implement me")
}

func (d *DemoHandler) Delete(ctx context.Context, req *demopb.DeleteReq) (*demopb.DeleteResp, error) {
	//TODO implement me
	panic("implement me")
}

func (d *DemoHandler) Update(ctx context.Context, data *demopb.DemoData) (*demopb.UpdateResp, error) {
	//TODO implement me
	panic("implement me")
}

func (d *DemoHandler) Get(ctx context.Context, req *demopb.GetReq) (*demopb.GetResp, error) {
	// Do something you want

	return d.client.Get(ctx, req)
}
