package dao

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"sync"
)

var (
	_idZones   []IDZone
	_curIDZone *IDZone
	_mutex     *sync.Mutex
)

func init() {
	_mutex = new(sync.Mutex)
}

// IDZone id区段表
type IDZone struct {
	ID      uint  `gorm:"primary_key"` // 数据库自增id
	StartID int64 `sql:"unique_index"` // 起始id
	EndID   int64 // 截至id
	CurID   int64 // 当前用户id
}

// TableName 表名
func (IDZone) TableName() string {
	return "_id_zones"
}

// GenerateAccountID 获取生成对用户友好的id
func GenerateAccountID() int64 {
	_mutex.Lock()
	defer _mutex.Unlock()
	// 该方法完成后已经将_curIDZone.CurID自增
	err := _db.Model(&_curIDZone).Update("cur_id", _curIDZone.CurID+1).Error
	if err != nil {
		panic(fmt.Sprintf("IDZone Update Err:%s\n", err.Error()))
	}
	// todo 区段用完切换区段 没有有效区段则提示报错
	return _curIDZone.CurID
}

// checkIDZones 检查idzones
func checkIDZones(idZones []*IDZone) {
	// todo 检查id区间之前是否有交集
	for _, IDZone := range idZones {
		if IDZone.StartID < 0 || IDZone.EndID < 0 {
			panic("IDZone StartID and EndID must greater than zero")
		}
		if IDZone.EndID-IDZone.StartID <= 0 {
			panic("IDZone EndID must greater than StartID")
		}
	}
}

// 初始化id区段表
func initIDZone(_db *gorm.DB) {
	// 从配置文件读取IDZone并插入数据库
	for _, idzone := range _config.IDZones {
		zone := &IDZone{}
		_db.Where("start_id = ?", idzone.StartID).Find(zone)
		if zone.ID == 0 {
			// not exist
			if idzone.CurID == 0 {
				idzone.CurID = idzone.StartID
			}
			err := _db.Create(idzone).Error
			if err != nil {
				panic(fmt.Sprintf("No error should happen when initIDZone, but got %+v", err))
			}
		} else {
			// 区段已经在数据库里面存在 则跳过
			if zone.StartID == idzone.StartID &&
				zone.EndID == idzone.EndID {
				continue
			}
			panic(fmt.Sprintf("zoneids not matching, database:[%d-%d], config[%d-%d]",
				zone.StartID, zone.EndID, idzone.StartID, idzone.EndID))
		}
	}

	// 从数据库读取所有的未使用完毕的IDZones
	err := _db.Where("end_id != cur_id").Find(&_idZones).Error
	if err != nil {
		panic(fmt.Sprintf("No error should happen when initIDZone, but got %+v", err))
	}
	if len(_idZones) == 0 {
		panic("IdZones no available ids")
	}

	// 设置当前id
	_curIDZone = &_idZones[0]
}
