package impl

import (
	"server/webapp/dao"
	idao "server/webapp/dao/interfaces"
	"server/webapp/global"
	errs "server/webapp/global/errors"
	"server/webapp/model"
	"time"
)

// todo 依赖注入
func newAccountSrv() *accountSrv {
	return &accountSrv{}
}

type accountSrv struct {
	AccountDao idao.IAccountDao `inject:"t"`
	SessionDao idao.ISessionDao `inject:"t"`
}

//用户登录
func (a *accountSrv) Login(userName, password string) (data interface{}, err errs.IErrCode) {
	//1、先判断用户名是否存在
	//1.1、不存在，进行注册（邮箱？）
	//1.2、存在，进行密码判断
	//1.2.1存在，返回token
	//1.2.2不存在，返回错误
	if data, err = a.AccountDao.GetByName(userName); err != nil {
		if err.EnumCode() == errs.UserNotExist {
			data, err = a.Register(userName, password, userName)
			if err != nil {
				return
			}
			data, err = a._loginSuccess(data.(*model.Account))
			return
		} else if err.EnumCode() != errs.OK {
			return
		}
	}
	account := data.(*model.Account)
	if account.Password != password {
		return nil, errs.New(errs.ParamsError)
	}
	data, err = a._loginSuccess(data.(*model.Account))
	return
}

//用户注册
func (a *accountSrv) Register(fUserName, fPassword, fEmail string) (data interface{}, err errs.IErrCode) {
	//1、检查用户是否唯一性
	//2、创建Account
	//3、创建UserInfo

	// todo密码需要加密存储
	if isExist, err := a._checkNameExist(fUserName); err != nil {
		return nil, err
	} else if isExist {
		return nil, errs.New(errs.UserNameAlreadyExist)
	}

	if isExist, err := a._checkEmailExist(fPassword); err != nil {
		return nil, err
	} else if isExist {
		return nil, errs.New(errs.EmailAlreadyExist)
	}

	account := &model.Account{
		UserName:  fUserName,
		Password:  fPassword,
		Email:     fEmail,
		AccountID: uint(dao.GenerateAccountID()),
	}

	if err = a.AccountDao.Create(account); err != nil {
		return nil, err
	}
	data = account
	// todo 创建userInfoDetail
	return
}

//用户登出
func (a *accountSrv) Logout(sessionToken string) (data interface{}, err errs.IErrCode) {
	return
}

func (a *accountSrv) _loginSuccess(fAccount *model.Account) (data interface{}, err errs.IErrCode) {
	// 过期时间默认为1小时
	session, err := a.SessionDao.Create(fAccount.AccountID, global.GameID, time.Now().Add(1*time.Hour))
	if err != nil {
		return nil, err
	}

	type AccountEx struct {
		*model.Account
		Token string
	}

	account := &AccountEx{}
	account.Account = fAccount
	account.Token = session.Token
	return account, nil
}

func (a *accountSrv) _checkNameExist(fUserName string) (isExist bool, err errs.IErrCode) {
	return a._checkExist(&model.Account{
		UserName: fUserName,
	})
}

func (a *accountSrv) _checkEmailExist(fEmail string) (isExist bool, err errs.IErrCode) {
	return a._checkExist(&model.Account{
		UserName: fEmail,
	})
}

func (a *accountSrv) _checkExist(fAccount *model.Account) (isExist bool, err errs.IErrCode) {
	count, err := a.AccountDao.Count(fAccount)
	if err != nil {
		return
	}
	if count > 0 {
		isExist = true
	}
	return
}
