package impl

import (
	"server/webapp/dao/interfaces"
)

// QueryInfo 查询信息
type QueryInfo struct {
	wheres [][]interface{}
	ors    [][]interface{}
	nots   [][]interface{}
	limit  int
	offset int
	order  string
}

func NewQueryInfo() interfaces.IQueryInfo {
	q := &QueryInfo{
		wheres: make([][]interface{}, 0),
		ors:    make([][]interface{}, 0),
		nots:   make([][]interface{}, 0),
		limit:  -1,
		offset: -1,
		order:  "",
	}
	return q
}

func (q *QueryInfo) Where(query interface{}, args ...interface{}) interfaces.IQueryInfo {
	params := make([]interface{}, 0)
	params = append(params, query)
	if args != nil {
		params = append(params, args...)
	}
	q.wheres = append(q.wheres, params)
	return q
}

func (q *QueryInfo) Or(query interface{}, args ...interface{}) interfaces.IQueryInfo {
	params := make([]interface{}, 0)
	params = append(params, query)
	if args != nil {
		params = append(params, args...)
	}
	q.ors = append(q.ors, params)
	return q
}

func (q *QueryInfo) Not(query interface{}, args ...interface{}) interfaces.IQueryInfo {
	params := make([]interface{}, 0)
	params = append(params, query)
	if args != nil {
		params = append(params, args...)
	}
	q.nots = append(q.nots, params)
	return q
}

func (q *QueryInfo) Limit(limit int) interfaces.IQueryInfo {
	q.limit = limit
	return q
}

func (q *QueryInfo) Offset(offset int) interfaces.IQueryInfo {
	q.offset = offset
	return q
}

func (q *QueryInfo) Order(value string) interfaces.IQueryInfo {
	q.order = value
	return q
}
