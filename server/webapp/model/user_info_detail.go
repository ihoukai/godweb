package model

import (
	"time"
)

// UserInfoDetail 基本用户信息
type UserInfoDetail struct {
	UserInfo
	Diamond        int64     // 钻石
	SignInCount    int32     // 玩家连续登录次数
	TotalSignCount int32     // 玩家累计登录次数
	LastSignInDate time.Time // 玩家上次登录时间
}

// TableName 数据库表名
func (UserInfoDetail) TableName() string {
	return "user_infos"
}
