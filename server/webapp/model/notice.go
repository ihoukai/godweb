package model

import "github.com/jinzhu/gorm"

// Notice 公告表
type Notice struct {
	gorm.Model
	Title   string // 公告标题
	Content string // 公告内容
	Type    int    // 公告类型 0-公告 1-活动
	Image   string // 公告图片
}
