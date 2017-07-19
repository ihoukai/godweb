package interfaces

type IQueryInfo interface {
	Where(query interface{}, args ...interface{}) IQueryInfo
	Or(query interface{}, args ...interface{}) IQueryInfo
	Not(query interface{}, args ...interface{}) IQueryInfo
	Limit(limit int) IQueryInfo
	Offset(offset int) IQueryInfo
	Order(value string) IQueryInfo
}
