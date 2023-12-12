// Code generated by Kitex v0.8.0. DO NOT EDIT.

package echo

import (
	"bytes"
	"fmt"
	"reflect"
	"strings"

	"github.com/apache/thrift/lib/go/thrift"
)

// unused protection
var (
	_ = fmt.Formatter(nil)
	_ = (*bytes.Buffer)(nil)
	_ = (*strings.Builder)(nil)
	_ = reflect.Type(nil)
	_ = thrift.TProtocol(nil)
)

type EchoServiceEchoBidirectionalArgs struct {
	Req1 *EchoRequest `thrift:"req1,1" frugal:"1,default,EchoRequest" json:"req1"`
}

func NewEchoServiceEchoBidirectionalArgs() *EchoServiceEchoBidirectionalArgs {
	return &EchoServiceEchoBidirectionalArgs{}
}

func (p *EchoServiceEchoBidirectionalArgs) InitDefault() {
	*p = EchoServiceEchoBidirectionalArgs{}
}

var EchoServiceEchoBidirectionalArgs_Req1_DEFAULT *EchoRequest

func (p *EchoServiceEchoBidirectionalArgs) GetReq1() (v *EchoRequest) {
	if !p.IsSetReq1() {
		return EchoServiceEchoBidirectionalArgs_Req1_DEFAULT
	}
	return p.Req1
}
func (p *EchoServiceEchoBidirectionalArgs) SetReq1(val *EchoRequest) {
	p.Req1 = val
}

func (p *EchoServiceEchoBidirectionalArgs) IsSetReq1() bool {
	return p.Req1 != nil
}

func (p *EchoServiceEchoBidirectionalArgs) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("EchoServiceEchoBidirectionalArgs(%+v)", *p)
}
func (p *EchoServiceEchoBidirectionalArgs) GetFirstArgument() interface{} {
	return p.Req1
}

type EchoServiceEchoBidirectionalResult struct {
	Success *EchoResponse `thrift:"success,0,optional" frugal:"0,optional,EchoResponse" json:"success,omitempty"`
}

func NewEchoServiceEchoBidirectionalResult() *EchoServiceEchoBidirectionalResult {
	return &EchoServiceEchoBidirectionalResult{}
}

func (p *EchoServiceEchoBidirectionalResult) InitDefault() {
	*p = EchoServiceEchoBidirectionalResult{}
}

var EchoServiceEchoBidirectionalResult_Success_DEFAULT *EchoResponse

func (p *EchoServiceEchoBidirectionalResult) GetSuccess() (v *EchoResponse) {
	if !p.IsSetSuccess() {
		return EchoServiceEchoBidirectionalResult_Success_DEFAULT
	}
	return p.Success
}
func (p *EchoServiceEchoBidirectionalResult) SetSuccess(x interface{}) {
	p.Success = x.(*EchoResponse)
}

func (p *EchoServiceEchoBidirectionalResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *EchoServiceEchoBidirectionalResult) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("EchoServiceEchoBidirectionalResult(%+v)", *p)
}
func (p *EchoServiceEchoBidirectionalResult) GetResult() interface{} {
	return p.Success
}

type EchoServiceEchoClientArgs struct {
	Req1 *EchoRequest `thrift:"req1,1" frugal:"1,default,EchoRequest" json:"req1"`
}

func NewEchoServiceEchoClientArgs() *EchoServiceEchoClientArgs {
	return &EchoServiceEchoClientArgs{}
}

func (p *EchoServiceEchoClientArgs) InitDefault() {
	*p = EchoServiceEchoClientArgs{}
}

var EchoServiceEchoClientArgs_Req1_DEFAULT *EchoRequest

func (p *EchoServiceEchoClientArgs) GetReq1() (v *EchoRequest) {
	if !p.IsSetReq1() {
		return EchoServiceEchoClientArgs_Req1_DEFAULT
	}
	return p.Req1
}
func (p *EchoServiceEchoClientArgs) SetReq1(val *EchoRequest) {
	p.Req1 = val
}

func (p *EchoServiceEchoClientArgs) IsSetReq1() bool {
	return p.Req1 != nil
}

func (p *EchoServiceEchoClientArgs) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("EchoServiceEchoClientArgs(%+v)", *p)
}
func (p *EchoServiceEchoClientArgs) GetFirstArgument() interface{} {
	return p.Req1
}

type EchoServiceEchoClientResult struct {
	Success *EchoResponse `thrift:"success,0,optional" frugal:"0,optional,EchoResponse" json:"success,omitempty"`
}

func NewEchoServiceEchoClientResult() *EchoServiceEchoClientResult {
	return &EchoServiceEchoClientResult{}
}

func (p *EchoServiceEchoClientResult) InitDefault() {
	*p = EchoServiceEchoClientResult{}
}

var EchoServiceEchoClientResult_Success_DEFAULT *EchoResponse

func (p *EchoServiceEchoClientResult) GetSuccess() (v *EchoResponse) {
	if !p.IsSetSuccess() {
		return EchoServiceEchoClientResult_Success_DEFAULT
	}
	return p.Success
}
func (p *EchoServiceEchoClientResult) SetSuccess(x interface{}) {
	p.Success = x.(*EchoResponse)
}

func (p *EchoServiceEchoClientResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *EchoServiceEchoClientResult) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("EchoServiceEchoClientResult(%+v)", *p)
}
func (p *EchoServiceEchoClientResult) GetResult() interface{} {
	return p.Success
}

type EchoServiceEchoServerArgs struct {
	Req1 *EchoRequest `thrift:"req1,1" frugal:"1,default,EchoRequest" json:"req1"`
}

func NewEchoServiceEchoServerArgs() *EchoServiceEchoServerArgs {
	return &EchoServiceEchoServerArgs{}
}

func (p *EchoServiceEchoServerArgs) InitDefault() {
	*p = EchoServiceEchoServerArgs{}
}

var EchoServiceEchoServerArgs_Req1_DEFAULT *EchoRequest

func (p *EchoServiceEchoServerArgs) GetReq1() (v *EchoRequest) {
	if !p.IsSetReq1() {
		return EchoServiceEchoServerArgs_Req1_DEFAULT
	}
	return p.Req1
}
func (p *EchoServiceEchoServerArgs) SetReq1(val *EchoRequest) {
	p.Req1 = val
}

func (p *EchoServiceEchoServerArgs) IsSetReq1() bool {
	return p.Req1 != nil
}

func (p *EchoServiceEchoServerArgs) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("EchoServiceEchoServerArgs(%+v)", *p)
}
func (p *EchoServiceEchoServerArgs) GetFirstArgument() interface{} {
	return p.Req1
}

type EchoServiceEchoServerResult struct {
	Success *EchoResponse `thrift:"success,0,optional" frugal:"0,optional,EchoResponse" json:"success,omitempty"`
}

func NewEchoServiceEchoServerResult() *EchoServiceEchoServerResult {
	return &EchoServiceEchoServerResult{}
}

func (p *EchoServiceEchoServerResult) InitDefault() {
	*p = EchoServiceEchoServerResult{}
}

var EchoServiceEchoServerResult_Success_DEFAULT *EchoResponse

func (p *EchoServiceEchoServerResult) GetSuccess() (v *EchoResponse) {
	if !p.IsSetSuccess() {
		return EchoServiceEchoServerResult_Success_DEFAULT
	}
	return p.Success
}
func (p *EchoServiceEchoServerResult) SetSuccess(x interface{}) {
	p.Success = x.(*EchoResponse)
}

func (p *EchoServiceEchoServerResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *EchoServiceEchoServerResult) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("EchoServiceEchoServerResult(%+v)", *p)
}
func (p *EchoServiceEchoServerResult) GetResult() interface{} {
	return p.Success
}

type EchoServiceEchoUnaryArgs struct {
	Req1 *EchoRequest `thrift:"req1,1" frugal:"1,default,EchoRequest" json:"req1"`
}

func NewEchoServiceEchoUnaryArgs() *EchoServiceEchoUnaryArgs {
	return &EchoServiceEchoUnaryArgs{}
}

func (p *EchoServiceEchoUnaryArgs) InitDefault() {
	*p = EchoServiceEchoUnaryArgs{}
}

var EchoServiceEchoUnaryArgs_Req1_DEFAULT *EchoRequest

func (p *EchoServiceEchoUnaryArgs) GetReq1() (v *EchoRequest) {
	if !p.IsSetReq1() {
		return EchoServiceEchoUnaryArgs_Req1_DEFAULT
	}
	return p.Req1
}
func (p *EchoServiceEchoUnaryArgs) SetReq1(val *EchoRequest) {
	p.Req1 = val
}

func (p *EchoServiceEchoUnaryArgs) IsSetReq1() bool {
	return p.Req1 != nil
}

func (p *EchoServiceEchoUnaryArgs) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("EchoServiceEchoUnaryArgs(%+v)", *p)
}
func (p *EchoServiceEchoUnaryArgs) GetFirstArgument() interface{} {
	return p.Req1
}

type EchoServiceEchoUnaryResult struct {
	Success *EchoResponse `thrift:"success,0,optional" frugal:"0,optional,EchoResponse" json:"success,omitempty"`
}

func NewEchoServiceEchoUnaryResult() *EchoServiceEchoUnaryResult {
	return &EchoServiceEchoUnaryResult{}
}

func (p *EchoServiceEchoUnaryResult) InitDefault() {
	*p = EchoServiceEchoUnaryResult{}
}

var EchoServiceEchoUnaryResult_Success_DEFAULT *EchoResponse

func (p *EchoServiceEchoUnaryResult) GetSuccess() (v *EchoResponse) {
	if !p.IsSetSuccess() {
		return EchoServiceEchoUnaryResult_Success_DEFAULT
	}
	return p.Success
}
func (p *EchoServiceEchoUnaryResult) SetSuccess(x interface{}) {
	p.Success = x.(*EchoResponse)
}

func (p *EchoServiceEchoUnaryResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *EchoServiceEchoUnaryResult) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("EchoServiceEchoUnaryResult(%+v)", *p)
}
func (p *EchoServiceEchoUnaryResult) GetResult() interface{} {
	return p.Success
}

type EchoServiceEchoPingPongArgs struct {
	Req1 *EchoRequest `thrift:"req1,1" frugal:"1,default,EchoRequest" json:"req1"`
	Req2 *EchoRequest `thrift:"req2,2" frugal:"2,default,EchoRequest" json:"req2"`
}

func NewEchoServiceEchoPingPongArgs() *EchoServiceEchoPingPongArgs {
	return &EchoServiceEchoPingPongArgs{}
}

func (p *EchoServiceEchoPingPongArgs) InitDefault() {
	*p = EchoServiceEchoPingPongArgs{}
}

var EchoServiceEchoPingPongArgs_Req1_DEFAULT *EchoRequest

func (p *EchoServiceEchoPingPongArgs) GetReq1() (v *EchoRequest) {
	if !p.IsSetReq1() {
		return EchoServiceEchoPingPongArgs_Req1_DEFAULT
	}
	return p.Req1
}

var EchoServiceEchoPingPongArgs_Req2_DEFAULT *EchoRequest

func (p *EchoServiceEchoPingPongArgs) GetReq2() (v *EchoRequest) {
	if !p.IsSetReq2() {
		return EchoServiceEchoPingPongArgs_Req2_DEFAULT
	}
	return p.Req2
}
func (p *EchoServiceEchoPingPongArgs) SetReq1(val *EchoRequest) {
	p.Req1 = val
}
func (p *EchoServiceEchoPingPongArgs) SetReq2(val *EchoRequest) {
	p.Req2 = val
}

func (p *EchoServiceEchoPingPongArgs) IsSetReq1() bool {
	return p.Req1 != nil
}

func (p *EchoServiceEchoPingPongArgs) IsSetReq2() bool {
	return p.Req2 != nil
}

func (p *EchoServiceEchoPingPongArgs) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("EchoServiceEchoPingPongArgs(%+v)", *p)
}
func (p *EchoServiceEchoPingPongArgs) GetFirstArgument() interface{} {
	return p.Req1
}

type EchoServiceEchoPingPongResult struct {
	Success *EchoResponse  `thrift:"success,0,optional" frugal:"0,optional,EchoResponse" json:"success,omitempty"`
	E       *EchoException `thrift:"e,1,optional" frugal:"1,optional,EchoException" json:"e,omitempty"`
}

func NewEchoServiceEchoPingPongResult() *EchoServiceEchoPingPongResult {
	return &EchoServiceEchoPingPongResult{}
}

func (p *EchoServiceEchoPingPongResult) InitDefault() {
	*p = EchoServiceEchoPingPongResult{}
}

var EchoServiceEchoPingPongResult_Success_DEFAULT *EchoResponse

func (p *EchoServiceEchoPingPongResult) GetSuccess() (v *EchoResponse) {
	if !p.IsSetSuccess() {
		return EchoServiceEchoPingPongResult_Success_DEFAULT
	}
	return p.Success
}

var EchoServiceEchoPingPongResult_E_DEFAULT *EchoException

func (p *EchoServiceEchoPingPongResult) GetE() (v *EchoException) {
	if !p.IsSetE() {
		return EchoServiceEchoPingPongResult_E_DEFAULT
	}
	return p.E
}
func (p *EchoServiceEchoPingPongResult) SetSuccess(x interface{}) {
	p.Success = x.(*EchoResponse)
}
func (p *EchoServiceEchoPingPongResult) SetE(val *EchoException) {
	p.E = val
}

func (p *EchoServiceEchoPingPongResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *EchoServiceEchoPingPongResult) IsSetE() bool {
	return p.E != nil
}

func (p *EchoServiceEchoPingPongResult) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("EchoServiceEchoPingPongResult(%+v)", *p)
}
func (p *EchoServiceEchoPingPongResult) GetResult() interface{} {
	return p.Success
}

type EchoServiceEchoOnewayArgs struct {
	Req1 *EchoRequest `thrift:"req1,1" frugal:"1,default,EchoRequest" json:"req1"`
}

func NewEchoServiceEchoOnewayArgs() *EchoServiceEchoOnewayArgs {
	return &EchoServiceEchoOnewayArgs{}
}

func (p *EchoServiceEchoOnewayArgs) InitDefault() {
	*p = EchoServiceEchoOnewayArgs{}
}

var EchoServiceEchoOnewayArgs_Req1_DEFAULT *EchoRequest

func (p *EchoServiceEchoOnewayArgs) GetReq1() (v *EchoRequest) {
	if !p.IsSetReq1() {
		return EchoServiceEchoOnewayArgs_Req1_DEFAULT
	}
	return p.Req1
}
func (p *EchoServiceEchoOnewayArgs) SetReq1(val *EchoRequest) {
	p.Req1 = val
}

func (p *EchoServiceEchoOnewayArgs) IsSetReq1() bool {
	return p.Req1 != nil
}

func (p *EchoServiceEchoOnewayArgs) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("EchoServiceEchoOnewayArgs(%+v)", *p)
}
func (p *EchoServiceEchoOnewayArgs) GetFirstArgument() interface{} {
	return p.Req1
}

type EchoServiceEchoOnewayResult struct {
}

func NewEchoServiceEchoOnewayResult() *EchoServiceEchoOnewayResult {
	return &EchoServiceEchoOnewayResult{}
}

func (p *EchoServiceEchoOnewayResult) InitDefault() {
	*p = EchoServiceEchoOnewayResult{}
}

func (p *EchoServiceEchoOnewayResult) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("EchoServiceEchoOnewayResult(%+v)", *p)
}
func (p *EchoServiceEchoOnewayResult) GetResult() interface{} {
	return nil
}

type EchoServicePingArgs struct {
}

func NewEchoServicePingArgs() *EchoServicePingArgs {
	return &EchoServicePingArgs{}
}

func (p *EchoServicePingArgs) InitDefault() {
	*p = EchoServicePingArgs{}
}

func (p *EchoServicePingArgs) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("EchoServicePingArgs(%+v)", *p)
}
func (p *EchoServicePingArgs) GetFirstArgument() interface{} {
	return nil
}

type EchoServicePingResult struct {
}

func NewEchoServicePingResult() *EchoServicePingResult {
	return &EchoServicePingResult{}
}

func (p *EchoServicePingResult) InitDefault() {
	*p = EchoServicePingResult{}
}

func (p *EchoServicePingResult) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("EchoServicePingResult(%+v)", *p)
}
func (p *EchoServicePingResult) GetResult() interface{} {
	return nil
}
