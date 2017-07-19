package dao

import (
	"fmt"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	. "gopkg.in/check.v1"
	daoImpl "server/webapp/dao/impl"
	. "server/webapp/dao/interfaces"
	"server/webapp/global/injector"
	"server/webapp/model"
	"testing"
)

var _ = Suite(&TestSuite{})

func Test(t *testing.T) { TestingT(t) }

type TestSuite struct {
	AccountDao IAccountDao `inject:"t"`
}

// SetUpSuite 框架方法(测试用例执行之前调用)
func (s *TestSuite) SetUpSuite(c *C) {
	tables := []interface{}{
		&model.Account{},
		&model.Session{}}
	Init(&Config{
		DB:       "postgres",
		Host:     "localhost",
		User:     "gorm",
		DBName:   "gorm-test",
		PassWord: "123456",
		IDZones: []*IDZone{
			&IDZone{
				StartID: 100000,
				EndID:   999999,
			},
			&IDZone{
				StartID: 200000,
				EndID:   299999,
				CurID:   299999,
			},
		},
	}, tables...)
	db, err := Connect()
	if err != nil {
		panic(fmt.Sprintf("err:%s\n", err.Error()))
	}
	db.LogMode(true)
	injector.Map(db)
	injector.SetFactory(daoImpl.NewFactory)
	injector.Apply(s)
}

// TearDownSuite 框架方法(所有测试用例执行之后调用)
func (s *TestSuite) TearDownSuite(c *C) {
	// 删除创建的表
	for _, table := range _tables {
		_db.DropTableIfExists(table)
	}
	DisConnect()
}

// SetUpTest 框架方法(每个测试用例执行之前都会调用)
func (s *TestSuite) SetUpTest(c *C) {
}

// TearDownTest 框架方法(每个测试用例执行之后都会调用)
func (s *TestSuite) TearDownTest(c *C) {
}
