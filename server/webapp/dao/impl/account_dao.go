package impl

import (
	"server/webapp/dao/interfaces"
	"server/webapp/global/errors"
	"server/webapp/model"
)

// accountDao account表
type accountDao struct {
	baseDao interfaces.IBaseDao
}

// Auth 验证用户名密码
func (u *accountDao) Auth(username, pwd string) (account *model.Account, err errors.IErrCode) {
	account = &model.Account{}
	e := u.baseDao.Find(account, NewQueryInfo().Where(&model.Account{
		UserName: username,
	}))
	if e != nil {
		account = nil
		err = errors.New(errors.UserNotExist, e)
		return
	}
	if account.Password != pwd {
		account = nil
		err = errors.New(errors.PwdError, e)
		return
	}
	return
}

func (u *accountDao) Get(accountID uint) (account *model.Account, err errors.IErrCode) {
	account = &model.Account{}
	e := u.baseDao.Find(account, NewQueryInfo().Where(&model.Account{
		AccountID: accountID,
	}))
	if e != nil {
		account = nil
		err = errors.New(errors.UserNotExist, e)
		return
	}
	return
}

func (u *accountDao) GetByName(fUserName string) (fAccount *model.Account, err errors.IErrCode) {
	fAccount = &model.Account{}
	e := u.baseDao.Find(fAccount, NewQueryInfo().Where(&model.Account{
		UserName: fUserName,
	}))
	if e != nil {
		fAccount = nil
		err = errors.New(errors.UserNotExist, e)
		return
	}
	return
}

func (u *accountDao) Create(account *model.Account) (err errors.IErrCode) {
	e := u.baseDao.Create(account)
	if e != nil {
		account = nil
		err = errors.New(errors.CreateUserFailed, e)
		return
	}
	return
}

func (u *accountDao) Count(account *model.Account) (count int, err errors.IErrCode) {
	e := u.baseDao.Count(&model.Account{}, &count, NewQueryInfo().Where(account))
	if e != nil {
		err = errors.New(errors.Unknow, err)
		return
	}
	return
}
