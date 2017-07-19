package interfaces

import (
	"server/webapp/global/errors"
)

// IAccountSrv 账户服务
type IAccountSrv interface {
	//用户登录   输入用户名密码，返回相应token
	Login(userName, password string) (data interface{}, err errors.IErrCode)

	//用户注册
	Register(userName, password, email string) (data interface{}, err errors.IErrCode)

	//用户登出
	Logout(sessionToken string) (data interface{}, err errors.IErrCode)
}
