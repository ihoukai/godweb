package model

import (
	"github.com/jinzhu/gorm"
)

// UserFriend 玩家好友表
type UserFriend struct {
	gorm.Model
	UserInfoID uint // 玩家自己id
	Friend     uint // 好友id
	State      int  // 好友状态	0-好友请求 1-好友 2-拒绝
}
