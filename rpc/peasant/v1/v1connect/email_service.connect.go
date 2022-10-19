// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: peasant/v1/email_service.proto

package v1connect

import (
	context "context"
	errors "errors"
	connect_go "github.com/bufbuild/connect-go"
	v1 "github.com/ysomad/answersuck/rpc/peasant/v1"
	http "net/http"
	strings "strings"
)

// This is a compile-time assertion to ensure that this generated file and the connect package are
// compatible. If you get a compiler error that this constant is not defined, this code was
// generated with a version of connect newer than the one compiled into your binary. You can fix the
// problem by either regenerating this code with an older version of connect or updating the connect
// version compiled into your binary.
const _ = connect_go.IsAtLeastVersion0_1_0

const (
	// EmailServiceName is the fully-qualified name of the EmailService service.
	EmailServiceName = "peasant.v1.EmailService"
)

// EmailServiceClient is a client for the peasant.v1.EmailService service.
type EmailServiceClient interface {
	// VerifyEmail verifies account email by provided code.
	VerifyEmail(context.Context, *connect_go.Request[v1.VerifyEmailRequest]) (*connect_go.Response[v1.VerifyEmailResponse], error)
	// UpdateEmail updates account email.
	UpdateEmail(context.Context, *connect_go.Request[v1.UpdateEmailRequest]) (*connect_go.Response[v1.UpdateEmailResponse], error)
	// SendVerification sends email verification with new code.
	SendVerification(context.Context, *connect_go.Request[v1.SendVerificationRequest]) (*connect_go.Response[v1.SendVerificationResponse], error)
}

// NewEmailServiceClient constructs a client for the peasant.v1.EmailService service. By default, it
// uses the Connect protocol with the binary Protobuf Codec, asks for gzipped responses, and sends
// uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the connect.WithGRPC() or
// connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewEmailServiceClient(httpClient connect_go.HTTPClient, baseURL string, opts ...connect_go.ClientOption) EmailServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &emailServiceClient{
		verifyEmail: connect_go.NewClient[v1.VerifyEmailRequest, v1.VerifyEmailResponse](
			httpClient,
			baseURL+"/peasant.v1.EmailService/VerifyEmail",
			opts...,
		),
		updateEmail: connect_go.NewClient[v1.UpdateEmailRequest, v1.UpdateEmailResponse](
			httpClient,
			baseURL+"/peasant.v1.EmailService/UpdateEmail",
			opts...,
		),
		sendVerification: connect_go.NewClient[v1.SendVerificationRequest, v1.SendVerificationResponse](
			httpClient,
			baseURL+"/peasant.v1.EmailService/SendVerification",
			opts...,
		),
	}
}

// emailServiceClient implements EmailServiceClient.
type emailServiceClient struct {
	verifyEmail      *connect_go.Client[v1.VerifyEmailRequest, v1.VerifyEmailResponse]
	updateEmail      *connect_go.Client[v1.UpdateEmailRequest, v1.UpdateEmailResponse]
	sendVerification *connect_go.Client[v1.SendVerificationRequest, v1.SendVerificationResponse]
}

// VerifyEmail calls peasant.v1.EmailService.VerifyEmail.
func (c *emailServiceClient) VerifyEmail(ctx context.Context, req *connect_go.Request[v1.VerifyEmailRequest]) (*connect_go.Response[v1.VerifyEmailResponse], error) {
	return c.verifyEmail.CallUnary(ctx, req)
}

// UpdateEmail calls peasant.v1.EmailService.UpdateEmail.
func (c *emailServiceClient) UpdateEmail(ctx context.Context, req *connect_go.Request[v1.UpdateEmailRequest]) (*connect_go.Response[v1.UpdateEmailResponse], error) {
	return c.updateEmail.CallUnary(ctx, req)
}

// SendVerification calls peasant.v1.EmailService.SendVerification.
func (c *emailServiceClient) SendVerification(ctx context.Context, req *connect_go.Request[v1.SendVerificationRequest]) (*connect_go.Response[v1.SendVerificationResponse], error) {
	return c.sendVerification.CallUnary(ctx, req)
}

// EmailServiceHandler is an implementation of the peasant.v1.EmailService service.
type EmailServiceHandler interface {
	// VerifyEmail verifies account email by provided code.
	VerifyEmail(context.Context, *connect_go.Request[v1.VerifyEmailRequest]) (*connect_go.Response[v1.VerifyEmailResponse], error)
	// UpdateEmail updates account email.
	UpdateEmail(context.Context, *connect_go.Request[v1.UpdateEmailRequest]) (*connect_go.Response[v1.UpdateEmailResponse], error)
	// SendVerification sends email verification with new code.
	SendVerification(context.Context, *connect_go.Request[v1.SendVerificationRequest]) (*connect_go.Response[v1.SendVerificationResponse], error)
}

// NewEmailServiceHandler builds an HTTP handler from the service implementation. It returns the
// path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewEmailServiceHandler(svc EmailServiceHandler, opts ...connect_go.HandlerOption) (string, http.Handler) {
	mux := http.NewServeMux()
	mux.Handle("/peasant.v1.EmailService/VerifyEmail", connect_go.NewUnaryHandler(
		"/peasant.v1.EmailService/VerifyEmail",
		svc.VerifyEmail,
		opts...,
	))
	mux.Handle("/peasant.v1.EmailService/UpdateEmail", connect_go.NewUnaryHandler(
		"/peasant.v1.EmailService/UpdateEmail",
		svc.UpdateEmail,
		opts...,
	))
	mux.Handle("/peasant.v1.EmailService/SendVerification", connect_go.NewUnaryHandler(
		"/peasant.v1.EmailService/SendVerification",
		svc.SendVerification,
		opts...,
	))
	return "/peasant.v1.EmailService/", mux
}

// UnimplementedEmailServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedEmailServiceHandler struct{}

func (UnimplementedEmailServiceHandler) VerifyEmail(context.Context, *connect_go.Request[v1.VerifyEmailRequest]) (*connect_go.Response[v1.VerifyEmailResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("peasant.v1.EmailService.VerifyEmail is not implemented"))
}

func (UnimplementedEmailServiceHandler) UpdateEmail(context.Context, *connect_go.Request[v1.UpdateEmailRequest]) (*connect_go.Response[v1.UpdateEmailResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("peasant.v1.EmailService.UpdateEmail is not implemented"))
}

func (UnimplementedEmailServiceHandler) SendVerification(context.Context, *connect_go.Request[v1.SendVerificationRequest]) (*connect_go.Response[v1.SendVerificationResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("peasant.v1.EmailService.SendVerification is not implemented"))
}
