// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto_device.proto

package device

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	api "github.com/micro/micro/v3/service/api"
	client "github.com/micro/micro/v3/service/client"
	server "github.com/micro/micro/v3/service/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ api.Endpoint
var _ context.Context
var _ client.Option
var _ server.Option

// Api Endpoints for DeviceService service

func NewDeviceServiceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for DeviceService service

type DeviceService interface {
	//设备信息上报
	DeviceUpload(ctx context.Context, in *DeviceUploadRep, opts ...client.CallOption) (*DeviceUploadRep, error)
	//根据udi查询设备
	DeviceGetByUdi(ctx context.Context, in *DeviceUdiGetReq, opts ...client.CallOption) (*DeviceRep, error)
	//根据id查询设备
	DeviceGet(ctx context.Context, in *DeviceIdGetReq, opts ...client.CallOption) (*DeviceRep, error)
	//设备列表
	DevicePageList(ctx context.Context, in *DevicePageReq, opts ...client.CallOption) (*DevicePageRep, error)
}

type deviceService struct {
	c    client.Client
	name string
}

func NewDeviceService(name string, c client.Client) DeviceService {
	return &deviceService{
		c:    c,
		name: name,
	}
}

func (c *deviceService) DeviceUpload(ctx context.Context, in *DeviceUploadRep, opts ...client.CallOption) (*DeviceUploadRep, error) {
	req := c.c.NewRequest(c.name, "DeviceService.DeviceUpload", in)
	out := new(DeviceUploadRep)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *deviceService) DeviceGetByUdi(ctx context.Context, in *DeviceUdiGetReq, opts ...client.CallOption) (*DeviceRep, error) {
	req := c.c.NewRequest(c.name, "DeviceService.DeviceGetByUdi", in)
	out := new(DeviceRep)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *deviceService) DeviceGet(ctx context.Context, in *DeviceIdGetReq, opts ...client.CallOption) (*DeviceRep, error) {
	req := c.c.NewRequest(c.name, "DeviceService.DeviceGet", in)
	out := new(DeviceRep)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *deviceService) DevicePageList(ctx context.Context, in *DevicePageReq, opts ...client.CallOption) (*DevicePageRep, error) {
	req := c.c.NewRequest(c.name, "DeviceService.DevicePageList", in)
	out := new(DevicePageRep)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for DeviceService service

type DeviceServiceHandler interface {
	//设备信息上报
	DeviceUpload(context.Context, *DeviceUploadRep, *DeviceUploadRep) error
	//根据udi查询设备
	DeviceGetByUdi(context.Context, *DeviceUdiGetReq, *DeviceRep) error
	//根据id查询设备
	DeviceGet(context.Context, *DeviceIdGetReq, *DeviceRep) error
	//设备列表
	DevicePageList(context.Context, *DevicePageReq, *DevicePageRep) error
}

func RegisterDeviceServiceHandler(s server.Server, hdlr DeviceServiceHandler, opts ...server.HandlerOption) error {
	type deviceService interface {
		DeviceUpload(ctx context.Context, in *DeviceUploadRep, out *DeviceUploadRep) error
		DeviceGetByUdi(ctx context.Context, in *DeviceUdiGetReq, out *DeviceRep) error
		DeviceGet(ctx context.Context, in *DeviceIdGetReq, out *DeviceRep) error
		DevicePageList(ctx context.Context, in *DevicePageReq, out *DevicePageRep) error
	}
	type DeviceService struct {
		deviceService
	}
	h := &deviceServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&DeviceService{h}, opts...))
}

type deviceServiceHandler struct {
	DeviceServiceHandler
}

func (h *deviceServiceHandler) DeviceUpload(ctx context.Context, in *DeviceUploadRep, out *DeviceUploadRep) error {
	return h.DeviceServiceHandler.DeviceUpload(ctx, in, out)
}

func (h *deviceServiceHandler) DeviceGetByUdi(ctx context.Context, in *DeviceUdiGetReq, out *DeviceRep) error {
	return h.DeviceServiceHandler.DeviceGetByUdi(ctx, in, out)
}

func (h *deviceServiceHandler) DeviceGet(ctx context.Context, in *DeviceIdGetReq, out *DeviceRep) error {
	return h.DeviceServiceHandler.DeviceGet(ctx, in, out)
}

func (h *deviceServiceHandler) DevicePageList(ctx context.Context, in *DevicePageReq, out *DevicePageRep) error {
	return h.DeviceServiceHandler.DevicePageList(ctx, in, out)
}
