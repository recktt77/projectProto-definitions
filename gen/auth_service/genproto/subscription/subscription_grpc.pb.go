// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v6.30.2
// source: subscription.proto

package subscription

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	SubscriptionService_CreatePlan_FullMethodName            = "/subscription.SubscriptionService/CreatePlan"
	SubscriptionService_GetPlanByID_FullMethodName           = "/subscription.SubscriptionService/GetPlanByID"
	SubscriptionService_UpdatePlan_FullMethodName            = "/subscription.SubscriptionService/UpdatePlan"
	SubscriptionService_GetListOfPlans_FullMethodName        = "/subscription.SubscriptionService/GetListOfPlans"
	SubscriptionService_Subscribe_FullMethodName             = "/subscription.SubscriptionService/Subscribe"
	SubscriptionService_CancelSubscription_FullMethodName    = "/subscription.SubscriptionService/CancelSubscription"
	SubscriptionService_GetSubscription_FullMethodName       = "/subscription.SubscriptionService/GetSubscription"
	SubscriptionService_GetSubscriptionStatus_FullMethodName = "/subscription.SubscriptionService/GetSubscriptionStatus"
)

// SubscriptionServiceClient is the client API for SubscriptionService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SubscriptionServiceClient interface {
	// Plan management
	CreatePlan(ctx context.Context, in *CreatePlanRequest, opts ...grpc.CallOption) (*CreatePlanResponse, error)
	GetPlanByID(ctx context.Context, in *GetPlanByIDRequest, opts ...grpc.CallOption) (*GetPlanByIDResponse, error)
	UpdatePlan(ctx context.Context, in *UpdatePlanRequest, opts ...grpc.CallOption) (*UpdatePlanResponse, error)
	GetListOfPlans(ctx context.Context, in *GetListOfPlansRequest, opts ...grpc.CallOption) (*GetListOfPlansResponse, error)
	// Subscription management
	Subscribe(ctx context.Context, in *SubscribeRequest, opts ...grpc.CallOption) (*SubscribeResponse, error)
	CancelSubscription(ctx context.Context, in *CancelSubscriptionRequest, opts ...grpc.CallOption) (*CancelSubscriptionResponse, error)
	GetSubscription(ctx context.Context, in *GetSubscriptionRequest, opts ...grpc.CallOption) (*GetSubscriptionResponse, error)
	GetSubscriptionStatus(ctx context.Context, in *GetSubscriptionStatusRequest, opts ...grpc.CallOption) (*GetSubscriptionStatusResponse, error)
}

type subscriptionServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSubscriptionServiceClient(cc grpc.ClientConnInterface) SubscriptionServiceClient {
	return &subscriptionServiceClient{cc}
}

func (c *subscriptionServiceClient) CreatePlan(ctx context.Context, in *CreatePlanRequest, opts ...grpc.CallOption) (*CreatePlanResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreatePlanResponse)
	err := c.cc.Invoke(ctx, SubscriptionService_CreatePlan_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *subscriptionServiceClient) GetPlanByID(ctx context.Context, in *GetPlanByIDRequest, opts ...grpc.CallOption) (*GetPlanByIDResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetPlanByIDResponse)
	err := c.cc.Invoke(ctx, SubscriptionService_GetPlanByID_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *subscriptionServiceClient) UpdatePlan(ctx context.Context, in *UpdatePlanRequest, opts ...grpc.CallOption) (*UpdatePlanResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdatePlanResponse)
	err := c.cc.Invoke(ctx, SubscriptionService_UpdatePlan_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *subscriptionServiceClient) GetListOfPlans(ctx context.Context, in *GetListOfPlansRequest, opts ...grpc.CallOption) (*GetListOfPlansResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetListOfPlansResponse)
	err := c.cc.Invoke(ctx, SubscriptionService_GetListOfPlans_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *subscriptionServiceClient) Subscribe(ctx context.Context, in *SubscribeRequest, opts ...grpc.CallOption) (*SubscribeResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SubscribeResponse)
	err := c.cc.Invoke(ctx, SubscriptionService_Subscribe_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *subscriptionServiceClient) CancelSubscription(ctx context.Context, in *CancelSubscriptionRequest, opts ...grpc.CallOption) (*CancelSubscriptionResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CancelSubscriptionResponse)
	err := c.cc.Invoke(ctx, SubscriptionService_CancelSubscription_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *subscriptionServiceClient) GetSubscription(ctx context.Context, in *GetSubscriptionRequest, opts ...grpc.CallOption) (*GetSubscriptionResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetSubscriptionResponse)
	err := c.cc.Invoke(ctx, SubscriptionService_GetSubscription_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *subscriptionServiceClient) GetSubscriptionStatus(ctx context.Context, in *GetSubscriptionStatusRequest, opts ...grpc.CallOption) (*GetSubscriptionStatusResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetSubscriptionStatusResponse)
	err := c.cc.Invoke(ctx, SubscriptionService_GetSubscriptionStatus_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SubscriptionServiceServer is the server API for SubscriptionService service.
// All implementations must embed UnimplementedSubscriptionServiceServer
// for forward compatibility.
type SubscriptionServiceServer interface {
	// Plan management
	CreatePlan(context.Context, *CreatePlanRequest) (*CreatePlanResponse, error)
	GetPlanByID(context.Context, *GetPlanByIDRequest) (*GetPlanByIDResponse, error)
	UpdatePlan(context.Context, *UpdatePlanRequest) (*UpdatePlanResponse, error)
	GetListOfPlans(context.Context, *GetListOfPlansRequest) (*GetListOfPlansResponse, error)
	// Subscription management
	Subscribe(context.Context, *SubscribeRequest) (*SubscribeResponse, error)
	CancelSubscription(context.Context, *CancelSubscriptionRequest) (*CancelSubscriptionResponse, error)
	GetSubscription(context.Context, *GetSubscriptionRequest) (*GetSubscriptionResponse, error)
	GetSubscriptionStatus(context.Context, *GetSubscriptionStatusRequest) (*GetSubscriptionStatusResponse, error)
	mustEmbedUnimplementedSubscriptionServiceServer()
}

// UnimplementedSubscriptionServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedSubscriptionServiceServer struct{}

func (UnimplementedSubscriptionServiceServer) CreatePlan(context.Context, *CreatePlanRequest) (*CreatePlanResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatePlan not implemented")
}
func (UnimplementedSubscriptionServiceServer) GetPlanByID(context.Context, *GetPlanByIDRequest) (*GetPlanByIDResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPlanByID not implemented")
}
func (UnimplementedSubscriptionServiceServer) UpdatePlan(context.Context, *UpdatePlanRequest) (*UpdatePlanResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdatePlan not implemented")
}
func (UnimplementedSubscriptionServiceServer) GetListOfPlans(context.Context, *GetListOfPlansRequest) (*GetListOfPlansResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetListOfPlans not implemented")
}
func (UnimplementedSubscriptionServiceServer) Subscribe(context.Context, *SubscribeRequest) (*SubscribeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Subscribe not implemented")
}
func (UnimplementedSubscriptionServiceServer) CancelSubscription(context.Context, *CancelSubscriptionRequest) (*CancelSubscriptionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CancelSubscription not implemented")
}
func (UnimplementedSubscriptionServiceServer) GetSubscription(context.Context, *GetSubscriptionRequest) (*GetSubscriptionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSubscription not implemented")
}
func (UnimplementedSubscriptionServiceServer) GetSubscriptionStatus(context.Context, *GetSubscriptionStatusRequest) (*GetSubscriptionStatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSubscriptionStatus not implemented")
}
func (UnimplementedSubscriptionServiceServer) mustEmbedUnimplementedSubscriptionServiceServer() {}
func (UnimplementedSubscriptionServiceServer) testEmbeddedByValue()                             {}

// UnsafeSubscriptionServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SubscriptionServiceServer will
// result in compilation errors.
type UnsafeSubscriptionServiceServer interface {
	mustEmbedUnimplementedSubscriptionServiceServer()
}

func RegisterSubscriptionServiceServer(s grpc.ServiceRegistrar, srv SubscriptionServiceServer) {
	// If the following call pancis, it indicates UnimplementedSubscriptionServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&SubscriptionService_ServiceDesc, srv)
}

func _SubscriptionService_CreatePlan_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreatePlanRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SubscriptionServiceServer).CreatePlan(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SubscriptionService_CreatePlan_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SubscriptionServiceServer).CreatePlan(ctx, req.(*CreatePlanRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SubscriptionService_GetPlanByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPlanByIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SubscriptionServiceServer).GetPlanByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SubscriptionService_GetPlanByID_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SubscriptionServiceServer).GetPlanByID(ctx, req.(*GetPlanByIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SubscriptionService_UpdatePlan_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdatePlanRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SubscriptionServiceServer).UpdatePlan(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SubscriptionService_UpdatePlan_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SubscriptionServiceServer).UpdatePlan(ctx, req.(*UpdatePlanRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SubscriptionService_GetListOfPlans_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetListOfPlansRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SubscriptionServiceServer).GetListOfPlans(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SubscriptionService_GetListOfPlans_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SubscriptionServiceServer).GetListOfPlans(ctx, req.(*GetListOfPlansRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SubscriptionService_Subscribe_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SubscribeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SubscriptionServiceServer).Subscribe(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SubscriptionService_Subscribe_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SubscriptionServiceServer).Subscribe(ctx, req.(*SubscribeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SubscriptionService_CancelSubscription_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CancelSubscriptionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SubscriptionServiceServer).CancelSubscription(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SubscriptionService_CancelSubscription_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SubscriptionServiceServer).CancelSubscription(ctx, req.(*CancelSubscriptionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SubscriptionService_GetSubscription_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSubscriptionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SubscriptionServiceServer).GetSubscription(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SubscriptionService_GetSubscription_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SubscriptionServiceServer).GetSubscription(ctx, req.(*GetSubscriptionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SubscriptionService_GetSubscriptionStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSubscriptionStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SubscriptionServiceServer).GetSubscriptionStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SubscriptionService_GetSubscriptionStatus_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SubscriptionServiceServer).GetSubscriptionStatus(ctx, req.(*GetSubscriptionStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// SubscriptionService_ServiceDesc is the grpc.ServiceDesc for SubscriptionService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SubscriptionService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "subscription.SubscriptionService",
	HandlerType: (*SubscriptionServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreatePlan",
			Handler:    _SubscriptionService_CreatePlan_Handler,
		},
		{
			MethodName: "GetPlanByID",
			Handler:    _SubscriptionService_GetPlanByID_Handler,
		},
		{
			MethodName: "UpdatePlan",
			Handler:    _SubscriptionService_UpdatePlan_Handler,
		},
		{
			MethodName: "GetListOfPlans",
			Handler:    _SubscriptionService_GetListOfPlans_Handler,
		},
		{
			MethodName: "Subscribe",
			Handler:    _SubscriptionService_Subscribe_Handler,
		},
		{
			MethodName: "CancelSubscription",
			Handler:    _SubscriptionService_CancelSubscription_Handler,
		},
		{
			MethodName: "GetSubscription",
			Handler:    _SubscriptionService_GetSubscription_Handler,
		},
		{
			MethodName: "GetSubscriptionStatus",
			Handler:    _SubscriptionService_GetSubscriptionStatus_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "subscription.proto",
}
