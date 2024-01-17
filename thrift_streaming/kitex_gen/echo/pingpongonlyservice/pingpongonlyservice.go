// Code generated by Kitex v0.8.0. DO NOT EDIT.

package pingpongonlyservice

import (
	"context"
	"errors"
	echo "github.com/cloudwego/kitex-tests/thrift_streaming/kitex_gen/echo"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
)

var errInvalidMessageType = errors.New("invalid message type for service method handler")

var serviceMethods = map[string]kitex.MethodInfo{
	"EchoPingPong": kitex.NewMethodInfo(
		echoPingPongHandler,
		newPingPongOnlyServiceEchoPingPongArgs,
		newPingPongOnlyServiceEchoPingPongResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
}

var (
	pingPongOnlyServiceServiceInfo                = NewServiceInfo()
	pingPongOnlyServiceServiceInfoForClient       = NewServiceInfoForClient()
	pingPongOnlyServiceServiceInfoForStreamClient = NewServiceInfoForStreamClient()
)

// for server
func serviceInfo() *kitex.ServiceInfo {
	return pingPongOnlyServiceServiceInfo
}

// for client
func serviceInfoForStreamClient() *kitex.ServiceInfo {
	return pingPongOnlyServiceServiceInfoForStreamClient
}

// for stream client
func serviceInfoForClient() *kitex.ServiceInfo {
	return pingPongOnlyServiceServiceInfoForClient
}

// NewServiceInfo creates a new ServiceInfo containing all methods
func NewServiceInfo() *kitex.ServiceInfo {
	return newServiceInfo(false, true, true)
}

// NewServiceInfo creates a new ServiceInfo containing non-streaming methods
func NewServiceInfoForClient() *kitex.ServiceInfo {
	return newServiceInfo(false, false, true)
}
func NewServiceInfoForStreamClient() *kitex.ServiceInfo {
	return newServiceInfo(true, true, false)
}

func newServiceInfo(hasStreaming bool, keepStreamingMethods bool, keepNonStreamingMethods bool) *kitex.ServiceInfo {
	serviceName := "PingPongOnlyService"
	handlerType := (*echo.PingPongOnlyService)(nil)
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

func echoPingPongHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*echo.PingPongOnlyServiceEchoPingPongArgs)
	realResult := result.(*echo.PingPongOnlyServiceEchoPingPongResult)
	success, err := handler.(echo.PingPongOnlyService).EchoPingPong(ctx, realArg.Req1)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newPingPongOnlyServiceEchoPingPongArgs() interface{} {
	return echo.NewPingPongOnlyServiceEchoPingPongArgs()
}

func newPingPongOnlyServiceEchoPingPongResult() interface{} {
	return echo.NewPingPongOnlyServiceEchoPingPongResult()
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) EchoPingPong(ctx context.Context, req1 *echo.EchoRequest) (r *echo.EchoResponse, err error) {
	var _args echo.PingPongOnlyServiceEchoPingPongArgs
	_args.Req1 = req1
	var _result echo.PingPongOnlyServiceEchoPingPongResult
	if err = p.c.Call(ctx, "EchoPingPong", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
