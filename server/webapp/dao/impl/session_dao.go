package impl

import (
	"fmt"
	"github.com/satori/go.uuid"
	"server/webapp/dao/interfaces"
	"server/webapp/global/errors"
	"server/webapp/model"
	"time"
)

// sessionDao session表
type sessionDao struct {
	baseDao interfaces.IBaseDao
}

// 验证session，判断有没有过期
func (s *sessionDao) Auth(sessionToken string) (session *model.Session, err errors.IErrCode) {
	var se model.Session
	e := s.baseDao.Find(&se, NewQueryInfo().Where("token = ?", sessionToken))
	if e != nil {
		session = nil
		err = errors.New(errors.InvalidSessionToken, e)
		return
	}
	// token 过期
	if se.ExpiresAt.Before(time.Now()) {
		session = nil
		s.Delete(sessionToken)
		err = errors.New(errors.ExpiredSessionToken, fmt.Errorf("token %s expires", sessionToken))
		return
	}
	session = &se
	return
}

// 创建
func (s *sessionDao) Create(accountID uint, gameID uint, expiresAt time.Time) (session *model.Session, err errors.IErrCode) {
	// 删除存在的session
	var sessions []model.Session
	s.baseDao.Find(&sessions, NewQueryInfo().Where("account_id = ?", accountID).Where("game_id = ?", gameID))
	if sessions != nil {
		for _, session := range sessions {
			s.Delete(session.Token)
		}
	}

	// 插入新的session
	session = &model.Session{
		AccountID: accountID,
		GameID:    gameID,
		ExpiresAt: expiresAt,
		Token:     uuid.NewV4().String(),
	}
	e := s.baseDao.Create(session)
	if e != nil {
		session = nil
		err = errors.New(errors.CreateTokenFailed, e)
	}
	return
}

// 删除
func (s *sessionDao) Delete(sessionToken string) (err errors.IErrCode) {
	return
}
