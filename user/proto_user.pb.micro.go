// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto_user.proto

package user

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

// Api Endpoints for UserService service

func NewUserServiceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for UserService service

type UserService interface {
	//用户登录校验
	Auth(ctx context.Context, in *AuthReq, opts ...client.CallOption) (*UserRep, error)
	//用户手机号登录
	LoginWithPhone(ctx context.Context, in *LoginPhoneReq, opts ...client.CallOption) (*LoginedUserRep, error)
	//用户微信登录
	LoginWithWx(ctx context.Context, in *LoginWxReq, opts ...client.CallOption) (*LoginedUserRep, error)
	//用户手机号绑定
	BindPhone(ctx context.Context, in *BindPhoneReq, opts ...client.CallOption) (*UserRep, error)
	//用户微信绑定
	BindWx(ctx context.Context, in *BindWxReq, opts ...client.CallOption) (*UserRep, error)
	//用户信息查询
	UserGet(ctx context.Context, in *UserGetReq, opts ...client.CallOption) (*UserRep, error)
	//根据用户id查询用户
	UserListByIds(ctx context.Context, in *UserListByIdsReq, opts ...client.CallOption) (*UserListByIdsRep, error)
	//所有用户分页列表查询
	UserPageList(ctx context.Context, in *UserPageListReq, opts ...client.CallOption) (*UserPageListRep, error)
}

type userService struct {
	c    client.Client
	name string
}

func NewUserService(name string, c client.Client) UserService {
	return &userService{
		c:    c,
		name: name,
	}
}

func (c *userService) Auth(ctx context.Context, in *AuthReq, opts ...client.CallOption) (*UserRep, error) {
	req := c.c.NewRequest(c.name, "UserService.Auth", in)
	out := new(UserRep)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userService) LoginWithPhone(ctx context.Context, in *LoginPhoneReq, opts ...client.CallOption) (*LoginedUserRep, error) {
	req := c.c.NewRequest(c.name, "UserService.LoginWithPhone", in)
	out := new(LoginedUserRep)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userService) LoginWithWx(ctx context.Context, in *LoginWxReq, opts ...client.CallOption) (*LoginedUserRep, error) {
	req := c.c.NewRequest(c.name, "UserService.LoginWithWx", in)
	out := new(LoginedUserRep)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userService) BindPhone(ctx context.Context, in *BindPhoneReq, opts ...client.CallOption) (*UserRep, error) {
	req := c.c.NewRequest(c.name, "UserService.BindPhone", in)
	out := new(UserRep)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userService) BindWx(ctx context.Context, in *BindWxReq, opts ...client.CallOption) (*UserRep, error) {
	req := c.c.NewRequest(c.name, "UserService.BindWx", in)
	out := new(UserRep)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userService) UserGet(ctx context.Context, in *UserGetReq, opts ...client.CallOption) (*UserRep, error) {
	req := c.c.NewRequest(c.name, "UserService.UserGet", in)
	out := new(UserRep)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userService) UserListByIds(ctx context.Context, in *UserListByIdsReq, opts ...client.CallOption) (*UserListByIdsRep, error) {
	req := c.c.NewRequest(c.name, "UserService.UserListByIds", in)
	out := new(UserListByIdsRep)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userService) UserPageList(ctx context.Context, in *UserPageListReq, opts ...client.CallOption) (*UserPageListRep, error) {
	req := c.c.NewRequest(c.name, "UserService.UserPageList", in)
	out := new(UserPageListRep)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for UserService service

type UserServiceHandler interface {
	//用户登录校验
	Auth(context.Context, *AuthReq, *UserRep) error
	//用户手机号登录
	LoginWithPhone(context.Context, *LoginPhoneReq, *LoginedUserRep) error
	//用户微信登录
	LoginWithWx(context.Context, *LoginWxReq, *LoginedUserRep) error
	//用户手机号绑定
	BindPhone(context.Context, *BindPhoneReq, *UserRep) error
	//用户微信绑定
	BindWx(context.Context, *BindWxReq, *UserRep) error
	//用户信息查询
	UserGet(context.Context, *UserGetReq, *UserRep) error
	//根据用户id查询用户
	UserListByIds(context.Context, *UserListByIdsReq, *UserListByIdsRep) error
	//所有用户分页列表查询
	UserPageList(context.Context, *UserPageListReq, *UserPageListRep) error
}

func RegisterUserServiceHandler(s server.Server, hdlr UserServiceHandler, opts ...server.HandlerOption) error {
	type userService interface {
		Auth(ctx context.Context, in *AuthReq, out *UserRep) error
		LoginWithPhone(ctx context.Context, in *LoginPhoneReq, out *LoginedUserRep) error
		LoginWithWx(ctx context.Context, in *LoginWxReq, out *LoginedUserRep) error
		BindPhone(ctx context.Context, in *BindPhoneReq, out *UserRep) error
		BindWx(ctx context.Context, in *BindWxReq, out *UserRep) error
		UserGet(ctx context.Context, in *UserGetReq, out *UserRep) error
		UserListByIds(ctx context.Context, in *UserListByIdsReq, out *UserListByIdsRep) error
		UserPageList(ctx context.Context, in *UserPageListReq, out *UserPageListRep) error
	}
	type UserService struct {
		userService
	}
	h := &userServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&UserService{h}, opts...))
}

type userServiceHandler struct {
	UserServiceHandler
}

func (h *userServiceHandler) Auth(ctx context.Context, in *AuthReq, out *UserRep) error {
	return h.UserServiceHandler.Auth(ctx, in, out)
}

func (h *userServiceHandler) LoginWithPhone(ctx context.Context, in *LoginPhoneReq, out *LoginedUserRep) error {
	return h.UserServiceHandler.LoginWithPhone(ctx, in, out)
}

func (h *userServiceHandler) LoginWithWx(ctx context.Context, in *LoginWxReq, out *LoginedUserRep) error {
	return h.UserServiceHandler.LoginWithWx(ctx, in, out)
}

func (h *userServiceHandler) BindPhone(ctx context.Context, in *BindPhoneReq, out *UserRep) error {
	return h.UserServiceHandler.BindPhone(ctx, in, out)
}

func (h *userServiceHandler) BindWx(ctx context.Context, in *BindWxReq, out *UserRep) error {
	return h.UserServiceHandler.BindWx(ctx, in, out)
}

func (h *userServiceHandler) UserGet(ctx context.Context, in *UserGetReq, out *UserRep) error {
	return h.UserServiceHandler.UserGet(ctx, in, out)
}

func (h *userServiceHandler) UserListByIds(ctx context.Context, in *UserListByIdsReq, out *UserListByIdsRep) error {
	return h.UserServiceHandler.UserListByIds(ctx, in, out)
}

func (h *userServiceHandler) UserPageList(ctx context.Context, in *UserPageListReq, out *UserPageListRep) error {
	return h.UserServiceHandler.UserPageList(ctx, in, out)
}