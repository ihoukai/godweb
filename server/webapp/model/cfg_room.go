package model

import "github.com/jinzhu/gorm"

// CfgRoom 房间配置表
type CfgRoom struct {
	gorm.Model
	ConfigID int    // 房间配置id
	Type     int    // 房间类型
	Name     string // 名称
	Bet      int64  // 单注金额
	MinTake  int64  // 最小携带
}
