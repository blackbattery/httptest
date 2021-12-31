// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// protoc-gen-go-http v2.1.2

package v1

import (
	context "context"
	http "github.com/go-kratos/kratos/v2/transport/http"
	binding "github.com/go-kratos/kratos/v2/transport/http/binding"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
var _ = new(context.Context)
var _ = binding.EncodeURL

const _ = http.SupportPackageIsVersion1

type PlaycardtensflowServiceHTTPServer interface {
	SendPlayCardToServer(context.Context, *ReqPlayCardinfo) (*RespPlayCardInfo, error)
}

func RegisterPlaycardtensflowServiceHTTPServer(s *http.Server, srv PlaycardtensflowServiceHTTPServer) {
	r := s.Route("/")
	r.GET("/playcardtensflowService/{player_position}", _PlaycardtensflowService_SendPlayCardToServer0_HTTP_Handler(srv))
}

func _PlaycardtensflowService_SendPlayCardToServer0_HTTP_Handler(srv PlaycardtensflowServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ReqPlayCardinfo
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/playcardtensflow.v1.playcardtensflowService/SendPlayCardToServer")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.SendPlayCardToServer(ctx, req.(*ReqPlayCardinfo))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*RespPlayCardInfo)
		return ctx.Result(200, reply)
	}
}

type PlaycardtensflowServiceHTTPClient interface {
	SendPlayCardToServer(ctx context.Context, req *ReqPlayCardinfo, opts ...http.CallOption) (rsp *RespPlayCardInfo, err error)
}

type PlaycardtensflowServiceHTTPClientImpl struct {
	cc *http.Client
}

func NewPlaycardtensflowServiceHTTPClient(client *http.Client) PlaycardtensflowServiceHTTPClient {
	return &PlaycardtensflowServiceHTTPClientImpl{client}
}

func (c *PlaycardtensflowServiceHTTPClientImpl) SendPlayCardToServer(ctx context.Context, in *ReqPlayCardinfo, opts ...http.CallOption) (*RespPlayCardInfo, error) {
	var out RespPlayCardInfo
	pattern := "/playcardtensflowService/{player_position}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/playcardtensflow.v1.playcardtensflowService/SendPlayCardToServer"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}
