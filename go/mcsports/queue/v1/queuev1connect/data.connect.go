// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: mcsports/queue/v1/data.proto

package queuev1connect

import (
	connect "connectrpc.com/connect"
	context "context"
	errors "errors"
	v1 "github.com/bufbuild/buf-tour/gen/mcsports/queue/v1"
	http "net/http"
	strings "strings"
)

// This is a compile-time assertion to ensure that this generated file and the connect package are
// compatible. If you get a compiler error that this constant is not defined, this code was
// generated with a version of connect newer than the one compiled into your binary. You can fix the
// problem by either regenerating this code with an older version of connect or updating the connect
// version compiled into your binary.
const _ = connect.IsAtLeastVersion1_13_0

const (
	// QueueDataName is the fully-qualified name of the QueueData service.
	QueueDataName = "mcsports.queue.v1.QueueData"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// QueueDataGetQueueProcedure is the fully-qualified name of the QueueData's GetQueue RPC.
	QueueDataGetQueueProcedure = "/mcsports.queue.v1.QueueData/GetQueue"
	// QueueDataGetAllQueuesProcedure is the fully-qualified name of the QueueData's GetAllQueues RPC.
	QueueDataGetAllQueuesProcedure = "/mcsports.queue.v1.QueueData/GetAllQueues"
	// QueueDataGetQueueByPlayerProcedure is the fully-qualified name of the QueueData's
	// GetQueueByPlayer RPC.
	QueueDataGetQueueByPlayerProcedure = "/mcsports.queue.v1.QueueData/GetQueueByPlayer"
	// QueueDataDeleteQueueProcedure is the fully-qualified name of the QueueData's DeleteQueue RPC.
	QueueDataDeleteQueueProcedure = "/mcsports.queue.v1.QueueData/DeleteQueue"
	// QueueDataUpdateQueueProcedure is the fully-qualified name of the QueueData's UpdateQueue RPC.
	QueueDataUpdateQueueProcedure = "/mcsports.queue.v1.QueueData/UpdateQueue"
	// QueueDataGetAllQueueTypesProcedure is the fully-qualified name of the QueueData's
	// GetAllQueueTypes RPC.
	QueueDataGetAllQueueTypesProcedure = "/mcsports.queue.v1.QueueData/GetAllQueueTypes"
	// QueueDataGetQueueTypePlayerInformationProcedure is the fully-qualified name of the QueueData's
	// GetQueueTypePlayerInformation RPC.
	QueueDataGetQueueTypePlayerInformationProcedure = "/mcsports.queue.v1.QueueData/GetQueueTypePlayerInformation"
)

// QueueDataClient is a client for the mcsports.queue.v1.QueueData service.
type QueueDataClient interface {
	GetQueue(context.Context, *connect.Request[v1.GetQueueRequest]) (*connect.Response[v1.GetQueueResponse], error)
	GetAllQueues(context.Context, *connect.Request[v1.GetAllQueuesRequest]) (*connect.Response[v1.GetAllQueuesResponse], error)
	GetQueueByPlayer(context.Context, *connect.Request[v1.GetQueueByPlayerRequest]) (*connect.Response[v1.GetQueueByPlayerResponse], error)
	DeleteQueue(context.Context, *connect.Request[v1.DeleteQueueRequest]) (*connect.Response[v1.DeleteQueueResponse], error)
	UpdateQueue(context.Context, *connect.Request[v1.UpdateQueueRequest]) (*connect.Response[v1.UpdateQueueResponse], error)
	GetAllQueueTypes(context.Context, *connect.Request[v1.GetAllQueueTypesRequest]) (*connect.Response[v1.GetAllQueueTypesResponse], error)
	GetQueueTypePlayerInformation(context.Context, *connect.Request[v1.GetQueueTypePlayerInformationRequest]) (*connect.Response[v1.GetQueueTypePlayerInformationResponse], error)
}

// NewQueueDataClient constructs a client for the mcsports.queue.v1.QueueData service. By default,
// it uses the Connect protocol with the binary Protobuf Codec, asks for gzipped responses, and
// sends uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the connect.WithGRPC()
// or connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewQueueDataClient(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) QueueDataClient {
	baseURL = strings.TrimRight(baseURL, "/")
	queueDataMethods := v1.File_mcsports_queue_v1_data_proto.Services().ByName("QueueData").Methods()
	return &queueDataClient{
		getQueue: connect.NewClient[v1.GetQueueRequest, v1.GetQueueResponse](
			httpClient,
			baseURL+QueueDataGetQueueProcedure,
			connect.WithSchema(queueDataMethods.ByName("GetQueue")),
			connect.WithClientOptions(opts...),
		),
		getAllQueues: connect.NewClient[v1.GetAllQueuesRequest, v1.GetAllQueuesResponse](
			httpClient,
			baseURL+QueueDataGetAllQueuesProcedure,
			connect.WithSchema(queueDataMethods.ByName("GetAllQueues")),
			connect.WithClientOptions(opts...),
		),
		getQueueByPlayer: connect.NewClient[v1.GetQueueByPlayerRequest, v1.GetQueueByPlayerResponse](
			httpClient,
			baseURL+QueueDataGetQueueByPlayerProcedure,
			connect.WithSchema(queueDataMethods.ByName("GetQueueByPlayer")),
			connect.WithClientOptions(opts...),
		),
		deleteQueue: connect.NewClient[v1.DeleteQueueRequest, v1.DeleteQueueResponse](
			httpClient,
			baseURL+QueueDataDeleteQueueProcedure,
			connect.WithSchema(queueDataMethods.ByName("DeleteQueue")),
			connect.WithClientOptions(opts...),
		),
		updateQueue: connect.NewClient[v1.UpdateQueueRequest, v1.UpdateQueueResponse](
			httpClient,
			baseURL+QueueDataUpdateQueueProcedure,
			connect.WithSchema(queueDataMethods.ByName("UpdateQueue")),
			connect.WithClientOptions(opts...),
		),
		getAllQueueTypes: connect.NewClient[v1.GetAllQueueTypesRequest, v1.GetAllQueueTypesResponse](
			httpClient,
			baseURL+QueueDataGetAllQueueTypesProcedure,
			connect.WithSchema(queueDataMethods.ByName("GetAllQueueTypes")),
			connect.WithClientOptions(opts...),
		),
		getQueueTypePlayerInformation: connect.NewClient[v1.GetQueueTypePlayerInformationRequest, v1.GetQueueTypePlayerInformationResponse](
			httpClient,
			baseURL+QueueDataGetQueueTypePlayerInformationProcedure,
			connect.WithSchema(queueDataMethods.ByName("GetQueueTypePlayerInformation")),
			connect.WithClientOptions(opts...),
		),
	}
}

// queueDataClient implements QueueDataClient.
type queueDataClient struct {
	getQueue                      *connect.Client[v1.GetQueueRequest, v1.GetQueueResponse]
	getAllQueues                  *connect.Client[v1.GetAllQueuesRequest, v1.GetAllQueuesResponse]
	getQueueByPlayer              *connect.Client[v1.GetQueueByPlayerRequest, v1.GetQueueByPlayerResponse]
	deleteQueue                   *connect.Client[v1.DeleteQueueRequest, v1.DeleteQueueResponse]
	updateQueue                   *connect.Client[v1.UpdateQueueRequest, v1.UpdateQueueResponse]
	getAllQueueTypes              *connect.Client[v1.GetAllQueueTypesRequest, v1.GetAllQueueTypesResponse]
	getQueueTypePlayerInformation *connect.Client[v1.GetQueueTypePlayerInformationRequest, v1.GetQueueTypePlayerInformationResponse]
}

// GetQueue calls mcsports.queue.v1.QueueData.GetQueue.
func (c *queueDataClient) GetQueue(ctx context.Context, req *connect.Request[v1.GetQueueRequest]) (*connect.Response[v1.GetQueueResponse], error) {
	return c.getQueue.CallUnary(ctx, req)
}

// GetAllQueues calls mcsports.queue.v1.QueueData.GetAllQueues.
func (c *queueDataClient) GetAllQueues(ctx context.Context, req *connect.Request[v1.GetAllQueuesRequest]) (*connect.Response[v1.GetAllQueuesResponse], error) {
	return c.getAllQueues.CallUnary(ctx, req)
}

// GetQueueByPlayer calls mcsports.queue.v1.QueueData.GetQueueByPlayer.
func (c *queueDataClient) GetQueueByPlayer(ctx context.Context, req *connect.Request[v1.GetQueueByPlayerRequest]) (*connect.Response[v1.GetQueueByPlayerResponse], error) {
	return c.getQueueByPlayer.CallUnary(ctx, req)
}

// DeleteQueue calls mcsports.queue.v1.QueueData.DeleteQueue.
func (c *queueDataClient) DeleteQueue(ctx context.Context, req *connect.Request[v1.DeleteQueueRequest]) (*connect.Response[v1.DeleteQueueResponse], error) {
	return c.deleteQueue.CallUnary(ctx, req)
}

// UpdateQueue calls mcsports.queue.v1.QueueData.UpdateQueue.
func (c *queueDataClient) UpdateQueue(ctx context.Context, req *connect.Request[v1.UpdateQueueRequest]) (*connect.Response[v1.UpdateQueueResponse], error) {
	return c.updateQueue.CallUnary(ctx, req)
}

// GetAllQueueTypes calls mcsports.queue.v1.QueueData.GetAllQueueTypes.
func (c *queueDataClient) GetAllQueueTypes(ctx context.Context, req *connect.Request[v1.GetAllQueueTypesRequest]) (*connect.Response[v1.GetAllQueueTypesResponse], error) {
	return c.getAllQueueTypes.CallUnary(ctx, req)
}

// GetQueueTypePlayerInformation calls mcsports.queue.v1.QueueData.GetQueueTypePlayerInformation.
func (c *queueDataClient) GetQueueTypePlayerInformation(ctx context.Context, req *connect.Request[v1.GetQueueTypePlayerInformationRequest]) (*connect.Response[v1.GetQueueTypePlayerInformationResponse], error) {
	return c.getQueueTypePlayerInformation.CallUnary(ctx, req)
}

// QueueDataHandler is an implementation of the mcsports.queue.v1.QueueData service.
type QueueDataHandler interface {
	GetQueue(context.Context, *connect.Request[v1.GetQueueRequest]) (*connect.Response[v1.GetQueueResponse], error)
	GetAllQueues(context.Context, *connect.Request[v1.GetAllQueuesRequest]) (*connect.Response[v1.GetAllQueuesResponse], error)
	GetQueueByPlayer(context.Context, *connect.Request[v1.GetQueueByPlayerRequest]) (*connect.Response[v1.GetQueueByPlayerResponse], error)
	DeleteQueue(context.Context, *connect.Request[v1.DeleteQueueRequest]) (*connect.Response[v1.DeleteQueueResponse], error)
	UpdateQueue(context.Context, *connect.Request[v1.UpdateQueueRequest]) (*connect.Response[v1.UpdateQueueResponse], error)
	GetAllQueueTypes(context.Context, *connect.Request[v1.GetAllQueueTypesRequest]) (*connect.Response[v1.GetAllQueueTypesResponse], error)
	GetQueueTypePlayerInformation(context.Context, *connect.Request[v1.GetQueueTypePlayerInformationRequest]) (*connect.Response[v1.GetQueueTypePlayerInformationResponse], error)
}

// NewQueueDataHandler builds an HTTP handler from the service implementation. It returns the path
// on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewQueueDataHandler(svc QueueDataHandler, opts ...connect.HandlerOption) (string, http.Handler) {
	queueDataMethods := v1.File_mcsports_queue_v1_data_proto.Services().ByName("QueueData").Methods()
	queueDataGetQueueHandler := connect.NewUnaryHandler(
		QueueDataGetQueueProcedure,
		svc.GetQueue,
		connect.WithSchema(queueDataMethods.ByName("GetQueue")),
		connect.WithHandlerOptions(opts...),
	)
	queueDataGetAllQueuesHandler := connect.NewUnaryHandler(
		QueueDataGetAllQueuesProcedure,
		svc.GetAllQueues,
		connect.WithSchema(queueDataMethods.ByName("GetAllQueues")),
		connect.WithHandlerOptions(opts...),
	)
	queueDataGetQueueByPlayerHandler := connect.NewUnaryHandler(
		QueueDataGetQueueByPlayerProcedure,
		svc.GetQueueByPlayer,
		connect.WithSchema(queueDataMethods.ByName("GetQueueByPlayer")),
		connect.WithHandlerOptions(opts...),
	)
	queueDataDeleteQueueHandler := connect.NewUnaryHandler(
		QueueDataDeleteQueueProcedure,
		svc.DeleteQueue,
		connect.WithSchema(queueDataMethods.ByName("DeleteQueue")),
		connect.WithHandlerOptions(opts...),
	)
	queueDataUpdateQueueHandler := connect.NewUnaryHandler(
		QueueDataUpdateQueueProcedure,
		svc.UpdateQueue,
		connect.WithSchema(queueDataMethods.ByName("UpdateQueue")),
		connect.WithHandlerOptions(opts...),
	)
	queueDataGetAllQueueTypesHandler := connect.NewUnaryHandler(
		QueueDataGetAllQueueTypesProcedure,
		svc.GetAllQueueTypes,
		connect.WithSchema(queueDataMethods.ByName("GetAllQueueTypes")),
		connect.WithHandlerOptions(opts...),
	)
	queueDataGetQueueTypePlayerInformationHandler := connect.NewUnaryHandler(
		QueueDataGetQueueTypePlayerInformationProcedure,
		svc.GetQueueTypePlayerInformation,
		connect.WithSchema(queueDataMethods.ByName("GetQueueTypePlayerInformation")),
		connect.WithHandlerOptions(opts...),
	)
	return "/mcsports.queue.v1.QueueData/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case QueueDataGetQueueProcedure:
			queueDataGetQueueHandler.ServeHTTP(w, r)
		case QueueDataGetAllQueuesProcedure:
			queueDataGetAllQueuesHandler.ServeHTTP(w, r)
		case QueueDataGetQueueByPlayerProcedure:
			queueDataGetQueueByPlayerHandler.ServeHTTP(w, r)
		case QueueDataDeleteQueueProcedure:
			queueDataDeleteQueueHandler.ServeHTTP(w, r)
		case QueueDataUpdateQueueProcedure:
			queueDataUpdateQueueHandler.ServeHTTP(w, r)
		case QueueDataGetAllQueueTypesProcedure:
			queueDataGetAllQueueTypesHandler.ServeHTTP(w, r)
		case QueueDataGetQueueTypePlayerInformationProcedure:
			queueDataGetQueueTypePlayerInformationHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedQueueDataHandler returns CodeUnimplemented from all methods.
type UnimplementedQueueDataHandler struct{}

func (UnimplementedQueueDataHandler) GetQueue(context.Context, *connect.Request[v1.GetQueueRequest]) (*connect.Response[v1.GetQueueResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("mcsports.queue.v1.QueueData.GetQueue is not implemented"))
}

func (UnimplementedQueueDataHandler) GetAllQueues(context.Context, *connect.Request[v1.GetAllQueuesRequest]) (*connect.Response[v1.GetAllQueuesResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("mcsports.queue.v1.QueueData.GetAllQueues is not implemented"))
}

func (UnimplementedQueueDataHandler) GetQueueByPlayer(context.Context, *connect.Request[v1.GetQueueByPlayerRequest]) (*connect.Response[v1.GetQueueByPlayerResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("mcsports.queue.v1.QueueData.GetQueueByPlayer is not implemented"))
}

func (UnimplementedQueueDataHandler) DeleteQueue(context.Context, *connect.Request[v1.DeleteQueueRequest]) (*connect.Response[v1.DeleteQueueResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("mcsports.queue.v1.QueueData.DeleteQueue is not implemented"))
}

func (UnimplementedQueueDataHandler) UpdateQueue(context.Context, *connect.Request[v1.UpdateQueueRequest]) (*connect.Response[v1.UpdateQueueResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("mcsports.queue.v1.QueueData.UpdateQueue is not implemented"))
}

func (UnimplementedQueueDataHandler) GetAllQueueTypes(context.Context, *connect.Request[v1.GetAllQueueTypesRequest]) (*connect.Response[v1.GetAllQueueTypesResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("mcsports.queue.v1.QueueData.GetAllQueueTypes is not implemented"))
}

func (UnimplementedQueueDataHandler) GetQueueTypePlayerInformation(context.Context, *connect.Request[v1.GetQueueTypePlayerInformationRequest]) (*connect.Response[v1.GetQueueTypePlayerInformationResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("mcsports.queue.v1.QueueData.GetQueueTypePlayerInformation is not implemented"))
}
