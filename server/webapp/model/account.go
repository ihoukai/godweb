package model

import "github.com/jinzhu/gorm"

// Account 玩家账号(多个游戏只有一个User)
type Account struct {
	gorm.Model
	AccountID uint   // id为8位整数, 便于玩家搜索添加好友，类似于QQ号码; todo 用分区段的方式生成用户id
	UserName  string `gorm:"index:idx_name_code"`
	Password  string `json:"-"`
	Email     string
	AuthData  string
}
