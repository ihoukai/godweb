package interfaces

type IBaseDao interface {
	Create(obj interface{}) (err error)
	Delete(id uint) (err error)
	Update(obj interface{}) (err error)
	Find(out interface{}, q IQueryInfo) (err error)
	Count(table interface{}, out interface{}, q IQueryInfo) (err error)
}
