package dao

// import (
// 	"encoding/json"
// 	"github.com/tuvistavie/structomap"
// 	. "gopkg.in/check.v1"
// 	"reflect"
// 	"server/webapp/global/errors"
// 	"server/webapp/model"
// )

// var _cfgShopSerializer structomap.Serializer

// func init() {
// 	_cfgShopSerializer = structomap.New().
// 		UseSnakeCase().
// 		Pick("PID",
// 			"Price",
// 			"PType",
// 			"Icon",
// 			"SortNo",
// 			"Desc",
// 			"Currency",
// 			"IsSale",
// 			"Discount",
// 			"Rewards",
// 			"ExtraIds")
// }

// // TestCfgShop 测试配置表
// func (s *TestSuite) TestCfgShop(c *C) {
// 	cfgShop1 := &model.CfgShop{
// 		PID:      1,
// 		Price:    1,
// 		PType:    1,
// 		Icon:     1,
// 		SortNo:   1,
// 		Desc:     "1",
// 		Currency: 1,
// 		IsSale:   true,
// 		Discount: 1,
// 		Rewards: model.CReward{
// 			Chips:   1,
// 			Diamond: 1,
// 			Vip1:    1,
// 			Vip2:    1,
// 		},
// 		ExtraIds: model.CExtraIDs{
// 			AppStoreID: "1",
// 			GoogPlayID: "1",
// 		},
// 	}

// 	cfgShop2 := &model.CfgShop{
// 		PID:      2,
// 		Price:    2,
// 		PType:    2,
// 		Icon:     2,
// 		SortNo:   2,
// 		Desc:     "2",
// 		Currency: 2,
// 		IsSale:   true,
// 		Discount: 2,
// 		Rewards: model.CReward{
// 			Chips:   2,
// 			Diamond: 2,
// 			Vip1:    2,
// 			Vip2:    2,
// 		},
// 		ExtraIds: model.CExtraIDs{
// 			AppStoreID: "2",
// 			GoogPlayID: "2",
// 		},
// 	}

// 	expectMap1 := _cfgShopSerializer.Transform(cfgShop1)
// 	expectjson1, err := json.MarshalIndent(expectMap1, "", "  ")
// 	if err != nil {
// 		c.Error(err)
// 	}
// 	expectMap2 := _cfgShopSerializer.Transform(cfgShop2)
// 	expectjson2, err := json.MarshalIndent(expectMap2, "", "  ")
// 	if err != nil {
// 		c.Error(err)
// 	}

// 	CfgShop().Create(cfgShop1)
// 	CfgShop().Create(cfgShop2)
// 	cfgs, retCode := CfgShop().Find(nil)
// 	if retCode != int(RetCode.OK) {
// 		c.Error(retCode)
// 	}
// 	if len(cfgs) != 2 {
// 		c.Errorf("Num not Match!")
// 	}

// 	for _, cfg := range cfgs {
// 		retMap := _cfgShopSerializer.Transform(cfg)
// 		retjson, err := json.MarshalIndent(retMap, "", "  ")
// 		if err != nil {
// 			c.Error(err)
// 		}

// 		if !reflect.DeepEqual(retjson, expectjson1) &&
// 			!reflect.DeepEqual(retjson, expectjson2) {
// 			c.Errorf("Result not Match!")
// 		}
// 	}
// }
