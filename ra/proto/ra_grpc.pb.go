// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package proto

import (
	context "context"
	proto "github.com/letsencrypt/boulder/core/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// RegistrationAuthorityClient is the client API for RegistrationAuthority service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RegistrationAuthorityClient interface {
	NewRegistration(ctx context.Context, in *proto.Registration, opts ...grpc.CallOption) (*proto.Registration, error)
	NewAuthorization(ctx context.Context, in *NewAuthorizationRequest, opts ...grpc.CallOption) (*proto.Authorization, error)
	NewCertificate(ctx context.Context, in *NewCertificateRequest, opts ...grpc.CallOption) (*proto.Certificate, error)
	UpdateRegistration(ctx context.Context, in *UpdateRegistrationRequest, opts ...grpc.CallOption) (*proto.Registration, error)
	PerformValidation(ctx context.Context, in *PerformValidationRequest, opts ...grpc.CallOption) (*proto.Authorization, error)
	RevokeCertificateWithReg(ctx context.Context, in *RevokeCertificateWithRegRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	DeactivateRegistration(ctx context.Context, in *proto.Registration, opts ...grpc.CallOption) (*emptypb.Empty, error)
	DeactivateAuthorization(ctx context.Context, in *proto.Authorization, opts ...grpc.CallOption) (*emptypb.Empty, error)
	AdministrativelyRevokeCertificate(ctx context.Context, in *AdministrativelyRevokeCertificateRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	NewOrder(ctx context.Context, in *NewOrderRequest, opts ...grpc.CallOption) (*proto.Order, error)
	FinalizeOrder(ctx context.Context, in *FinalizeOrderRequest, opts ...grpc.CallOption) (*proto.Order, error)
}

type registrationAuthorityClient struct {
	cc grpc.ClientConnInterface
}

func NewRegistrationAuthorityClient(cc grpc.ClientConnInterface) RegistrationAuthorityClient {
	return &registrationAuthorityClient{cc}
}

func (c *registrationAuthorityClient) NewRegistration(ctx context.Context, in *proto.Registration, opts ...grpc.CallOption) (*proto.Registration, error) {
	out := new(proto.Registration)
	err := c.cc.Invoke(ctx, "/ra.RegistrationAuthority/NewRegistration", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *registrationAuthorityClient) NewAuthorization(ctx context.Context, in *NewAuthorizationRequest, opts ...grpc.CallOption) (*proto.Authorization, error) {
	out := new(proto.Authorization)
	err := c.cc.Invoke(ctx, "/ra.RegistrationAuthority/NewAuthorization", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *registrationAuthorityClient) NewCertificate(ctx context.Context, in *NewCertificateRequest, opts ...grpc.CallOption) (*proto.Certificate, error) {
	out := new(proto.Certificate)
	err := c.cc.Invoke(ctx, "/ra.RegistrationAuthority/NewCertificate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *registrationAuthorityClient) UpdateRegistration(ctx context.Context, in *UpdateRegistrationRequest, opts ...grpc.CallOption) (*proto.Registration, error) {
	out := new(proto.Registration)
	err := c.cc.Invoke(ctx, "/ra.RegistrationAuthority/UpdateRegistration", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *registrationAuthorityClient) PerformValidation(ctx context.Context, in *PerformValidationRequest, opts ...grpc.CallOption) (*proto.Authorization, error) {
	out := new(proto.Authorization)
	err := c.cc.Invoke(ctx, "/ra.RegistrationAuthority/PerformValidation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *registrationAuthorityClient) RevokeCertificateWithReg(ctx context.Context, in *RevokeCertificateWithRegRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/ra.RegistrationAuthority/RevokeCertificateWithReg", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *registrationAuthorityClient) DeactivateRegistration(ctx context.Context, in *proto.Registration, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/ra.RegistrationAuthority/DeactivateRegistration", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *registrationAuthorityClient) DeactivateAuthorization(ctx context.Context, in *proto.Authorization, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/ra.RegistrationAuthority/DeactivateAuthorization", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *registrationAuthorityClient) AdministrativelyRevokeCertificate(ctx context.Context, in *AdministrativelyRevokeCertificateRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/ra.RegistrationAuthority/AdministrativelyRevokeCertificate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *registrationAuthorityClient) NewOrder(ctx context.Context, in *NewOrderRequest, opts ...grpc.CallOption) (*proto.Order, error) {
	out := new(proto.Order)
	err := c.cc.Invoke(ctx, "/ra.RegistrationAuthority/NewOrder", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *registrationAuthorityClient) FinalizeOrder(ctx context.Context, in *FinalizeOrderRequest, opts ...grpc.CallOption) (*proto.Order, error) {
	out := new(proto.Order)
	err := c.cc.Invoke(ctx, "/ra.RegistrationAuthority/FinalizeOrder", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RegistrationAuthorityServer is the server API for RegistrationAuthority service.
// All implementations must embed UnimplementedRegistrationAuthorityServer
// for forward compatibility
type RegistrationAuthorityServer interface {
	NewRegistration(context.Context, *proto.Registration) (*proto.Registration, error)
	NewAuthorization(context.Context, *NewAuthorizationRequest) (*proto.Authorization, error)
	NewCertificate(context.Context, *NewCertificateRequest) (*proto.Certificate, error)
	UpdateRegistration(context.Context, *UpdateRegistrationRequest) (*proto.Registration, error)
	PerformValidation(context.Context, *PerformValidationRequest) (*proto.Authorization, error)
	RevokeCertificateWithReg(context.Context, *RevokeCertificateWithRegRequest) (*emptypb.Empty, error)
	DeactivateRegistration(context.Context, *proto.Registration) (*emptypb.Empty, error)
	DeactivateAuthorization(context.Context, *proto.Authorization) (*emptypb.Empty, error)
	AdministrativelyRevokeCertificate(context.Context, *AdministrativelyRevokeCertificateRequest) (*emptypb.Empty, error)
	NewOrder(context.Context, *NewOrderRequest) (*proto.Order, error)
	FinalizeOrder(context.Context, *FinalizeOrderRequest) (*proto.Order, error)
	mustEmbedUnimplementedRegistrationAuthorityServer()
}

// UnimplementedRegistrationAuthorityServer must be embedded to have forward compatible implementations.
type UnimplementedRegistrationAuthorityServer struct {
}

func (UnimplementedRegistrationAuthorityServer) NewRegistration(context.Context, *proto.Registration) (*proto.Registration, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NewRegistration not implemented")
}
func (UnimplementedRegistrationAuthorityServer) NewAuthorization(context.Context, *NewAuthorizationRequest) (*proto.Authorization, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NewAuthorization not implemented")
}
func (UnimplementedRegistrationAuthorityServer) NewCertificate(context.Context, *NewCertificateRequest) (*proto.Certificate, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NewCertificate not implemented")
}
func (UnimplementedRegistrationAuthorityServer) UpdateRegistration(context.Context, *UpdateRegistrationRequest) (*proto.Registration, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateRegistration not implemented")
}
func (UnimplementedRegistrationAuthorityServer) PerformValidation(context.Context, *PerformValidationRequest) (*proto.Authorization, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PerformValidation not implemented")
}
func (UnimplementedRegistrationAuthorityServer) RevokeCertificateWithReg(context.Context, *RevokeCertificateWithRegRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RevokeCertificateWithReg not implemented")
}
func (UnimplementedRegistrationAuthorityServer) DeactivateRegistration(context.Context, *proto.Registration) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeactivateRegistration not implemented")
}
func (UnimplementedRegistrationAuthorityServer) DeactivateAuthorization(context.Context, *proto.Authorization) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeactivateAuthorization not implemented")
}
func (UnimplementedRegistrationAuthorityServer) AdministrativelyRevokeCertificate(context.Context, *AdministrativelyRevokeCertificateRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AdministrativelyRevokeCertificate not implemented")
}
func (UnimplementedRegistrationAuthorityServer) NewOrder(context.Context, *NewOrderRequest) (*proto.Order, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NewOrder not implemented")
}
func (UnimplementedRegistrationAuthorityServer) FinalizeOrder(context.Context, *FinalizeOrderRequest) (*proto.Order, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FinalizeOrder not implemented")
}
func (UnimplementedRegistrationAuthorityServer) mustEmbedUnimplementedRegistrationAuthorityServer() {}

// UnsafeRegistrationAuthorityServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RegistrationAuthorityServer will
// result in compilation errors.
type UnsafeRegistrationAuthorityServer interface {
	mustEmbedUnimplementedRegistrationAuthorityServer()
}

func RegisterRegistrationAuthorityServer(s grpc.ServiceRegistrar, srv RegistrationAuthorityServer) {
	s.RegisterService(&RegistrationAuthority_ServiceDesc, srv)
}

func _RegistrationAuthority_NewRegistration_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(proto.Registration)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RegistrationAuthorityServer).NewRegistration(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ra.RegistrationAuthority/NewRegistration",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RegistrationAuthorityServer).NewRegistration(ctx, req.(*proto.Registration))
	}
	return interceptor(ctx, in, info, handler)
}

func _RegistrationAuthority_NewAuthorization_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NewAuthorizationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RegistrationAuthorityServer).NewAuthorization(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ra.RegistrationAuthority/NewAuthorization",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RegistrationAuthorityServer).NewAuthorization(ctx, req.(*NewAuthorizationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RegistrationAuthority_NewCertificate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NewCertificateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RegistrationAuthorityServer).NewCertificate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ra.RegistrationAuthority/NewCertificate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RegistrationAuthorityServer).NewCertificate(ctx, req.(*NewCertificateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RegistrationAuthority_UpdateRegistration_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateRegistrationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RegistrationAuthorityServer).UpdateRegistration(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ra.RegistrationAuthority/UpdateRegistration",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RegistrationAuthorityServer).UpdateRegistration(ctx, req.(*UpdateRegistrationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RegistrationAuthority_PerformValidation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PerformValidationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RegistrationAuthorityServer).PerformValidation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ra.RegistrationAuthority/PerformValidation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RegistrationAuthorityServer).PerformValidation(ctx, req.(*PerformValidationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RegistrationAuthority_RevokeCertificateWithReg_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RevokeCertificateWithRegRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RegistrationAuthorityServer).RevokeCertificateWithReg(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ra.RegistrationAuthority/RevokeCertificateWithReg",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RegistrationAuthorityServer).RevokeCertificateWithReg(ctx, req.(*RevokeCertificateWithRegRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RegistrationAuthority_DeactivateRegistration_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(proto.Registration)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RegistrationAuthorityServer).DeactivateRegistration(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ra.RegistrationAuthority/DeactivateRegistration",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RegistrationAuthorityServer).DeactivateRegistration(ctx, req.(*proto.Registration))
	}
	return interceptor(ctx, in, info, handler)
}

func _RegistrationAuthority_DeactivateAuthorization_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(proto.Authorization)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RegistrationAuthorityServer).DeactivateAuthorization(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ra.RegistrationAuthority/DeactivateAuthorization",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RegistrationAuthorityServer).DeactivateAuthorization(ctx, req.(*proto.Authorization))
	}
	return interceptor(ctx, in, info, handler)
}

func _RegistrationAuthority_AdministrativelyRevokeCertificate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AdministrativelyRevokeCertificateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RegistrationAuthorityServer).AdministrativelyRevokeCertificate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ra.RegistrationAuthority/AdministrativelyRevokeCertificate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RegistrationAuthorityServer).AdministrativelyRevokeCertificate(ctx, req.(*AdministrativelyRevokeCertificateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RegistrationAuthority_NewOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NewOrderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RegistrationAuthorityServer).NewOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ra.RegistrationAuthority/NewOrder",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RegistrationAuthorityServer).NewOrder(ctx, req.(*NewOrderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RegistrationAuthority_FinalizeOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FinalizeOrderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RegistrationAuthorityServer).FinalizeOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ra.RegistrationAuthority/FinalizeOrder",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RegistrationAuthorityServer).FinalizeOrder(ctx, req.(*FinalizeOrderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// RegistrationAuthority_ServiceDesc is the grpc.ServiceDesc for RegistrationAuthority service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RegistrationAuthority_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "ra.RegistrationAuthority",
	HandlerType: (*RegistrationAuthorityServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "NewRegistration",
			Handler:    _RegistrationAuthority_NewRegistration_Handler,
		},
		{
			MethodName: "NewAuthorization",
			Handler:    _RegistrationAuthority_NewAuthorization_Handler,
		},
		{
			MethodName: "NewCertificate",
			Handler:    _RegistrationAuthority_NewCertificate_Handler,
		},
		{
			MethodName: "UpdateRegistration",
			Handler:    _RegistrationAuthority_UpdateRegistration_Handler,
		},
		{
			MethodName: "PerformValidation",
			Handler:    _RegistrationAuthority_PerformValidation_Handler,
		},
		{
			MethodName: "RevokeCertificateWithReg",
			Handler:    _RegistrationAuthority_RevokeCertificateWithReg_Handler,
		},
		{
			MethodName: "DeactivateRegistration",
			Handler:    _RegistrationAuthority_DeactivateRegistration_Handler,
		},
		{
			MethodName: "DeactivateAuthorization",
			Handler:    _RegistrationAuthority_DeactivateAuthorization_Handler,
		},
		{
			MethodName: "AdministrativelyRevokeCertificate",
			Handler:    _RegistrationAuthority_AdministrativelyRevokeCertificate_Handler,
		},
		{
			MethodName: "NewOrder",
			Handler:    _RegistrationAuthority_NewOrder_Handler,
		},
		{
			MethodName: "FinalizeOrder",
			Handler:    _RegistrationAuthority_FinalizeOrder_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "ra.proto",
}
