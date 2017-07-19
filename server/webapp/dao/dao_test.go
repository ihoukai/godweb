package dao

import (
	"fmt"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	. "gopkg.in/check.v1"
	"server/webapp/model"
	"testing"
)

var _ = Suite(&TestSuite{})

func Test(t *testing.T) { TestingT(t) }

type TestSuite struct {
	Name string "inject"
}

// SetUpSuite 框架方法(测试用例执行之前调用)
func (s *TestSuite) SetUpSuite(c *C) {
	tables := []interface{}{&model.CfgDailySignIn{},
		&model.CfgRoom{},
		&model.CfgShop{},
		&model.LogChips{},
		&model.LogDiamond{},
		&model.LogShop{},
		&model.Notice{},
		&model.UserStatistics{},
		&model.UserInfoDetail{},
		&model.Account{},
		&model.Session{},
		&model.UserFriend{}}
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
	err := Connect()
	if err != nil {
		panic(fmt.Sprintf("err:%s\n", err.Error()))
	}
	_db.LogMode(true)
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
