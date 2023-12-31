package main

import (
	"context"
	user "github.com/Asong6824/douyin-user-service/kitex_gen/user"
	"github.com/Asong6824/douyin-user-service/internal/dao"
	"github.com/Asong6824/douyin-user-service/global"
	"github.com/Asong6824/douyin-user-service/package/errno"

)

// UserServiceImpl implements the last service interface defined in the IDL.
type UserServiceImpl struct{
	dao *dao.Dao
}

func NewUserServiceImpl() *UserServiceImpl {
	dao := dao.NewDao(global.DBEngine, global.Cache)
    return &UserServiceImpl{
        dao: dao,
    }
}

// Register implements the UserServiceImpl interface.
func (s *UserServiceImpl) Register(ctx context.Context, req *user.RegisterRequest) (resp *user.RegisterResponse, err error) {
	resp = user.NewRegisterResponse()
	base := user.NewBaseResp()
	id, err := s.dao.Register(req.Username, req.Password)
	if err != nil {
		errInformation := errno.ConvertErr(err)
		base.Code = errInformation.ErrCode
		base.Msg = errInformation.ErrMsg
		return resp, err
	}
	resp.UserId = id
	return resp, nil
}

// Login implements the UserServiceImpl interface.
func (s *UserServiceImpl) Login(ctx context.Context, req *user.LoginRequest) (resp *user.LoginResponse, err error) {
	resp = user.NewLoginResponse()
	base := user.NewBaseResp()
	id, err := s.dao.Login(req.Username, req.Password)
	if err != nil {
		errInformation := errno.ConvertErr(err)
		base.Code = errInformation.ErrCode
		base.Msg = errInformation.ErrMsg
		return resp, err
	}
	resp.UserId = id
	return resp, nil
}

// Info implements the UserServiceImpl interface.
func (s *UserServiceImpl) Info(ctx context.Context, req *user.InfoRequest) (resp *user.InfoResponse, err error) {
	// TODO: Your code here...
	return
}
