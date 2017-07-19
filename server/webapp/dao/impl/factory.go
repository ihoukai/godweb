package impl

import (
	"github.com/jinzhu/gorm"
	"server/webapp/dao/interfaces"
)

// Factory 工厂类
type Factory struct {
	db      *gorm.DB
	baseDao interfaces.IBaseDao
}

// NewFactory 实例化工厂对象
func NewFactory(db *gorm.DB) *Factory {
	f := &Factory{}
	f.db = db
	f.baseDao = &baseDao{
		db: f.db,
	}
	return f
}

// NewAccountDao 数据访问对象
func (f *Factory) NewAccountDao() interfaces.IAccountDao {
	return &accountDao{
		baseDao: f.baseDao,
	}
}

// NewCfgShopDao 数据访问对象
func (f *Factory) NewCfgShopDao() interfaces.ICfgShopDao {
	return &cfgShopDao{
		baseDao: f.baseDao,
	}
}

// NewSessionDao 数据访问对象
func (f *Factory) NewSessionDao() interfaces.ISessionDao {
	return &sessionDao{
		baseDao: f.baseDao,
	}
}
