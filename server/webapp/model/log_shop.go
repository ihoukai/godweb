package model

import "github.com/jinzhu/gorm"

// LogShop 商城购买记录
type LogShop struct {
	gorm.Model
	UserInfoID uint   // 玩家信息
	ShopItemID uint   // 道具信息
	OrderID    string // 票据id
	IsSuccess  bool   // 购买状态 true 成功 true 失败
}
