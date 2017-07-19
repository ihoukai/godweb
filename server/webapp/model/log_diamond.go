package model

import "github.com/jinzhu/gorm"

// LogDiamond 钻石日志记录
type LogDiamond struct {
	gorm.Model
	UserInfoID uint  // 玩家Id
	Type       int   // 钻石变更类型 1 商城 2 每日签到
	Count      int64 // 扣减或者增加的数量
}
