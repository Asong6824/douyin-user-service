package main

import (
	"context"
	user "github.com/Asong6824/douyin-user-service/kitex_gen/user"
)

// UserHandlerImpl implements the last service interface defined in the IDL.
type UserHandlerImpl struct{}

// User implements the UserHandlerImpl interface.
func (s *UserHandlerImpl) User(ctx context.Context, req *user.DouyinUserRequest) (resp *user.DouyinUserResponse, err error) {
	// TODO: Your code here...
	return
}

// UserRegister implements the UserHandlerImpl interface.
func (s *UserHandlerImpl) UserRegister(ctx context.Context, req *user.DouyinUserRegisterRequest) (resp *user.DouyinUserRegisterResponse, err error) {
	// TODO: Your code here...
	return
}

// UserLogin implements the UserHandlerImpl interface.
func (s *UserHandlerImpl) UserLogin(ctx context.Context, req *user.DouyinUserLoginRequest) (resp *user.DouyinUserLoginResponse, err error) {
	// TODO: Your code here...
	return
}
