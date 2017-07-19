package model

import (
	"database/sql/driver"
	"encoding/json"
)

// CReward 奖励信息
type CReward struct {
	Chips   int64 // 筹码
	Diamond int64 // 钻石
	Vip1    int   // 月卡?
	Vip2    int   // 年卡?
}

// Value 数据库读取之后解析
func (c CReward) Value() (driver.Value, error) {
	return json.Marshal(c)
}

// Scan 存入数据库之前编码
func (c *CReward) Scan(input interface{}) error {
	return json.Unmarshal(input.([]byte), c)
}

// CExtraIDs 商城道具外部id
type CExtraIDs struct {
	AppStoreID string // AppStore ID
	GoogPlayID string // GoogPlay ID
}

// Value 数据库读取之后解析
func (c CExtraIDs) Value() (driver.Value, error) {
	return json.Marshal(c)
}

// Scan 存入数据库之前编码
func (c *CExtraIDs) Scan(input interface{}) error {
	return json.Unmarshal(input.([]byte), c)
}
