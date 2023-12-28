// Code generated by Kitex v0.8.0. DO NOT EDIT.

package userhandler

import (
	"context"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
	streaming "github.com/cloudwego/kitex/pkg/streaming"
	proto "google.golang.org/protobuf/proto"
	user "github.com/Asong6824/douyin-user-service/kitex_gen/user"
)

func serviceInfo() *kitex.ServiceInfo {
	return userHandlerServiceInfo
}

var userHandlerServiceInfo = NewServiceInfo()

func NewServiceInfo() *kitex.ServiceInfo {
	serviceName := "UserHandler"
	handlerType := (*user.UserHandler)(nil)
	methods := map[string]kitex.MethodInfo{
		"User":         kitex.NewMethodInfo(userHandler, newUserArgs, newUserResult, false),
		"UserRegister": kitex.NewMethodInfo(userRegisterHandler, newUserRegisterArgs, newUserRegisterResult, false),
		"UserLogin":    kitex.NewMethodInfo(userLoginHandler, newUserLoginArgs, newUserLoginResult, false),
	}
	extra := map[string]interface{}{
		"PackageName":     "user",
		"ServiceFilePath": ``,
	}
	svcInfo := &kitex.ServiceInfo{
		ServiceName:     serviceName,
		HandlerType:     handlerType,
		Methods:         methods,
		PayloadCodec:    kitex.Protobuf,
		KiteXGenVersion: "v0.8.0",
		Extra:           extra,
	}
	return svcInfo
}

func userHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(user.DouyinUserRequest)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(user.UserHandler).User(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *UserArgs:
		success, err := handler.(user.UserHandler).User(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*UserResult)
		realResult.Success = success
	}
	return nil
}
func newUserArgs() interface{} {
	return &UserArgs{}
}

func newUserResult() interface{} {
	return &UserResult{}
}

type UserArgs struct {
	Req *user.DouyinUserRequest
}

func (p *UserArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(user.DouyinUserRequest)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *UserArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *UserArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *UserArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, nil
	}
	return proto.Marshal(p.Req)
}

func (p *UserArgs) Unmarshal(in []byte) error {
	msg := new(user.DouyinUserRequest)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var UserArgs_Req_DEFAULT *user.DouyinUserRequest

func (p *UserArgs) GetReq() *user.DouyinUserRequest {
	if !p.IsSetReq() {
		return UserArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *UserArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *UserArgs) GetFirstArgument() interface{} {
	return p.Req
}

type UserResult struct {
	Success *user.DouyinUserResponse
}

var UserResult_Success_DEFAULT *user.DouyinUserResponse

func (p *UserResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(user.DouyinUserResponse)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *UserResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *UserResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *UserResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, nil
	}
	return proto.Marshal(p.Success)
}

func (p *UserResult) Unmarshal(in []byte) error {
	msg := new(user.DouyinUserResponse)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *UserResult) GetSuccess() *user.DouyinUserResponse {
	if !p.IsSetSuccess() {
		return UserResult_Success_DEFAULT
	}
	return p.Success
}

func (p *UserResult) SetSuccess(x interface{}) {
	p.Success = x.(*user.DouyinUserResponse)
}

func (p *UserResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *UserResult) GetResult() interface{} {
	return p.Success
}

func userRegisterHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(user.DouyinUserRegisterRequest)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(user.UserHandler).UserRegister(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *UserRegisterArgs:
		success, err := handler.(user.UserHandler).UserRegister(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*UserRegisterResult)
		realResult.Success = success
	}
	return nil
}
func newUserRegisterArgs() interface{} {
	return &UserRegisterArgs{}
}

func newUserRegisterResult() interface{} {
	return &UserRegisterResult{}
}

type UserRegisterArgs struct {
	Req *user.DouyinUserRegisterRequest
}

func (p *UserRegisterArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(user.DouyinUserRegisterRequest)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *UserRegisterArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *UserRegisterArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *UserRegisterArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, nil
	}
	return proto.Marshal(p.Req)
}

func (p *UserRegisterArgs) Unmarshal(in []byte) error {
	msg := new(user.DouyinUserRegisterRequest)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var UserRegisterArgs_Req_DEFAULT *user.DouyinUserRegisterRequest

func (p *UserRegisterArgs) GetReq() *user.DouyinUserRegisterRequest {
	if !p.IsSetReq() {
		return UserRegisterArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *UserRegisterArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *UserRegisterArgs) GetFirstArgument() interface{} {
	return p.Req
}

type UserRegisterResult struct {
	Success *user.DouyinUserRegisterResponse
}

var UserRegisterResult_Success_DEFAULT *user.DouyinUserRegisterResponse

func (p *UserRegisterResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(user.DouyinUserRegisterResponse)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *UserRegisterResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *UserRegisterResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *UserRegisterResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, nil
	}
	return proto.Marshal(p.Success)
}

func (p *UserRegisterResult) Unmarshal(in []byte) error {
	msg := new(user.DouyinUserRegisterResponse)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *UserRegisterResult) GetSuccess() *user.DouyinUserRegisterResponse {
	if !p.IsSetSuccess() {
		return UserRegisterResult_Success_DEFAULT
	}
	return p.Success
}

func (p *UserRegisterResult) SetSuccess(x interface{}) {
	p.Success = x.(*user.DouyinUserRegisterResponse)
}

func (p *UserRegisterResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *UserRegisterResult) GetResult() interface{} {
	return p.Success
}

func userLoginHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	switch s := arg.(type) {
	case *streaming.Args:
		st := s.Stream
		req := new(user.DouyinUserLoginRequest)
		if err := st.RecvMsg(req); err != nil {
			return err
		}
		resp, err := handler.(user.UserHandler).UserLogin(ctx, req)
		if err != nil {
			return err
		}
		if err := st.SendMsg(resp); err != nil {
			return err
		}
	case *UserLoginArgs:
		success, err := handler.(user.UserHandler).UserLogin(ctx, s.Req)
		if err != nil {
			return err
		}
		realResult := result.(*UserLoginResult)
		realResult.Success = success
	}
	return nil
}
func newUserLoginArgs() interface{} {
	return &UserLoginArgs{}
}

func newUserLoginResult() interface{} {
	return &UserLoginResult{}
}

type UserLoginArgs struct {
	Req *user.DouyinUserLoginRequest
}

func (p *UserLoginArgs) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetReq() {
		p.Req = new(user.DouyinUserLoginRequest)
	}
	return p.Req.FastRead(buf, _type, number)
}

func (p *UserLoginArgs) FastWrite(buf []byte) (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.FastWrite(buf)
}

func (p *UserLoginArgs) Size() (n int) {
	if !p.IsSetReq() {
		return 0
	}
	return p.Req.Size()
}

func (p *UserLoginArgs) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetReq() {
		return out, nil
	}
	return proto.Marshal(p.Req)
}

func (p *UserLoginArgs) Unmarshal(in []byte) error {
	msg := new(user.DouyinUserLoginRequest)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Req = msg
	return nil
}

var UserLoginArgs_Req_DEFAULT *user.DouyinUserLoginRequest

func (p *UserLoginArgs) GetReq() *user.DouyinUserLoginRequest {
	if !p.IsSetReq() {
		return UserLoginArgs_Req_DEFAULT
	}
	return p.Req
}

func (p *UserLoginArgs) IsSetReq() bool {
	return p.Req != nil
}

func (p *UserLoginArgs) GetFirstArgument() interface{} {
	return p.Req
}

type UserLoginResult struct {
	Success *user.DouyinUserLoginResponse
}

var UserLoginResult_Success_DEFAULT *user.DouyinUserLoginResponse

func (p *UserLoginResult) FastRead(buf []byte, _type int8, number int32) (n int, err error) {
	if !p.IsSetSuccess() {
		p.Success = new(user.DouyinUserLoginResponse)
	}
	return p.Success.FastRead(buf, _type, number)
}

func (p *UserLoginResult) FastWrite(buf []byte) (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.FastWrite(buf)
}

func (p *UserLoginResult) Size() (n int) {
	if !p.IsSetSuccess() {
		return 0
	}
	return p.Success.Size()
}

func (p *UserLoginResult) Marshal(out []byte) ([]byte, error) {
	if !p.IsSetSuccess() {
		return out, nil
	}
	return proto.Marshal(p.Success)
}

func (p *UserLoginResult) Unmarshal(in []byte) error {
	msg := new(user.DouyinUserLoginResponse)
	if err := proto.Unmarshal(in, msg); err != nil {
		return err
	}
	p.Success = msg
	return nil
}

func (p *UserLoginResult) GetSuccess() *user.DouyinUserLoginResponse {
	if !p.IsSetSuccess() {
		return UserLoginResult_Success_DEFAULT
	}
	return p.Success
}

func (p *UserLoginResult) SetSuccess(x interface{}) {
	p.Success = x.(*user.DouyinUserLoginResponse)
}

func (p *UserLoginResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *UserLoginResult) GetResult() interface{} {
	return p.Success
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) User(ctx context.Context, Req *user.DouyinUserRequest) (r *user.DouyinUserResponse, err error) {
	var _args UserArgs
	_args.Req = Req
	var _result UserResult
	if err = p.c.Call(ctx, "User", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) UserRegister(ctx context.Context, Req *user.DouyinUserRegisterRequest) (r *user.DouyinUserRegisterResponse, err error) {
	var _args UserRegisterArgs
	_args.Req = Req
	var _result UserRegisterResult
	if err = p.c.Call(ctx, "UserRegister", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) UserLogin(ctx context.Context, Req *user.DouyinUserLoginRequest) (r *user.DouyinUserLoginResponse, err error) {
	var _args UserLoginArgs
	_args.Req = Req
	var _result UserLoginResult
	if err = p.c.Call(ctx, "UserLogin", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
