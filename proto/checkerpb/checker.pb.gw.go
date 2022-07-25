// Code generated by protoc-gen-grpc-gateway. DO NOT EDIT.
// source: checker.proto

/*
Package checkerpb is a reverse proxy.

It translates gRPC into RESTful JSON APIs.
*/
package checkerpb

import (
	"context"
	"io"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/grpc-ecosystem/grpc-gateway/v2/utilities"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
)

// Suppress "imported and not used" errors
var _ codes.Code
var _ io.Reader
var _ status.Status
var _ = runtime.String
var _ = utilities.NewDoubleArray
var _ = metadata.Join

func request_CheckerService_CheckPermission_0(ctx context.Context, marshaler runtime.Marshaler, client CheckerServiceClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq CheckPermissionReq
	var metadata runtime.ServerMetadata

	newReader, berr := utilities.IOReaderFactory(req.Body)
	if berr != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", berr)
	}
	if err := marshaler.NewDecoder(newReader()).Decode(&protoReq); err != nil && err != io.EOF {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	msg, err := client.CheckPermission(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

func local_request_CheckerService_CheckPermission_0(ctx context.Context, marshaler runtime.Marshaler, server CheckerServiceServer, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq CheckPermissionReq
	var metadata runtime.ServerMetadata

	newReader, berr := utilities.IOReaderFactory(req.Body)
	if berr != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", berr)
	}
	if err := marshaler.NewDecoder(newReader()).Decode(&protoReq); err != nil && err != io.EOF {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	msg, err := server.CheckPermission(ctx, &protoReq)
	return msg, metadata, err

}

// RegisterCheckerServiceHandlerServer registers the http handlers for service CheckerService to "mux".
// UnaryRPC     :call CheckerServiceServer directly.
// StreamingRPC :currently unsupported pending https://github.com/grpc/grpc-go/issues/906.
// Note that using this registration option will cause many gRPC library features to stop working. Consider using RegisterCheckerServiceHandlerFromEndpoint instead.
func RegisterCheckerServiceHandlerServer(ctx context.Context, mux *runtime.ServeMux, server CheckerServiceServer) error {

	mux.Handle("POST", pattern_CheckerService_CheckPermission_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		var stream runtime.ServerTransportStream
		ctx = grpc.NewContextWithServerTransportStream(ctx, &stream)
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		var err error
		ctx, err = runtime.AnnotateIncomingContext(ctx, mux, req, "/checkerpb.CheckerService/CheckPermission", runtime.WithHTTPPathPattern("/checkerpb.CheckerService/CheckPermission"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := local_request_CheckerService_CheckPermission_0(ctx, inboundMarshaler, server, req, pathParams)
		md.HeaderMD, md.TrailerMD = metadata.Join(md.HeaderMD, stream.Header()), metadata.Join(md.TrailerMD, stream.Trailer())
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_CheckerService_CheckPermission_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	return nil
}

// RegisterCheckerServiceHandlerFromEndpoint is same as RegisterCheckerServiceHandler but
// automatically dials to "endpoint" and closes the connection when "ctx" gets done.
func RegisterCheckerServiceHandlerFromEndpoint(ctx context.Context, mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) (err error) {
	conn, err := grpc.Dial(endpoint, opts...)
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			if cerr := conn.Close(); cerr != nil {
				grpclog.Infof("Failed to close conn to %s: %v", endpoint, cerr)
			}
			return
		}
		go func() {
			<-ctx.Done()
			if cerr := conn.Close(); cerr != nil {
				grpclog.Infof("Failed to close conn to %s: %v", endpoint, cerr)
			}
		}()
	}()

	return RegisterCheckerServiceHandler(ctx, mux, conn)
}

// RegisterCheckerServiceHandler registers the http handlers for service CheckerService to "mux".
// The handlers forward requests to the grpc endpoint over "conn".
func RegisterCheckerServiceHandler(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error {
	return RegisterCheckerServiceHandlerClient(ctx, mux, NewCheckerServiceClient(conn))
}

// RegisterCheckerServiceHandlerClient registers the http handlers for service CheckerService
// to "mux". The handlers forward requests to the grpc endpoint over the given implementation of "CheckerServiceClient".
// Note: the gRPC framework executes interceptors within the gRPC handler. If the passed in "CheckerServiceClient"
// doesn't go through the normal gRPC flow (creating a gRPC client etc.) then it will be up to the passed in
// "CheckerServiceClient" to call the correct interceptors.
func RegisterCheckerServiceHandlerClient(ctx context.Context, mux *runtime.ServeMux, client CheckerServiceClient) error {

	mux.Handle("POST", pattern_CheckerService_CheckPermission_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		var err error
		ctx, err = runtime.AnnotateContext(ctx, mux, req, "/checkerpb.CheckerService/CheckPermission", runtime.WithHTTPPathPattern("/checkerpb.CheckerService/CheckPermission"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_CheckerService_CheckPermission_0(ctx, inboundMarshaler, client, req, pathParams)
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_CheckerService_CheckPermission_0(ctx, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	return nil
}

var (
	pattern_CheckerService_CheckPermission_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1}, []string{"checkerpb.CheckerService", "CheckPermission"}, ""))
)

var (
	forward_CheckerService_CheckPermission_0 = runtime.ForwardResponseMessage
)