package model

import "github.com/jinzhu/gorm"

// UserStatistics 玩家数据统计
type UserStatistics struct {
	gorm.Model
	UserInfoID  uint  // 玩家个人信息指针
	MaxChips    int64 // 历史最大筹码
	MaxWinChips int64 // 历史赢取最大筹码
	PlayCount   int32 // 游玩次数
	WinCount    int32 // 获胜次数
	MaxHand     string
}
