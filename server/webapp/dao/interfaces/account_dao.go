package interfaces

import (
	"server/webapp/global/errors"
	"server/webapp/model"
)

type IAccountDao interface {
	// 创建用户
	Create(fAccount *model.Account) errors.IErrCode

	// 权限验证
	Auth(fUsername, fPassword string) (*model.Account, errors.IErrCode)

	// 获取账户
	Get(accountID uint) (*model.Account, errors.IErrCode)
	GetByName(fUserName string) (*model.Account, errors.IErrCode)

	Count(account *model.Account) (int, errors.IErrCode)
}
