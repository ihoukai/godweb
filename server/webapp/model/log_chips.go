package model

import "github.com/jinzhu/gorm"

// LogChips 筹码日志记录
type LogChips struct {
	gorm.Model
	UserInfoID uint  // 玩家账户id
	Type       int   // 筹码变动类型 1 商城 2 每日签到
	Count      int64 // 扣减活增加的数量
}
