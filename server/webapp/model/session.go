package model

import (
	"github.com/jinzhu/gorm"
	"time"
)

// Session 表
type Session struct {
	gorm.Model
	Token     string    // 账户token
	ExpiresAt time.Time // 过期时间
	AccountID uint      // 用户id
	GameID    uint      // 游戏id, for 多游戏公用一个账户系统
}
