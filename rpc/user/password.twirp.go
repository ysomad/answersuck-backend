// Code generated by protoc-gen-twirp v8.1.2, DO NOT EDIT.
// source: user/password.proto

package user

import context "context"
import fmt "fmt"
import http "net/http"
import ioutil "io/ioutil"
import json "encoding/json"
import strconv "strconv"
import strings "strings"

import protojson "google.golang.org/protobuf/encoding/protojson"
import proto "google.golang.org/protobuf/proto"
import twirp "github.com/twitchtv/twirp"
import ctxsetters "github.com/twitchtv/twirp/ctxsetters"

import google_protobuf "google.golang.org/protobuf/types/known/emptypb"

// Version compatibility assertion.
// If the constant is not defined in the package, that likely means
// the package needs to be updated to work with this generated code.
// See https://twitchtv.github.io/twirp/docs/version_matrix.html
const _ = twirp.TwirpPackageMinVersion_8_1_0

// =========================
// PasswordService Interface
// =========================

type PasswordService interface {
	// ResetPassword sends an email to account email with short-term token
	// which may be used for updating account password.
	ResetPassword(context.Context, *ResetPasswordRequest) (*google_protobuf.Empty, error)

	// SetPassword sets new account password by using token from ResetPassword.
	SetPassword(context.Context, *SetPasswordRequest) (*google_protobuf.Empty, error)

	// UpdatePassword updates account password if old one is correct.
	UpdatePassword(context.Context, *UpdatePasswordRequest) (*google_protobuf.Empty, error)
}

// ===============================
// PasswordService Protobuf Client
// ===============================

type passwordServiceProtobufClient struct {
	client      HTTPClient
	urls        [3]string
	interceptor twirp.Interceptor
	opts        twirp.ClientOptions
}

// NewPasswordServiceProtobufClient creates a Protobuf client that implements the PasswordService interface.
// It communicates using Protobuf and can be configured with a custom HTTPClient.
func NewPasswordServiceProtobufClient(baseURL string, client HTTPClient, opts ...twirp.ClientOption) PasswordService {
	if c, ok := client.(*http.Client); ok {
		client = withoutRedirects(c)
	}

	clientOpts := twirp.ClientOptions{}
	for _, o := range opts {
		o(&clientOpts)
	}

	// Using ReadOpt allows backwards and forwads compatibility with new options in the future
	literalURLs := false
	_ = clientOpts.ReadOpt("literalURLs", &literalURLs)
	var pathPrefix string
	if ok := clientOpts.ReadOpt("pathPrefix", &pathPrefix); !ok {
		pathPrefix = "/twirp" // default prefix
	}

	// Build method URLs: <baseURL>[<prefix>]/<package>.<Service>/<Method>
	serviceURL := sanitizeBaseURL(baseURL)
	serviceURL += baseServicePath(pathPrefix, "users.password", "PasswordService")
	urls := [3]string{
		serviceURL + "ResetPassword",
		serviceURL + "SetPassword",
		serviceURL + "UpdatePassword",
	}

	return &passwordServiceProtobufClient{
		client:      client,
		urls:        urls,
		interceptor: twirp.ChainInterceptors(clientOpts.Interceptors...),
		opts:        clientOpts,
	}
}

func (c *passwordServiceProtobufClient) ResetPassword(ctx context.Context, in *ResetPasswordRequest) (*google_protobuf.Empty, error) {
	ctx = ctxsetters.WithPackageName(ctx, "users.password")
	ctx = ctxsetters.WithServiceName(ctx, "PasswordService")
	ctx = ctxsetters.WithMethodName(ctx, "ResetPassword")
	caller := c.callResetPassword
	if c.interceptor != nil {
		caller = func(ctx context.Context, req *ResetPasswordRequest) (*google_protobuf.Empty, error) {
			resp, err := c.interceptor(
				func(ctx context.Context, req interface{}) (interface{}, error) {
					typedReq, ok := req.(*ResetPasswordRequest)
					if !ok {
						return nil, twirp.InternalError("failed type assertion req.(*ResetPasswordRequest) when calling interceptor")
					}
					return c.callResetPassword(ctx, typedReq)
				},
			)(ctx, req)
			if resp != nil {
				typedResp, ok := resp.(*google_protobuf.Empty)
				if !ok {
					return nil, twirp.InternalError("failed type assertion resp.(*google_protobuf.Empty) when calling interceptor")
				}
				return typedResp, err
			}
			return nil, err
		}
	}
	return caller(ctx, in)
}

func (c *passwordServiceProtobufClient) callResetPassword(ctx context.Context, in *ResetPasswordRequest) (*google_protobuf.Empty, error) {
	out := new(google_protobuf.Empty)
	ctx, err := doProtobufRequest(ctx, c.client, c.opts.Hooks, c.urls[0], in, out)
	if err != nil {
		twerr, ok := err.(twirp.Error)
		if !ok {
			twerr = twirp.InternalErrorWith(err)
		}
		callClientError(ctx, c.opts.Hooks, twerr)
		return nil, err
	}

	callClientResponseReceived(ctx, c.opts.Hooks)

	return out, nil
}

func (c *passwordServiceProtobufClient) SetPassword(ctx context.Context, in *SetPasswordRequest) (*google_protobuf.Empty, error) {
	ctx = ctxsetters.WithPackageName(ctx, "users.password")
	ctx = ctxsetters.WithServiceName(ctx, "PasswordService")
	ctx = ctxsetters.WithMethodName(ctx, "SetPassword")
	caller := c.callSetPassword
	if c.interceptor != nil {
		caller = func(ctx context.Context, req *SetPasswordRequest) (*google_protobuf.Empty, error) {
			resp, err := c.interceptor(
				func(ctx context.Context, req interface{}) (interface{}, error) {
					typedReq, ok := req.(*SetPasswordRequest)
					if !ok {
						return nil, twirp.InternalError("failed type assertion req.(*SetPasswordRequest) when calling interceptor")
					}
					return c.callSetPassword(ctx, typedReq)
				},
			)(ctx, req)
			if resp != nil {
				typedResp, ok := resp.(*google_protobuf.Empty)
				if !ok {
					return nil, twirp.InternalError("failed type assertion resp.(*google_protobuf.Empty) when calling interceptor")
				}
				return typedResp, err
			}
			return nil, err
		}
	}
	return caller(ctx, in)
}

func (c *passwordServiceProtobufClient) callSetPassword(ctx context.Context, in *SetPasswordRequest) (*google_protobuf.Empty, error) {
	out := new(google_protobuf.Empty)
	ctx, err := doProtobufRequest(ctx, c.client, c.opts.Hooks, c.urls[1], in, out)
	if err != nil {
		twerr, ok := err.(twirp.Error)
		if !ok {
			twerr = twirp.InternalErrorWith(err)
		}
		callClientError(ctx, c.opts.Hooks, twerr)
		return nil, err
	}

	callClientResponseReceived(ctx, c.opts.Hooks)

	return out, nil
}

func (c *passwordServiceProtobufClient) UpdatePassword(ctx context.Context, in *UpdatePasswordRequest) (*google_protobuf.Empty, error) {
	ctx = ctxsetters.WithPackageName(ctx, "users.password")
	ctx = ctxsetters.WithServiceName(ctx, "PasswordService")
	ctx = ctxsetters.WithMethodName(ctx, "UpdatePassword")
	caller := c.callUpdatePassword
	if c.interceptor != nil {
		caller = func(ctx context.Context, req *UpdatePasswordRequest) (*google_protobuf.Empty, error) {
			resp, err := c.interceptor(
				func(ctx context.Context, req interface{}) (interface{}, error) {
					typedReq, ok := req.(*UpdatePasswordRequest)
					if !ok {
						return nil, twirp.InternalError("failed type assertion req.(*UpdatePasswordRequest) when calling interceptor")
					}
					return c.callUpdatePassword(ctx, typedReq)
				},
			)(ctx, req)
			if resp != nil {
				typedResp, ok := resp.(*google_protobuf.Empty)
				if !ok {
					return nil, twirp.InternalError("failed type assertion resp.(*google_protobuf.Empty) when calling interceptor")
				}
				return typedResp, err
			}
			return nil, err
		}
	}
	return caller(ctx, in)
}

func (c *passwordServiceProtobufClient) callUpdatePassword(ctx context.Context, in *UpdatePasswordRequest) (*google_protobuf.Empty, error) {
	out := new(google_protobuf.Empty)
	ctx, err := doProtobufRequest(ctx, c.client, c.opts.Hooks, c.urls[2], in, out)
	if err != nil {
		twerr, ok := err.(twirp.Error)
		if !ok {
			twerr = twirp.InternalErrorWith(err)
		}
		callClientError(ctx, c.opts.Hooks, twerr)
		return nil, err
	}

	callClientResponseReceived(ctx, c.opts.Hooks)

	return out, nil
}

// ===========================
// PasswordService JSON Client
// ===========================

type passwordServiceJSONClient struct {
	client      HTTPClient
	urls        [3]string
	interceptor twirp.Interceptor
	opts        twirp.ClientOptions
}

// NewPasswordServiceJSONClient creates a JSON client that implements the PasswordService interface.
// It communicates using JSON and can be configured with a custom HTTPClient.
func NewPasswordServiceJSONClient(baseURL string, client HTTPClient, opts ...twirp.ClientOption) PasswordService {
	if c, ok := client.(*http.Client); ok {
		client = withoutRedirects(c)
	}

	clientOpts := twirp.ClientOptions{}
	for _, o := range opts {
		o(&clientOpts)
	}

	// Using ReadOpt allows backwards and forwads compatibility with new options in the future
	literalURLs := false
	_ = clientOpts.ReadOpt("literalURLs", &literalURLs)
	var pathPrefix string
	if ok := clientOpts.ReadOpt("pathPrefix", &pathPrefix); !ok {
		pathPrefix = "/twirp" // default prefix
	}

	// Build method URLs: <baseURL>[<prefix>]/<package>.<Service>/<Method>
	serviceURL := sanitizeBaseURL(baseURL)
	serviceURL += baseServicePath(pathPrefix, "users.password", "PasswordService")
	urls := [3]string{
		serviceURL + "ResetPassword",
		serviceURL + "SetPassword",
		serviceURL + "UpdatePassword",
	}

	return &passwordServiceJSONClient{
		client:      client,
		urls:        urls,
		interceptor: twirp.ChainInterceptors(clientOpts.Interceptors...),
		opts:        clientOpts,
	}
}

func (c *passwordServiceJSONClient) ResetPassword(ctx context.Context, in *ResetPasswordRequest) (*google_protobuf.Empty, error) {
	ctx = ctxsetters.WithPackageName(ctx, "users.password")
	ctx = ctxsetters.WithServiceName(ctx, "PasswordService")
	ctx = ctxsetters.WithMethodName(ctx, "ResetPassword")
	caller := c.callResetPassword
	if c.interceptor != nil {
		caller = func(ctx context.Context, req *ResetPasswordRequest) (*google_protobuf.Empty, error) {
			resp, err := c.interceptor(
				func(ctx context.Context, req interface{}) (interface{}, error) {
					typedReq, ok := req.(*ResetPasswordRequest)
					if !ok {
						return nil, twirp.InternalError("failed type assertion req.(*ResetPasswordRequest) when calling interceptor")
					}
					return c.callResetPassword(ctx, typedReq)
				},
			)(ctx, req)
			if resp != nil {
				typedResp, ok := resp.(*google_protobuf.Empty)
				if !ok {
					return nil, twirp.InternalError("failed type assertion resp.(*google_protobuf.Empty) when calling interceptor")
				}
				return typedResp, err
			}
			return nil, err
		}
	}
	return caller(ctx, in)
}

func (c *passwordServiceJSONClient) callResetPassword(ctx context.Context, in *ResetPasswordRequest) (*google_protobuf.Empty, error) {
	out := new(google_protobuf.Empty)
	ctx, err := doJSONRequest(ctx, c.client, c.opts.Hooks, c.urls[0], in, out)
	if err != nil {
		twerr, ok := err.(twirp.Error)
		if !ok {
			twerr = twirp.InternalErrorWith(err)
		}
		callClientError(ctx, c.opts.Hooks, twerr)
		return nil, err
	}

	callClientResponseReceived(ctx, c.opts.Hooks)

	return out, nil
}

func (c *passwordServiceJSONClient) SetPassword(ctx context.Context, in *SetPasswordRequest) (*google_protobuf.Empty, error) {
	ctx = ctxsetters.WithPackageName(ctx, "users.password")
	ctx = ctxsetters.WithServiceName(ctx, "PasswordService")
	ctx = ctxsetters.WithMethodName(ctx, "SetPassword")
	caller := c.callSetPassword
	if c.interceptor != nil {
		caller = func(ctx context.Context, req *SetPasswordRequest) (*google_protobuf.Empty, error) {
			resp, err := c.interceptor(
				func(ctx context.Context, req interface{}) (interface{}, error) {
					typedReq, ok := req.(*SetPasswordRequest)
					if !ok {
						return nil, twirp.InternalError("failed type assertion req.(*SetPasswordRequest) when calling interceptor")
					}
					return c.callSetPassword(ctx, typedReq)
				},
			)(ctx, req)
			if resp != nil {
				typedResp, ok := resp.(*google_protobuf.Empty)
				if !ok {
					return nil, twirp.InternalError("failed type assertion resp.(*google_protobuf.Empty) when calling interceptor")
				}
				return typedResp, err
			}
			return nil, err
		}
	}
	return caller(ctx, in)
}

func (c *passwordServiceJSONClient) callSetPassword(ctx context.Context, in *SetPasswordRequest) (*google_protobuf.Empty, error) {
	out := new(google_protobuf.Empty)
	ctx, err := doJSONRequest(ctx, c.client, c.opts.Hooks, c.urls[1], in, out)
	if err != nil {
		twerr, ok := err.(twirp.Error)
		if !ok {
			twerr = twirp.InternalErrorWith(err)
		}
		callClientError(ctx, c.opts.Hooks, twerr)
		return nil, err
	}

	callClientResponseReceived(ctx, c.opts.Hooks)

	return out, nil
}

func (c *passwordServiceJSONClient) UpdatePassword(ctx context.Context, in *UpdatePasswordRequest) (*google_protobuf.Empty, error) {
	ctx = ctxsetters.WithPackageName(ctx, "users.password")
	ctx = ctxsetters.WithServiceName(ctx, "PasswordService")
	ctx = ctxsetters.WithMethodName(ctx, "UpdatePassword")
	caller := c.callUpdatePassword
	if c.interceptor != nil {
		caller = func(ctx context.Context, req *UpdatePasswordRequest) (*google_protobuf.Empty, error) {
			resp, err := c.interceptor(
				func(ctx context.Context, req interface{}) (interface{}, error) {
					typedReq, ok := req.(*UpdatePasswordRequest)
					if !ok {
						return nil, twirp.InternalError("failed type assertion req.(*UpdatePasswordRequest) when calling interceptor")
					}
					return c.callUpdatePassword(ctx, typedReq)
				},
			)(ctx, req)
			if resp != nil {
				typedResp, ok := resp.(*google_protobuf.Empty)
				if !ok {
					return nil, twirp.InternalError("failed type assertion resp.(*google_protobuf.Empty) when calling interceptor")
				}
				return typedResp, err
			}
			return nil, err
		}
	}
	return caller(ctx, in)
}

func (c *passwordServiceJSONClient) callUpdatePassword(ctx context.Context, in *UpdatePasswordRequest) (*google_protobuf.Empty, error) {
	out := new(google_protobuf.Empty)
	ctx, err := doJSONRequest(ctx, c.client, c.opts.Hooks, c.urls[2], in, out)
	if err != nil {
		twerr, ok := err.(twirp.Error)
		if !ok {
			twerr = twirp.InternalErrorWith(err)
		}
		callClientError(ctx, c.opts.Hooks, twerr)
		return nil, err
	}

	callClientResponseReceived(ctx, c.opts.Hooks)

	return out, nil
}

// ==============================
// PasswordService Server Handler
// ==============================

type passwordServiceServer struct {
	PasswordService
	interceptor      twirp.Interceptor
	hooks            *twirp.ServerHooks
	pathPrefix       string // prefix for routing
	jsonSkipDefaults bool   // do not include unpopulated fields (default values) in the response
	jsonCamelCase    bool   // JSON fields are serialized as lowerCamelCase rather than keeping the original proto names
}

// NewPasswordServiceServer builds a TwirpServer that can be used as an http.Handler to handle
// HTTP requests that are routed to the right method in the provided svc implementation.
// The opts are twirp.ServerOption modifiers, for example twirp.WithServerHooks(hooks).
func NewPasswordServiceServer(svc PasswordService, opts ...interface{}) TwirpServer {
	serverOpts := newServerOpts(opts)

	// Using ReadOpt allows backwards and forwads compatibility with new options in the future
	jsonSkipDefaults := false
	_ = serverOpts.ReadOpt("jsonSkipDefaults", &jsonSkipDefaults)
	jsonCamelCase := false
	_ = serverOpts.ReadOpt("jsonCamelCase", &jsonCamelCase)
	var pathPrefix string
	if ok := serverOpts.ReadOpt("pathPrefix", &pathPrefix); !ok {
		pathPrefix = "/twirp" // default prefix
	}

	return &passwordServiceServer{
		PasswordService:  svc,
		hooks:            serverOpts.Hooks,
		interceptor:      twirp.ChainInterceptors(serverOpts.Interceptors...),
		pathPrefix:       pathPrefix,
		jsonSkipDefaults: jsonSkipDefaults,
		jsonCamelCase:    jsonCamelCase,
	}
}

// writeError writes an HTTP response with a valid Twirp error format, and triggers hooks.
// If err is not a twirp.Error, it will get wrapped with twirp.InternalErrorWith(err)
func (s *passwordServiceServer) writeError(ctx context.Context, resp http.ResponseWriter, err error) {
	writeError(ctx, resp, err, s.hooks)
}

// handleRequestBodyError is used to handle error when the twirp server cannot read request
func (s *passwordServiceServer) handleRequestBodyError(ctx context.Context, resp http.ResponseWriter, msg string, err error) {
	if context.Canceled == ctx.Err() {
		s.writeError(ctx, resp, twirp.NewError(twirp.Canceled, "failed to read request: context canceled"))
		return
	}
	if context.DeadlineExceeded == ctx.Err() {
		s.writeError(ctx, resp, twirp.NewError(twirp.DeadlineExceeded, "failed to read request: deadline exceeded"))
		return
	}
	s.writeError(ctx, resp, twirp.WrapError(malformedRequestError(msg), err))
}

// PasswordServicePathPrefix is a convenience constant that may identify URL paths.
// Should be used with caution, it only matches routes generated by Twirp Go clients,
// with the default "/twirp" prefix and default CamelCase service and method names.
// More info: https://twitchtv.github.io/twirp/docs/routing.html
const PasswordServicePathPrefix = "/twirp/users.password.PasswordService/"

func (s *passwordServiceServer) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	ctx := req.Context()
	ctx = ctxsetters.WithPackageName(ctx, "users.password")
	ctx = ctxsetters.WithServiceName(ctx, "PasswordService")
	ctx = ctxsetters.WithResponseWriter(ctx, resp)

	var err error
	ctx, err = callRequestReceived(ctx, s.hooks)
	if err != nil {
		s.writeError(ctx, resp, err)
		return
	}

	if req.Method != "POST" {
		msg := fmt.Sprintf("unsupported method %q (only POST is allowed)", req.Method)
		s.writeError(ctx, resp, badRouteError(msg, req.Method, req.URL.Path))
		return
	}

	// Verify path format: [<prefix>]/<package>.<Service>/<Method>
	prefix, pkgService, method := parseTwirpPath(req.URL.Path)
	if pkgService != "users.password.PasswordService" {
		msg := fmt.Sprintf("no handler for path %q", req.URL.Path)
		s.writeError(ctx, resp, badRouteError(msg, req.Method, req.URL.Path))
		return
	}
	if prefix != s.pathPrefix {
		msg := fmt.Sprintf("invalid path prefix %q, expected %q, on path %q", prefix, s.pathPrefix, req.URL.Path)
		s.writeError(ctx, resp, badRouteError(msg, req.Method, req.URL.Path))
		return
	}

	switch method {
	case "ResetPassword":
		s.serveResetPassword(ctx, resp, req)
		return
	case "SetPassword":
		s.serveSetPassword(ctx, resp, req)
		return
	case "UpdatePassword":
		s.serveUpdatePassword(ctx, resp, req)
		return
	default:
		msg := fmt.Sprintf("no handler for path %q", req.URL.Path)
		s.writeError(ctx, resp, badRouteError(msg, req.Method, req.URL.Path))
		return
	}
}

func (s *passwordServiceServer) serveResetPassword(ctx context.Context, resp http.ResponseWriter, req *http.Request) {
	header := req.Header.Get("Content-Type")
	i := strings.Index(header, ";")
	if i == -1 {
		i = len(header)
	}
	switch strings.TrimSpace(strings.ToLower(header[:i])) {
	case "application/json":
		s.serveResetPasswordJSON(ctx, resp, req)
	case "application/protobuf":
		s.serveResetPasswordProtobuf(ctx, resp, req)
	default:
		msg := fmt.Sprintf("unexpected Content-Type: %q", req.Header.Get("Content-Type"))
		twerr := badRouteError(msg, req.Method, req.URL.Path)
		s.writeError(ctx, resp, twerr)
	}
}

func (s *passwordServiceServer) serveResetPasswordJSON(ctx context.Context, resp http.ResponseWriter, req *http.Request) {
	var err error
	ctx = ctxsetters.WithMethodName(ctx, "ResetPassword")
	ctx, err = callRequestRouted(ctx, s.hooks)
	if err != nil {
		s.writeError(ctx, resp, err)
		return
	}

	d := json.NewDecoder(req.Body)
	rawReqBody := json.RawMessage{}
	if err := d.Decode(&rawReqBody); err != nil {
		s.handleRequestBodyError(ctx, resp, "the json request could not be decoded", err)
		return
	}
	reqContent := new(ResetPasswordRequest)
	unmarshaler := protojson.UnmarshalOptions{DiscardUnknown: true}
	if err = unmarshaler.Unmarshal(rawReqBody, reqContent); err != nil {
		s.handleRequestBodyError(ctx, resp, "the json request could not be decoded", err)
		return
	}

	handler := s.PasswordService.ResetPassword
	if s.interceptor != nil {
		handler = func(ctx context.Context, req *ResetPasswordRequest) (*google_protobuf.Empty, error) {
			resp, err := s.interceptor(
				func(ctx context.Context, req interface{}) (interface{}, error) {
					typedReq, ok := req.(*ResetPasswordRequest)
					if !ok {
						return nil, twirp.InternalError("failed type assertion req.(*ResetPasswordRequest) when calling interceptor")
					}
					return s.PasswordService.ResetPassword(ctx, typedReq)
				},
			)(ctx, req)
			if resp != nil {
				typedResp, ok := resp.(*google_protobuf.Empty)
				if !ok {
					return nil, twirp.InternalError("failed type assertion resp.(*google_protobuf.Empty) when calling interceptor")
				}
				return typedResp, err
			}
			return nil, err
		}
	}

	// Call service method
	var respContent *google_protobuf.Empty
	func() {
		defer ensurePanicResponses(ctx, resp, s.hooks)
		respContent, err = handler(ctx, reqContent)
	}()

	if err != nil {
		s.writeError(ctx, resp, err)
		return
	}
	if respContent == nil {
		s.writeError(ctx, resp, twirp.InternalError("received a nil *google_protobuf.Empty and nil error while calling ResetPassword. nil responses are not supported"))
		return
	}

	ctx = callResponsePrepared(ctx, s.hooks)

	marshaler := &protojson.MarshalOptions{UseProtoNames: !s.jsonCamelCase, EmitUnpopulated: !s.jsonSkipDefaults}
	respBytes, err := marshaler.Marshal(respContent)
	if err != nil {
		s.writeError(ctx, resp, wrapInternal(err, "failed to marshal json response"))
		return
	}

	ctx = ctxsetters.WithStatusCode(ctx, http.StatusOK)
	resp.Header().Set("Content-Type", "application/json")
	resp.Header().Set("Content-Length", strconv.Itoa(len(respBytes)))
	resp.WriteHeader(http.StatusOK)

	if n, err := resp.Write(respBytes); err != nil {
		msg := fmt.Sprintf("failed to write response, %d of %d bytes written: %s", n, len(respBytes), err.Error())
		twerr := twirp.NewError(twirp.Unknown, msg)
		ctx = callError(ctx, s.hooks, twerr)
	}
	callResponseSent(ctx, s.hooks)
}

func (s *passwordServiceServer) serveResetPasswordProtobuf(ctx context.Context, resp http.ResponseWriter, req *http.Request) {
	var err error
	ctx = ctxsetters.WithMethodName(ctx, "ResetPassword")
	ctx, err = callRequestRouted(ctx, s.hooks)
	if err != nil {
		s.writeError(ctx, resp, err)
		return
	}

	buf, err := ioutil.ReadAll(req.Body)
	if err != nil {
		s.handleRequestBodyError(ctx, resp, "failed to read request body", err)
		return
	}
	reqContent := new(ResetPasswordRequest)
	if err = proto.Unmarshal(buf, reqContent); err != nil {
		s.writeError(ctx, resp, malformedRequestError("the protobuf request could not be decoded"))
		return
	}

	handler := s.PasswordService.ResetPassword
	if s.interceptor != nil {
		handler = func(ctx context.Context, req *ResetPasswordRequest) (*google_protobuf.Empty, error) {
			resp, err := s.interceptor(
				func(ctx context.Context, req interface{}) (interface{}, error) {
					typedReq, ok := req.(*ResetPasswordRequest)
					if !ok {
						return nil, twirp.InternalError("failed type assertion req.(*ResetPasswordRequest) when calling interceptor")
					}
					return s.PasswordService.ResetPassword(ctx, typedReq)
				},
			)(ctx, req)
			if resp != nil {
				typedResp, ok := resp.(*google_protobuf.Empty)
				if !ok {
					return nil, twirp.InternalError("failed type assertion resp.(*google_protobuf.Empty) when calling interceptor")
				}
				return typedResp, err
			}
			return nil, err
		}
	}

	// Call service method
	var respContent *google_protobuf.Empty
	func() {
		defer ensurePanicResponses(ctx, resp, s.hooks)
		respContent, err = handler(ctx, reqContent)
	}()

	if err != nil {
		s.writeError(ctx, resp, err)
		return
	}
	if respContent == nil {
		s.writeError(ctx, resp, twirp.InternalError("received a nil *google_protobuf.Empty and nil error while calling ResetPassword. nil responses are not supported"))
		return
	}

	ctx = callResponsePrepared(ctx, s.hooks)

	respBytes, err := proto.Marshal(respContent)
	if err != nil {
		s.writeError(ctx, resp, wrapInternal(err, "failed to marshal proto response"))
		return
	}

	ctx = ctxsetters.WithStatusCode(ctx, http.StatusOK)
	resp.Header().Set("Content-Type", "application/protobuf")
	resp.Header().Set("Content-Length", strconv.Itoa(len(respBytes)))
	resp.WriteHeader(http.StatusOK)
	if n, err := resp.Write(respBytes); err != nil {
		msg := fmt.Sprintf("failed to write response, %d of %d bytes written: %s", n, len(respBytes), err.Error())
		twerr := twirp.NewError(twirp.Unknown, msg)
		ctx = callError(ctx, s.hooks, twerr)
	}
	callResponseSent(ctx, s.hooks)
}

func (s *passwordServiceServer) serveSetPassword(ctx context.Context, resp http.ResponseWriter, req *http.Request) {
	header := req.Header.Get("Content-Type")
	i := strings.Index(header, ";")
	if i == -1 {
		i = len(header)
	}
	switch strings.TrimSpace(strings.ToLower(header[:i])) {
	case "application/json":
		s.serveSetPasswordJSON(ctx, resp, req)
	case "application/protobuf":
		s.serveSetPasswordProtobuf(ctx, resp, req)
	default:
		msg := fmt.Sprintf("unexpected Content-Type: %q", req.Header.Get("Content-Type"))
		twerr := badRouteError(msg, req.Method, req.URL.Path)
		s.writeError(ctx, resp, twerr)
	}
}

func (s *passwordServiceServer) serveSetPasswordJSON(ctx context.Context, resp http.ResponseWriter, req *http.Request) {
	var err error
	ctx = ctxsetters.WithMethodName(ctx, "SetPassword")
	ctx, err = callRequestRouted(ctx, s.hooks)
	if err != nil {
		s.writeError(ctx, resp, err)
		return
	}

	d := json.NewDecoder(req.Body)
	rawReqBody := json.RawMessage{}
	if err := d.Decode(&rawReqBody); err != nil {
		s.handleRequestBodyError(ctx, resp, "the json request could not be decoded", err)
		return
	}
	reqContent := new(SetPasswordRequest)
	unmarshaler := protojson.UnmarshalOptions{DiscardUnknown: true}
	if err = unmarshaler.Unmarshal(rawReqBody, reqContent); err != nil {
		s.handleRequestBodyError(ctx, resp, "the json request could not be decoded", err)
		return
	}

	handler := s.PasswordService.SetPassword
	if s.interceptor != nil {
		handler = func(ctx context.Context, req *SetPasswordRequest) (*google_protobuf.Empty, error) {
			resp, err := s.interceptor(
				func(ctx context.Context, req interface{}) (interface{}, error) {
					typedReq, ok := req.(*SetPasswordRequest)
					if !ok {
						return nil, twirp.InternalError("failed type assertion req.(*SetPasswordRequest) when calling interceptor")
					}
					return s.PasswordService.SetPassword(ctx, typedReq)
				},
			)(ctx, req)
			if resp != nil {
				typedResp, ok := resp.(*google_protobuf.Empty)
				if !ok {
					return nil, twirp.InternalError("failed type assertion resp.(*google_protobuf.Empty) when calling interceptor")
				}
				return typedResp, err
			}
			return nil, err
		}
	}

	// Call service method
	var respContent *google_protobuf.Empty
	func() {
		defer ensurePanicResponses(ctx, resp, s.hooks)
		respContent, err = handler(ctx, reqContent)
	}()

	if err != nil {
		s.writeError(ctx, resp, err)
		return
	}
	if respContent == nil {
		s.writeError(ctx, resp, twirp.InternalError("received a nil *google_protobuf.Empty and nil error while calling SetPassword. nil responses are not supported"))
		return
	}

	ctx = callResponsePrepared(ctx, s.hooks)

	marshaler := &protojson.MarshalOptions{UseProtoNames: !s.jsonCamelCase, EmitUnpopulated: !s.jsonSkipDefaults}
	respBytes, err := marshaler.Marshal(respContent)
	if err != nil {
		s.writeError(ctx, resp, wrapInternal(err, "failed to marshal json response"))
		return
	}

	ctx = ctxsetters.WithStatusCode(ctx, http.StatusOK)
	resp.Header().Set("Content-Type", "application/json")
	resp.Header().Set("Content-Length", strconv.Itoa(len(respBytes)))
	resp.WriteHeader(http.StatusOK)

	if n, err := resp.Write(respBytes); err != nil {
		msg := fmt.Sprintf("failed to write response, %d of %d bytes written: %s", n, len(respBytes), err.Error())
		twerr := twirp.NewError(twirp.Unknown, msg)
		ctx = callError(ctx, s.hooks, twerr)
	}
	callResponseSent(ctx, s.hooks)
}

func (s *passwordServiceServer) serveSetPasswordProtobuf(ctx context.Context, resp http.ResponseWriter, req *http.Request) {
	var err error
	ctx = ctxsetters.WithMethodName(ctx, "SetPassword")
	ctx, err = callRequestRouted(ctx, s.hooks)
	if err != nil {
		s.writeError(ctx, resp, err)
		return
	}

	buf, err := ioutil.ReadAll(req.Body)
	if err != nil {
		s.handleRequestBodyError(ctx, resp, "failed to read request body", err)
		return
	}
	reqContent := new(SetPasswordRequest)
	if err = proto.Unmarshal(buf, reqContent); err != nil {
		s.writeError(ctx, resp, malformedRequestError("the protobuf request could not be decoded"))
		return
	}

	handler := s.PasswordService.SetPassword
	if s.interceptor != nil {
		handler = func(ctx context.Context, req *SetPasswordRequest) (*google_protobuf.Empty, error) {
			resp, err := s.interceptor(
				func(ctx context.Context, req interface{}) (interface{}, error) {
					typedReq, ok := req.(*SetPasswordRequest)
					if !ok {
						return nil, twirp.InternalError("failed type assertion req.(*SetPasswordRequest) when calling interceptor")
					}
					return s.PasswordService.SetPassword(ctx, typedReq)
				},
			)(ctx, req)
			if resp != nil {
				typedResp, ok := resp.(*google_protobuf.Empty)
				if !ok {
					return nil, twirp.InternalError("failed type assertion resp.(*google_protobuf.Empty) when calling interceptor")
				}
				return typedResp, err
			}
			return nil, err
		}
	}

	// Call service method
	var respContent *google_protobuf.Empty
	func() {
		defer ensurePanicResponses(ctx, resp, s.hooks)
		respContent, err = handler(ctx, reqContent)
	}()

	if err != nil {
		s.writeError(ctx, resp, err)
		return
	}
	if respContent == nil {
		s.writeError(ctx, resp, twirp.InternalError("received a nil *google_protobuf.Empty and nil error while calling SetPassword. nil responses are not supported"))
		return
	}

	ctx = callResponsePrepared(ctx, s.hooks)

	respBytes, err := proto.Marshal(respContent)
	if err != nil {
		s.writeError(ctx, resp, wrapInternal(err, "failed to marshal proto response"))
		return
	}

	ctx = ctxsetters.WithStatusCode(ctx, http.StatusOK)
	resp.Header().Set("Content-Type", "application/protobuf")
	resp.Header().Set("Content-Length", strconv.Itoa(len(respBytes)))
	resp.WriteHeader(http.StatusOK)
	if n, err := resp.Write(respBytes); err != nil {
		msg := fmt.Sprintf("failed to write response, %d of %d bytes written: %s", n, len(respBytes), err.Error())
		twerr := twirp.NewError(twirp.Unknown, msg)
		ctx = callError(ctx, s.hooks, twerr)
	}
	callResponseSent(ctx, s.hooks)
}

func (s *passwordServiceServer) serveUpdatePassword(ctx context.Context, resp http.ResponseWriter, req *http.Request) {
	header := req.Header.Get("Content-Type")
	i := strings.Index(header, ";")
	if i == -1 {
		i = len(header)
	}
	switch strings.TrimSpace(strings.ToLower(header[:i])) {
	case "application/json":
		s.serveUpdatePasswordJSON(ctx, resp, req)
	case "application/protobuf":
		s.serveUpdatePasswordProtobuf(ctx, resp, req)
	default:
		msg := fmt.Sprintf("unexpected Content-Type: %q", req.Header.Get("Content-Type"))
		twerr := badRouteError(msg, req.Method, req.URL.Path)
		s.writeError(ctx, resp, twerr)
	}
}

func (s *passwordServiceServer) serveUpdatePasswordJSON(ctx context.Context, resp http.ResponseWriter, req *http.Request) {
	var err error
	ctx = ctxsetters.WithMethodName(ctx, "UpdatePassword")
	ctx, err = callRequestRouted(ctx, s.hooks)
	if err != nil {
		s.writeError(ctx, resp, err)
		return
	}

	d := json.NewDecoder(req.Body)
	rawReqBody := json.RawMessage{}
	if err := d.Decode(&rawReqBody); err != nil {
		s.handleRequestBodyError(ctx, resp, "the json request could not be decoded", err)
		return
	}
	reqContent := new(UpdatePasswordRequest)
	unmarshaler := protojson.UnmarshalOptions{DiscardUnknown: true}
	if err = unmarshaler.Unmarshal(rawReqBody, reqContent); err != nil {
		s.handleRequestBodyError(ctx, resp, "the json request could not be decoded", err)
		return
	}

	handler := s.PasswordService.UpdatePassword
	if s.interceptor != nil {
		handler = func(ctx context.Context, req *UpdatePasswordRequest) (*google_protobuf.Empty, error) {
			resp, err := s.interceptor(
				func(ctx context.Context, req interface{}) (interface{}, error) {
					typedReq, ok := req.(*UpdatePasswordRequest)
					if !ok {
						return nil, twirp.InternalError("failed type assertion req.(*UpdatePasswordRequest) when calling interceptor")
					}
					return s.PasswordService.UpdatePassword(ctx, typedReq)
				},
			)(ctx, req)
			if resp != nil {
				typedResp, ok := resp.(*google_protobuf.Empty)
				if !ok {
					return nil, twirp.InternalError("failed type assertion resp.(*google_protobuf.Empty) when calling interceptor")
				}
				return typedResp, err
			}
			return nil, err
		}
	}

	// Call service method
	var respContent *google_protobuf.Empty
	func() {
		defer ensurePanicResponses(ctx, resp, s.hooks)
		respContent, err = handler(ctx, reqContent)
	}()

	if err != nil {
		s.writeError(ctx, resp, err)
		return
	}
	if respContent == nil {
		s.writeError(ctx, resp, twirp.InternalError("received a nil *google_protobuf.Empty and nil error while calling UpdatePassword. nil responses are not supported"))
		return
	}

	ctx = callResponsePrepared(ctx, s.hooks)

	marshaler := &protojson.MarshalOptions{UseProtoNames: !s.jsonCamelCase, EmitUnpopulated: !s.jsonSkipDefaults}
	respBytes, err := marshaler.Marshal(respContent)
	if err != nil {
		s.writeError(ctx, resp, wrapInternal(err, "failed to marshal json response"))
		return
	}

	ctx = ctxsetters.WithStatusCode(ctx, http.StatusOK)
	resp.Header().Set("Content-Type", "application/json")
	resp.Header().Set("Content-Length", strconv.Itoa(len(respBytes)))
	resp.WriteHeader(http.StatusOK)

	if n, err := resp.Write(respBytes); err != nil {
		msg := fmt.Sprintf("failed to write response, %d of %d bytes written: %s", n, len(respBytes), err.Error())
		twerr := twirp.NewError(twirp.Unknown, msg)
		ctx = callError(ctx, s.hooks, twerr)
	}
	callResponseSent(ctx, s.hooks)
}

func (s *passwordServiceServer) serveUpdatePasswordProtobuf(ctx context.Context, resp http.ResponseWriter, req *http.Request) {
	var err error
	ctx = ctxsetters.WithMethodName(ctx, "UpdatePassword")
	ctx, err = callRequestRouted(ctx, s.hooks)
	if err != nil {
		s.writeError(ctx, resp, err)
		return
	}

	buf, err := ioutil.ReadAll(req.Body)
	if err != nil {
		s.handleRequestBodyError(ctx, resp, "failed to read request body", err)
		return
	}
	reqContent := new(UpdatePasswordRequest)
	if err = proto.Unmarshal(buf, reqContent); err != nil {
		s.writeError(ctx, resp, malformedRequestError("the protobuf request could not be decoded"))
		return
	}

	handler := s.PasswordService.UpdatePassword
	if s.interceptor != nil {
		handler = func(ctx context.Context, req *UpdatePasswordRequest) (*google_protobuf.Empty, error) {
			resp, err := s.interceptor(
				func(ctx context.Context, req interface{}) (interface{}, error) {
					typedReq, ok := req.(*UpdatePasswordRequest)
					if !ok {
						return nil, twirp.InternalError("failed type assertion req.(*UpdatePasswordRequest) when calling interceptor")
					}
					return s.PasswordService.UpdatePassword(ctx, typedReq)
				},
			)(ctx, req)
			if resp != nil {
				typedResp, ok := resp.(*google_protobuf.Empty)
				if !ok {
					return nil, twirp.InternalError("failed type assertion resp.(*google_protobuf.Empty) when calling interceptor")
				}
				return typedResp, err
			}
			return nil, err
		}
	}

	// Call service method
	var respContent *google_protobuf.Empty
	func() {
		defer ensurePanicResponses(ctx, resp, s.hooks)
		respContent, err = handler(ctx, reqContent)
	}()

	if err != nil {
		s.writeError(ctx, resp, err)
		return
	}
	if respContent == nil {
		s.writeError(ctx, resp, twirp.InternalError("received a nil *google_protobuf.Empty and nil error while calling UpdatePassword. nil responses are not supported"))
		return
	}

	ctx = callResponsePrepared(ctx, s.hooks)

	respBytes, err := proto.Marshal(respContent)
	if err != nil {
		s.writeError(ctx, resp, wrapInternal(err, "failed to marshal proto response"))
		return
	}

	ctx = ctxsetters.WithStatusCode(ctx, http.StatusOK)
	resp.Header().Set("Content-Type", "application/protobuf")
	resp.Header().Set("Content-Length", strconv.Itoa(len(respBytes)))
	resp.WriteHeader(http.StatusOK)
	if n, err := resp.Write(respBytes); err != nil {
		msg := fmt.Sprintf("failed to write response, %d of %d bytes written: %s", n, len(respBytes), err.Error())
		twerr := twirp.NewError(twirp.Unknown, msg)
		ctx = callError(ctx, s.hooks, twerr)
	}
	callResponseSent(ctx, s.hooks)
}

func (s *passwordServiceServer) ServiceDescriptor() ([]byte, int) {
	return twirpFileDescriptor2, 0
}

func (s *passwordServiceServer) ProtocGenTwirpVersion() string {
	return "v8.1.2"
}

// PathPrefix returns the base service path, in the form: "/<prefix>/<package>.<Service>/"
// that is everything in a Twirp route except for the <Method>. This can be used for routing,
// for example to identify the requests that are targeted to this service in a mux.
func (s *passwordServiceServer) PathPrefix() string {
	return baseServicePath(s.pathPrefix, "users.password", "PasswordService")
}

var twirpFileDescriptor2 = []byte{
	// 341 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x91, 0x4f, 0x4b, 0xf3, 0x40,
	0x10, 0xc6, 0xd9, 0xb7, 0xf4, 0xb5, 0x9d, 0xd6, 0x56, 0xd7, 0xaa, 0x25, 0x5e, 0x34, 0x28, 0x8a,
	0xe0, 0x06, 0xf4, 0x1b, 0x04, 0x14, 0x7a, 0x90, 0x4a, 0x4a, 0x2f, 0x5e, 0x42, 0x9a, 0x8c, 0x25,
	0x98, 0x66, 0xe3, 0xee, 0xa6, 0xc5, 0x9b, 0x9f, 0xc1, 0x4f, 0xe2, 0xe7, 0x2b, 0x1e, 0x24, 0x7f,
	0x56, 0x6d, 0xac, 0xe0, 0x6d, 0x67, 0x9e, 0xd9, 0xdf, 0x0c, 0xcf, 0x03, 0x3b, 0xa9, 0x44, 0x61,
	0x25, 0x9e, 0x94, 0x0b, 0x2e, 0x02, 0x96, 0x08, 0xae, 0x38, 0xed, 0x64, 0x4d, 0xc9, 0x74, 0xd7,
	0x38, 0x98, 0x72, 0x3e, 0x8d, 0xd0, 0xca, 0xd5, 0x49, 0xfa, 0x60, 0xe1, 0x2c, 0x51, 0xcf, 0xc5,
	0xb0, 0xb1, 0x3f, 0xf7, 0xa2, 0x30, 0xf0, 0x14, 0x5a, 0xfa, 0x51, 0x08, 0xa6, 0x0d, 0x3d, 0x07,
	0x25, 0xaa, 0xbb, 0x12, 0xe3, 0xe0, 0x53, 0x8a, 0x52, 0xd1, 0x73, 0xd8, 0xc6, 0x99, 0x17, 0x46,
	0x2e, 0x17, 0x6e, 0xb6, 0x28, 0xf6, 0x66, 0xd8, 0x27, 0x87, 0xe4, 0xac, 0xe9, 0x74, 0x73, 0x61,
	0x28, 0xc6, 0x65, 0xdb, 0xbc, 0x01, 0x3a, 0xfa, 0x49, 0xe8, 0x41, 0x5d, 0xf1, 0x47, 0x8c, 0xcb,
	0x5f, 0x45, 0x41, 0x0d, 0x68, 0xe8, 0x8b, 0xfb, 0xff, 0x72, 0xe1, 0xb3, 0x36, 0x5f, 0x09, 0xec,
	0x8e, 0x93, 0xec, 0xb8, 0x2a, 0xeb, 0x14, 0xc0, 0xf3, 0x7d, 0x9e, 0xc6, 0xca, 0x0d, 0x83, 0x02,
	0x68, 0x37, 0x96, 0x76, 0x5d, 0xd4, 0xde, 0x08, 0x71, 0x9a, 0xa5, 0x36, 0x08, 0xe8, 0x11, 0xb4,
	0x79, 0x14, 0xb8, 0x95, 0x15, 0x2d, 0x1e, 0x05, 0x1a, 0x49, 0x2f, 0xa0, 0x1d, 0xe3, 0xe2, 0x6b,
	0xa4, 0x96, 0xd3, 0x60, 0x69, 0x6f, 0x88, 0xfa, 0x16, 0xf4, 0x5f, 0x88, 0xd3, 0x8a, 0x71, 0xa1,
	0xc7, 0x2f, 0xdf, 0x09, 0x74, 0x75, 0x31, 0x42, 0x31, 0x0f, 0x7d, 0xa4, 0xb7, 0xb0, 0xb9, 0x62,
	0x1a, 0x3d, 0x66, 0xab, 0x61, 0xb0, 0x75, 0x9e, 0x1a, 0x7b, 0xac, 0x88, 0x88, 0xe9, 0x88, 0xd8,
	0x75, 0x16, 0x11, 0x1d, 0x40, 0xeb, 0x9b, 0x7f, 0xd4, 0xac, 0xc2, 0x46, 0x7f, 0x47, 0x0d, 0xa1,
	0xb3, 0xea, 0x20, 0x3d, 0xa9, 0xd2, 0xd6, 0x3a, 0xfc, 0x1b, 0xd0, 0x86, 0xfb, 0x86, 0x48, 0x7c,
	0x2b, 0x63, 0x4c, 0xfe, 0xe7, 0xda, 0xd5, 0x47, 0x00, 0x00, 0x00, 0xff, 0xff, 0x38, 0x37, 0xc0,
	0x61, 0x8f, 0x02, 0x00, 0x00,
}
