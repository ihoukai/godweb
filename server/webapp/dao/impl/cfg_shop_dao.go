package impl

import (
	"server/webapp/dao/interfaces"
	"server/webapp/global/errors"
	"server/webapp/model"
)

type cfgShopDao struct {
	baseDao interfaces.IBaseDao
}

func (c *cfgShopDao) Create(cfgShop *model.CfgShop) (err errors.IErrCode) {
	e := c.baseDao.Create(cfgShop)
	if e != nil {
		err = errors.New(errors.CreateCfgShopFailed, e)
	}
	return
}

func (c *cfgShopDao) Find(q interfaces.IQueryInfo) (cfgShops []model.CfgShop, err errors.IErrCode) {
	e := c.baseDao.Find(&cfgShops, q)
	if e != nil {
		err = errors.New(errors.NotFound, e)
	}
	return
}
