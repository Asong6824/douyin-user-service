package errno

import (
	"github.com/Asong6824/douyin-user-service/kitex_gen/user"
)

var (
	Success                = NewErrNo(int64(user.ErrCode_SuccessCode), "Success")
	UserServiceErr             = NewErrNo(int64(user.ErrCode_UserServiceErrCode), "User Service is unable to start successfully")
	UserAlreadyExistErr    = NewErrNo(int64(user.ErrCode_UserAlreadyExistErrCode), "User already exists")
	UserNotExistErr        = NewErrNo(int64(user.ErrCode_UserNotExistErrcode), "User does not exist")
)
