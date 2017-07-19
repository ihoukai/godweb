package model

import (
	"github.com/jinzhu/gorm"
	"time"
)

// UserInfo 玩家基本信息定义(一个游戏至少有一个UserInfo, 类似角色)
type UserInfo struct {
	gorm.Model
	AccountID  uint      // 账户id
	UserName   string    // 玩家昵称
	Avatar     string    // 玩家头像
	Chips      int64     // 筹码
	Vip        int32     // 玩家VIP等级 0-无等级 1-月卡 2-年卡
	VipEndDate time.Time // VIP到期时间
	Signature  string    // 玩家签名
	IsInRoom   bool      // 玩家是否在房间中
	IsOnLine   bool      // 玩家是否在线
}

// TableName 数据库表名
func (UserInfo) TableName() string {
	return "user_infos"
}
