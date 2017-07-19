package interfaces

import (
	"server/webapp/global/errors"
	"server/webapp/model"
	"time"
)

type ISessionDao interface {
	// 验证session，判断有没有过期
	Auth(sessionToken string) (*model.Session, errors.IErrCode)

	// 创建
	Create(accountID uint, gameID uint, expiresAt time.Time) (*model.Session, errors.IErrCode)

	// 删除
	Delete(sessionToken string) errors.IErrCode
}
