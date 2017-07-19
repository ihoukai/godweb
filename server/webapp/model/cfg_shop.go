package model

import "github.com/jinzhu/gorm"

// CfgShop 商店配置表
type CfgShop struct {
	gorm.Model
	PID      int64     // 商品ID
	Rewards  CReward   `gorm:"type:jsonb"` // 奖励配置
	ExtraIds CExtraIDs `gorm:"type:jsonb"` // 外部渠道商品ID
	Price    float32   // 价格
	PType    int       // 商品类型 0-钻石 1-筹码 2-vip 3 其他
	Icon     int       // 道具Icon 配置
	SortNo   int       // 商品在商城中的排序
	Desc     string    // 描述信息
	Currency int       // 货比类型 0-钞票 1-钻石 2-筹码
	IsSale   bool      // 是否在售 false-下架 true-在售
	Discount float32   // 折扣
}
