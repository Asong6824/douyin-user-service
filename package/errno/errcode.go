package errno

import (
	"github.com/Asong6824/douyin-user-service/kitex_gen/user"
)

var (
	Success                = NewErrNo(int64(user.ErrCode_SuccessCode), "Success")
	ServiceErr             = NewErrNo(int64(user.ErrCode_ServiceErrCode), "Service is unable to start successfully")
	ParamErr               = NewErrNo(int64(user.ErrCode_ParamErrCode), "Wrong Parameter has been given")
	UserAlreadyExistErr    = NewErrNo(int64(user.ErrCode_UserAlreadyExistErrCode), "User already exists")
	AuthorizationFailedErr = NewErrNo(int64(user.ErrCode_AuthorizationFailedErrCode), "Authorization failed")
)
