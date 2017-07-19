package impl_test

import (
	errs "server/webapp/global/errors"
	"server/webapp/log"
	"server/webapp/model"
)

// todo 依赖注入
func newAccountSrv() *accountSrv {
	return &accountSrv{}
}

type accountSrv struct {
}

//用户登录
func (a *accountSrv) Login(fUserName, fPassword string) (data interface{}, err errs.IErrCode) {
	log.Info("Login => %s %s\n", fUserName, fPassword)
	account := &model.Account{
		UserName: fUserName,
		Password: fPassword,
	}
	data = account
	return account, nil
}

//用户注册
func (a *accountSrv) Register(fUserName, fPassword, fEmail string) (data interface{}, err errs.IErrCode) {
	log.Info("Register => %s %s %s\n", fUserName, fPassword, fEmail)
	account := &model.Account{
		UserName: fUserName,
		Password: fPassword,
		Email:    fEmail,
	}
	data = account
	return account, nil
}

//用户登出
func (a *accountSrv) Logout(sessionToken string) (data interface{}, err errs.IErrCode) {
	log.Info("Logout => %s\n", sessionToken)
	return "", nil
}
