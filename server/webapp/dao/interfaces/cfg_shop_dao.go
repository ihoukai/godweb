package interfaces

import (
	"server/webapp/global/errors"
	"server/webapp/model"
)

// ICfgShopDao 商店配置表访问接口
type ICfgShopDao interface {
	Create(cfgShop *model.CfgShop) (err errors.IErrCode)
	Find(q IQueryInfo) (cfgShops []model.CfgShop, err errors.IErrCode)
}
