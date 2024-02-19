// Code generated by Kitex v0.8.0. DO NOT EDIT.

package lowerservice

import (
	"context"
	"errors"
	"fmt"
	echo "github.com/cloudwego/kitex-tests/thrift_streaming/kitex_gen/echo"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
	streaming "github.com/cloudwego/kitex/pkg/streaming"
)

var errInvalidMessageType = errors.New("invalid message type for service method handler")

var serviceMethods = map[string]kitex.MethodInfo{
	"echo_bidirectional": kitex.NewMethodInfo(
		echoBidirectionalHandler,
		newLowerServiceEchoBidirectionalArgs,
		newLowerServiceEchoBidirectionalResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingBidirectional),
	),
	"echo_client": kitex.NewMethodInfo(
		echoClientHandler,
		newLowerServiceEchoClientArgs,
		newLowerServiceEchoClientResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingClient),
	),
	"echo_server": kitex.NewMethodInfo(
		echoServerHandler,
		newLowerServiceEchoServerArgs,
		newLowerServiceEchoServerResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingServer),
	),
	"echo_unary": kitex.NewMethodInfo(
		echoUnaryHandler,
		newLowerServiceEchoUnaryArgs,
		newLowerServiceEchoUnaryResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingUnary),
	),
	"echo_pingPong": kitex.NewMethodInfo(
		echoPingPongHandler,
		newLowerServiceEchoPingPongArgs,
		newLowerServiceEchoPingPongResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
}

var (
	lowerServiceServiceInfo                = NewServiceInfo()
	lowerServiceServiceInfoForClient       = NewServiceInfoForClient()
	lowerServiceServiceInfoForStreamClient = NewServiceInfoForStreamClient()
)

// for server
func serviceInfo() *kitex.ServiceInfo {
	return lowerServiceServiceInfo
}

// for client
func serviceInfoForStreamClient() *kitex.ServiceInfo {
	return lowerServiceServiceInfoForStreamClient
}

// for stream client
func serviceInfoForClient() *kitex.ServiceInfo {
	return lowerServiceServiceInfoForClient
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
	serviceName := "lower_service"
	handlerType := (*echo.LowerService)(nil)
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
		return errors.New("LowerService.EchoBidirectional is a thrift streaming method, please call with Kitex StreamClient")
	}
	stream := &lowerServiceecho_bidirectionalServer{st.Stream}
	return handler.(echo.LowerService).EchoBidirectional(stream)
}

type lowerServiceecho_bidirectionalClient struct {
	streaming.Stream
}

func (x *lowerServiceecho_bidirectionalClient) Send(m *echo.LowerRequest) error {
	return x.Stream.SendMsg(m)
}
func (x *lowerServiceecho_bidirectionalClient) Recv() (*echo.LowerResponse, error) {
	m := new(echo.LowerResponse)
	return m, x.Stream.RecvMsg(m)
}

type lowerServiceecho_bidirectionalServer struct {
	streaming.Stream
}

func (x *lowerServiceecho_bidirectionalServer) Send(m *echo.LowerResponse) error {
	return x.Stream.SendMsg(m)
}

func (x *lowerServiceecho_bidirectionalServer) Recv() (*echo.LowerRequest, error) {
	m := new(echo.LowerRequest)
	return m, x.Stream.RecvMsg(m)
}

func newLowerServiceEchoBidirectionalArgs() interface{} {
	return echo.NewLowerServiceEchoBidirectionalArgs()
}

func newLowerServiceEchoBidirectionalResult() interface{} {
	return echo.NewLowerServiceEchoBidirectionalResult()
}

func echoClientHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	st, ok := arg.(*streaming.Args)
	if !ok {
		return errors.New("LowerService.EchoClient is a thrift streaming method, please call with Kitex StreamClient")
	}
	stream := &lowerServiceecho_clientServer{st.Stream}
	return handler.(echo.LowerService).EchoClient(stream)
}

type lowerServiceecho_clientClient struct {
	streaming.Stream
}

func (x *lowerServiceecho_clientClient) Send(m *echo.LowerRequest) error {
	return x.Stream.SendMsg(m)
}
func (x *lowerServiceecho_clientClient) CloseAndRecv() (*echo.LowerResponse, error) {
	if err := x.Stream.Close(); err != nil {
		return nil, err
	}
	m := new(echo.LowerResponse)
	return m, x.Stream.RecvMsg(m)
}

type lowerServiceecho_clientServer struct {
	streaming.Stream
}

func (x *lowerServiceecho_clientServer) SendAndClose(m *echo.LowerResponse) error {
	return x.Stream.SendMsg(m)
}

func (x *lowerServiceecho_clientServer) Recv() (*echo.LowerRequest, error) {
	m := new(echo.LowerRequest)
	return m, x.Stream.RecvMsg(m)
}

func newLowerServiceEchoClientArgs() interface{} {
	return echo.NewLowerServiceEchoClientArgs()
}

func newLowerServiceEchoClientResult() interface{} {
	return echo.NewLowerServiceEchoClientResult()
}

func echoServerHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	st, ok := arg.(*streaming.Args)
	if !ok {
		return errors.New("LowerService.EchoServer is a thrift streaming method, please call with Kitex StreamClient")
	}
	stream := &lowerServiceecho_serverServer{st.Stream}
	req := new(echo.LowerRequest)
	if err := st.Stream.RecvMsg(req); err != nil {
		return err
	}
	return handler.(echo.LowerService).EchoServer(req, stream)
}

type lowerServiceecho_serverClient struct {
	streaming.Stream
}

func (x *lowerServiceecho_serverClient) Recv() (*echo.LowerResponse, error) {
	m := new(echo.LowerResponse)
	return m, x.Stream.RecvMsg(m)
}

type lowerServiceecho_serverServer struct {
	streaming.Stream
}

func (x *lowerServiceecho_serverServer) Send(m *echo.LowerResponse) error {
	return x.Stream.SendMsg(m)
}

func newLowerServiceEchoServerArgs() interface{} {
	return echo.NewLowerServiceEchoServerArgs()
}

func newLowerServiceEchoServerResult() interface{} {
	return echo.NewLowerServiceEchoServerResult()
}

func echoUnaryHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	if streaming.GetStream(ctx) == nil {
		return errors.New("LowerService.EchoUnary is a thrift streaming unary method, please call with Kitex StreamClient or remove the annotation streaming.mode")
	}
	realArg := arg.(*echo.LowerServiceEchoUnaryArgs)
	realResult := result.(*echo.LowerServiceEchoUnaryResult)
	success, err := handler.(echo.LowerService).EchoUnary(ctx, realArg.Req1)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newLowerServiceEchoUnaryArgs() interface{} {
	return echo.NewLowerServiceEchoUnaryArgs()
}

func newLowerServiceEchoUnaryResult() interface{} {
	return echo.NewLowerServiceEchoUnaryResult()
}

func echoPingPongHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*echo.LowerServiceEchoPingPongArgs)
	realResult := result.(*echo.LowerServiceEchoPingPongResult)
	success, err := handler.(echo.LowerService).EchoPingPong(ctx, realArg.Req1)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newLowerServiceEchoPingPongArgs() interface{} {
	return echo.NewLowerServiceEchoPingPongArgs()
}

func newLowerServiceEchoPingPongResult() interface{} {
	return echo.NewLowerServiceEchoPingPongResult()
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) EchoBidirectional(ctx context.Context) (LowerService_echo_bidirectionalClient, error) {
	streamClient, ok := p.c.(client.Streaming)
	if !ok {
		return nil, fmt.Errorf("client not support streaming")
	}
	res := new(streaming.Result)
	err := streamClient.Stream(ctx, "echo_bidirectional", nil, res)
	if err != nil {
		return nil, err
	}
	stream := &lowerServiceecho_bidirectionalClient{res.Stream}
	return stream, nil
}

func (p *kClient) EchoClient(ctx context.Context) (LowerService_echo_clientClient, error) {
	streamClient, ok := p.c.(client.Streaming)
	if !ok {
		return nil, fmt.Errorf("client not support streaming")
	}
	res := new(streaming.Result)
	err := streamClient.Stream(ctx, "echo_client", nil, res)
	if err != nil {
		return nil, err
	}
	stream := &lowerServiceecho_clientClient{res.Stream}
	return stream, nil
}

func (p *kClient) EchoServer(ctx context.Context, req1 *echo.LowerRequest) (LowerService_echo_serverClient, error) {
	streamClient, ok := p.c.(client.Streaming)
	if !ok {
		return nil, fmt.Errorf("client not support streaming")
	}
	res := new(streaming.Result)
	err := streamClient.Stream(ctx, "echo_server", nil, res)
	if err != nil {
		return nil, err
	}
	stream := &lowerServiceecho_serverClient{res.Stream}

	if err := stream.Stream.SendMsg(req1); err != nil {
		return nil, err
	}
	if err := stream.Stream.Close(); err != nil {
		return nil, err
	}
	return stream, nil
}

func (p *kClient) EchoUnary(ctx context.Context, req1 *echo.LowerRequest) (r *echo.LowerResponse, err error) {
	var _args echo.LowerServiceEchoUnaryArgs
	_args.Req1 = req1
	var _result echo.LowerServiceEchoUnaryResult
	if err = p.c.Call(ctx, "echo_unary", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) EchoPingPong(ctx context.Context, req1 *echo.LowerRequest) (r *echo.LowerResponse, err error) {
	var _args echo.LowerServiceEchoPingPongArgs
	_args.Req1 = req1
	var _result echo.LowerServiceEchoPingPongResult
	if err = p.c.Call(ctx, "echo_pingPong", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
