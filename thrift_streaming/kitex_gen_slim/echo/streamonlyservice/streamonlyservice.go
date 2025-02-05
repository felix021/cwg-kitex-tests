// Code generated by Kitex v0.8.0. DO NOT EDIT.

package streamonlyservice

import (
	"context"
	"errors"
	"fmt"
	echo "github.com/cloudwego/kitex-tests/thrift_streaming/kitex_gen_slim/echo"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
	streaming "github.com/cloudwego/kitex/pkg/streaming"
)

var errInvalidMessageType = errors.New("invalid message type for service method handler")

var serviceMethods = map[string]kitex.MethodInfo{
	"EchoBidirectional": kitex.NewMethodInfo(
		echoBidirectionalHandler,
		newStreamOnlyServiceEchoBidirectionalArgs,
		newStreamOnlyServiceEchoBidirectionalResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingBidirectional),
	),
}

var (
	streamOnlyServiceServiceInfo                = NewServiceInfo()
	streamOnlyServiceServiceInfoForClient       = NewServiceInfoForClient()
	streamOnlyServiceServiceInfoForStreamClient = NewServiceInfoForStreamClient()
)

// for server
func serviceInfo() *kitex.ServiceInfo {
	return streamOnlyServiceServiceInfo
}

// for client
func serviceInfoForStreamClient() *kitex.ServiceInfo {
	return streamOnlyServiceServiceInfoForStreamClient
}

// for stream client
func serviceInfoForClient() *kitex.ServiceInfo {
	return streamOnlyServiceServiceInfoForClient
}

// NewServiceInfo creates a new ServiceInfo containing all methods
func NewServiceInfo() *kitex.ServiceInfo {
	return newServiceInfo(true, true, true)
}

// NewServiceInfo creates a new ServiceInfo containing non-streaming methods
func NewServiceInfoForClient() *kitex.ServiceInfo {
	return newServiceInfo(false, false, true)
}
func NewServiceInfoForStreamClient() *kitex.ServiceInfo {
	return newServiceInfo(true, true, false)
}

func newServiceInfo(hasStreaming bool, keepStreamingMethods bool, keepNonStreamingMethods bool) *kitex.ServiceInfo {
	serviceName := "StreamOnlyService"
	handlerType := (*echo.StreamOnlyService)(nil)
	methods := map[string]kitex.MethodInfo{}
	for name, m := range serviceMethods {
		if m.IsStreaming() && !keepStreamingMethods {
			continue
		}
		if !m.IsStreaming() && !keepNonStreamingMethods {
			continue
		}
		methods[name] = m
	}
	extra := map[string]interface{}{
		"PackageName": "echo",
	}
	if hasStreaming {
		extra["streaming"] = hasStreaming
	}
	svcInfo := &kitex.ServiceInfo{
		ServiceName:     serviceName,
		HandlerType:     handlerType,
		Methods:         methods,
		PayloadCodec:    kitex.Thrift,
		KiteXGenVersion: "v0.8.0",
		Extra:           extra,
	}
	return svcInfo
}

func echoBidirectionalHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	st, ok := arg.(*streaming.Args)
	if !ok {
		return errors.New("StreamOnlyService.EchoBidirectional is a thrift streaming method, please call with Kitex StreamClient")
	}
	stream := &streamOnlyServiceEchoBidirectionalServer{st.Stream}
	return handler.(echo.StreamOnlyService).EchoBidirectional(stream)
}

type streamOnlyServiceEchoBidirectionalClient struct {
	streaming.Stream
}

func (x *streamOnlyServiceEchoBidirectionalClient) Send(m *echo.EchoRequest) error {
	return x.Stream.SendMsg(m)
}
func (x *streamOnlyServiceEchoBidirectionalClient) Recv() (*echo.EchoResponse, error) {
	m := new(echo.EchoResponse)
	return m, x.Stream.RecvMsg(m)
}

type streamOnlyServiceEchoBidirectionalServer struct {
	streaming.Stream
}

func (x *streamOnlyServiceEchoBidirectionalServer) Send(m *echo.EchoResponse) error {
	return x.Stream.SendMsg(m)
}

func (x *streamOnlyServiceEchoBidirectionalServer) Recv() (*echo.EchoRequest, error) {
	m := new(echo.EchoRequest)
	return m, x.Stream.RecvMsg(m)
}

func newStreamOnlyServiceEchoBidirectionalArgs() interface{} {
	return echo.NewStreamOnlyServiceEchoBidirectionalArgs()
}

func newStreamOnlyServiceEchoBidirectionalResult() interface{} {
	return echo.NewStreamOnlyServiceEchoBidirectionalResult()
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) EchoBidirectional(ctx context.Context) (StreamOnlyService_EchoBidirectionalClient, error) {
	streamClient, ok := p.c.(client.Streaming)
	if !ok {
		return nil, fmt.Errorf("client not support streaming")
	}
	res := new(streaming.Result)
	err := streamClient.Stream(ctx, "EchoBidirectional", nil, res)
	if err != nil {
		return nil, err
	}
	stream := &streamOnlyServiceEchoBidirectionalClient{res.Stream}
	return stream, nil
}
