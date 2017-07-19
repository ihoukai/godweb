package model

import "github.com/jinzhu/gorm"

// CfgDailySignIn 每日签到表
type CfgDailySignIn struct {
	gorm.Model
	Type    int     // 签到类型 1-连续签到 2-累计签到
	Days    int     // 领取奖励需要达到的天数
	Rewards CReward `gorm:"type:jsonb"` // 奖励信息
}
