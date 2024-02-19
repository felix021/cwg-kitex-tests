// Code generated by Kitex v0.8.0. DO NOT EDIT.

package lowerservice

import (
	"context"
	echo "github.com/cloudwego/kitex-tests/thrift_streaming/kitex_gen_old/echo"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
)

func serviceInfo() *kitex.ServiceInfo {
	return lowerServiceServiceInfo
}

var lowerServiceServiceInfo = NewServiceInfo()

func NewServiceInfo() *kitex.ServiceInfo {
	serviceName := "lower_service"
	handlerType := (*echo.LowerService)(nil)
	methods := map[string]kitex.MethodInfo{
		"echo_bidirectional": kitex.NewMethodInfo(echoBidirectionalHandler, newLowerServiceEchoBidirectionalArgs, newLowerServiceEchoBidirectionalResult, false),
		"echo_client":        kitex.NewMethodInfo(echoClientHandler, newLowerServiceEchoClientArgs, newLowerServiceEchoClientResult, false),
		"echo_server":        kitex.NewMethodInfo(echoServerHandler, newLowerServiceEchoServerArgs, newLowerServiceEchoServerResult, false),
		"echo_unary":         kitex.NewMethodInfo(echoUnaryHandler, newLowerServiceEchoUnaryArgs, newLowerServiceEchoUnaryResult, false),
		"echo_pingPong":      kitex.NewMethodInfo(echoPingPongHandler, newLowerServiceEchoPingPongArgs, newLowerServiceEchoPingPongResult, false),
	}
	extra := map[string]interface{}{
		"PackageName":     "echo",
		"ServiceFilePath": `idl/api.thrift`,
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
	realArg := arg.(*echo.LowerServiceEchoBidirectionalArgs)
	realResult := result.(*echo.LowerServiceEchoBidirectionalResult)
	success, err := handler.(echo.LowerService).EchoBidirectional(ctx, realArg.Req1)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newLowerServiceEchoBidirectionalArgs() interface{} {
	return echo.NewLowerServiceEchoBidirectionalArgs()
}

func newLowerServiceEchoBidirectionalResult() interface{} {
	return echo.NewLowerServiceEchoBidirectionalResult()
}

func echoClientHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*echo.LowerServiceEchoClientArgs)
	realResult := result.(*echo.LowerServiceEchoClientResult)
	success, err := handler.(echo.LowerService).EchoClient(ctx, realArg.Req1)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newLowerServiceEchoClientArgs() interface{} {
	return echo.NewLowerServiceEchoClientArgs()
}

func newLowerServiceEchoClientResult() interface{} {
	return echo.NewLowerServiceEchoClientResult()
}

func echoServerHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*echo.LowerServiceEchoServerArgs)
	realResult := result.(*echo.LowerServiceEchoServerResult)
	success, err := handler.(echo.LowerService).EchoServer(ctx, realArg.Req1)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newLowerServiceEchoServerArgs() interface{} {
	return echo.NewLowerServiceEchoServerArgs()
}

func newLowerServiceEchoServerResult() interface{} {
	return echo.NewLowerServiceEchoServerResult()
}

func echoUnaryHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
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

func (p *kClient) EchoBidirectional(ctx context.Context, req1 *echo.LowerRequest) (r *echo.LowerResponse, err error) {
	var _args echo.LowerServiceEchoBidirectionalArgs
	_args.Req1 = req1
	var _result echo.LowerServiceEchoBidirectionalResult
	if err = p.c.Call(ctx, "echo_bidirectional", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) EchoClient(ctx context.Context, req1 *echo.LowerRequest) (r *echo.LowerResponse, err error) {
	var _args echo.LowerServiceEchoClientArgs
	_args.Req1 = req1
	var _result echo.LowerServiceEchoClientResult
	if err = p.c.Call(ctx, "echo_client", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) EchoServer(ctx context.Context, req1 *echo.LowerRequest) (r *echo.LowerResponse, err error) {
	var _args echo.LowerServiceEchoServerArgs
	_args.Req1 = req1
	var _result echo.LowerServiceEchoServerResult
	if err = p.c.Call(ctx, "echo_server", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
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
