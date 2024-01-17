// Code generated by Kitex v0.8.0. DO NOT EDIT.

package echoservice

import (
	"context"
	echo "github.com/cloudwego/kitex-tests/thrift_streaming/kitex_gen/echo"
	client "github.com/cloudwego/kitex/client"
	callopt "github.com/cloudwego/kitex/client/callopt"
	"github.com/cloudwego/kitex/client/callopt/streamcall"
	"github.com/cloudwego/kitex/client/streamclient"
	streaming "github.com/cloudwego/kitex/pkg/streaming"
	transport "github.com/cloudwego/kitex/transport"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	EchoPingPong(ctx context.Context, req1 *echo.EchoRequest, req2 *echo.EchoRequest, callOptions ...callopt.Option) (r *echo.EchoResponse, err error)
	EchoOneway(ctx context.Context, req1 *echo.EchoRequest, callOptions ...callopt.Option) (err error)
	Ping(ctx context.Context, callOptions ...callopt.Option) (err error)
}

// StreamClient is designed to provide Interface for Streaming APIs.
type StreamClient interface {
	EchoBidirectional(ctx context.Context, callOptions ...streamcall.Option) (stream EchoService_EchoBidirectionalClient, err error)
	EchoClient(ctx context.Context, callOptions ...streamcall.Option) (stream EchoService_EchoClientClient, err error)
	EchoServer(ctx context.Context, req1 *echo.EchoRequest, callOptions ...streamcall.Option) (stream EchoService_EchoServerClient, err error)
	EchoUnary(ctx context.Context, req1 *echo.EchoRequest, callOptions ...streamcall.Option) (r *echo.EchoResponse, err error)
}

type EchoService_EchoBidirectionalClient interface {
	streaming.Stream
	Send(*echo.EchoRequest) error
	Recv() (*echo.EchoResponse, error)
}

type EchoService_EchoClientClient interface {
	streaming.Stream
	Send(*echo.EchoRequest) error
	CloseAndRecv() (*echo.EchoResponse, error)
}

type EchoService_EchoServerClient interface {
	streaming.Stream
	Recv() (*echo.EchoResponse, error)
}

// NewClient creates a client for the service defined in IDL.
func NewClient(destService string, opts ...client.Option) (Client, error) {
	var options []client.Option
	options = append(options, client.WithDestService(destService))

	options = append(options, opts...)

	kc, err := client.NewClient(serviceInfoForClient(), options...)
	if err != nil {
		return nil, err
	}
	return &kEchoServiceClient{
		kClient: newServiceClient(kc),
	}, nil
}

// MustNewClient creates a client for the service defined in IDL. It panics if any error occurs.
func MustNewClient(destService string, opts ...client.Option) Client {
	kc, err := NewClient(destService, opts...)
	if err != nil {
		panic(err)
	}
	return kc
}

type kEchoServiceClient struct {
	*kClient
}

func (p *kEchoServiceClient) EchoPingPong(ctx context.Context, req1 *echo.EchoRequest, req2 *echo.EchoRequest, callOptions ...callopt.Option) (r *echo.EchoResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.EchoPingPong(ctx, req1, req2)
}

func (p *kEchoServiceClient) EchoOneway(ctx context.Context, req1 *echo.EchoRequest, callOptions ...callopt.Option) (err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.EchoOneway(ctx, req1)
}

func (p *kEchoServiceClient) Ping(ctx context.Context, callOptions ...callopt.Option) (err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.Ping(ctx)
}

// NewStreamClient creates a stream client for the service's streaming APIs defined in IDL.
func NewStreamClient(destService string, opts ...streamclient.Option) (StreamClient, error) {
	var options []client.Option
	options = append(options, client.WithDestService(destService))
	options = append(options, client.WithTransportProtocol(transport.GRPC))
	options = append(options, streamclient.GetClientOptions(opts)...)

	kc, err := client.NewClient(serviceInfoForStreamClient(), options...)
	if err != nil {
		return nil, err
	}
	return &kEchoServiceStreamClient{
		kClient: newServiceClient(kc),
	}, nil
}

// MustNewStreamClient creates a stream client for the service's streaming APIs defined in IDL.
// It panics if any error occurs.
func MustNewStreamClient(destService string, opts ...streamclient.Option) StreamClient {
	kc, err := NewStreamClient(destService, opts...)
	if err != nil {
		panic(err)
	}
	return kc
}

type kEchoServiceStreamClient struct {
	*kClient
}

func (p *kEchoServiceStreamClient) EchoBidirectional(ctx context.Context, callOptions ...streamcall.Option) (stream EchoService_EchoBidirectionalClient, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, streamcall.GetCallOptions(callOptions))
	return p.kClient.EchoBidirectional(ctx)
}

func (p *kEchoServiceStreamClient) EchoClient(ctx context.Context, callOptions ...streamcall.Option) (stream EchoService_EchoClientClient, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, streamcall.GetCallOptions(callOptions))
	return p.kClient.EchoClient(ctx)
}

func (p *kEchoServiceStreamClient) EchoServer(ctx context.Context, req1 *echo.EchoRequest, callOptions ...streamcall.Option) (stream EchoService_EchoServerClient, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, streamcall.GetCallOptions(callOptions))
	return p.kClient.EchoServer(ctx, req1)
}

func (p *kEchoServiceStreamClient) EchoUnary(ctx context.Context, req1 *echo.EchoRequest, callOptions ...streamcall.Option) (r *echo.EchoResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, streamcall.GetCallOptions(callOptions))
	return p.kClient.EchoUnary(ctx, req1)
}
