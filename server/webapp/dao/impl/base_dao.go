package impl

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"server/webapp/dao/interfaces"
)

type baseDao struct {
	db *gorm.DB
}

func (b *baseDao) Create(obj interface{}) (err error) {
	return b.db.Create(obj).Error
}

func (b *baseDao) Update(obj interface{}) (err error) {
	return
}

func (b *baseDao) Delete(id uint) (err error) {
	return
}

func (b *baseDao) Find(out interface{}, q interfaces.IQueryInfo) (err error) {
	if q == nil {
		return b.db.Find(out).Error
	}
	queryInfo := q.(*QueryInfo)
	if queryInfo == nil {
		return b.db.Find(out).Error
	}
	return handleQueryInfo(b.db, queryInfo).Find(out).Error
}

func (b *baseDao) Count(table interface{}, out interface{}, q interfaces.IQueryInfo) (err error) {
	if q == nil {
		return b.db.Find(out).Error
	}
	queryInfo := q.(*QueryInfo)
	if queryInfo == nil {
		return b.db.Find(out).Error
	}
	return handleQueryInfo(b.db, queryInfo).Model(table).Count(out).Error
}

func handleQueryInfo(_db *gorm.DB, queryInfo *QueryInfo) *gorm.DB {
	cur := _db
	if queryInfo.wheres != nil {
		for _, whereArray := range queryInfo.wheres {
			if len(whereArray) != 0 {
				if len(whereArray) == 1 {
					cur = cur.Where(whereArray[0])
				} else {
					cur = cur.Where(whereArray[0], whereArray[1:]...)
				}
			}
		}
	}
	if queryInfo.ors != nil {
		for _, orArray := range queryInfo.ors {
			if len(orArray) != 0 {
				if len(orArray) == 1 {
					cur = cur.Or(orArray[0])
				} else {
					cur = cur.Or(orArray[0], orArray[1:]...)
				}
			}
		}
	}
	if queryInfo.nots != nil {
		for _, notArray := range queryInfo.nots {
			if len(notArray) != 0 {
				if len(notArray) == 1 {
					cur = cur.Or(notArray[0])
				} else {
					cur = cur.Or(notArray[0], notArray[1:]...)
				}
			}
		}
	}
	if queryInfo.limit != -1 {
		cur = cur.Limit(queryInfo.limit)
	}
	if queryInfo.offset != -1 {
		cur = cur.Offset(queryInfo.offset)
	}
	if queryInfo.order != "" {
		cur = cur.Order(queryInfo.order)
	}
	return cur
}
