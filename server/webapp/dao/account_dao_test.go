package dao

import (
	"fmt"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	. "gopkg.in/check.v1"
	"server/webapp/model"
)

// TestAccount 测试用户创建及查询
func (s *TestSuite) TestAccount(c *C) {
	username := "houkai"
	username1 := "houkai1"
	s.AccountDao.Create(&model.Account{
		UserName: username,
		Password: "123456",
	})

	account, err := s.AccountDao.Auth(username, "123456")
	if err != nil {
		c.Error(err)
	}
	if account.UserName != username {
		c.Errorf("Not Equal, Expect=%s Cur=%s", username, account.UserName)
	}

	s.AccountDao.Create(&model.Account{
		UserName: username1,
		Password: "123456",
	})

	var users []model.Account
	e := _db.Where("id != 0").Find(&users).Error
	if e != nil {
		c.Error(e)
	}
	for _, user := range users {
		if user.UserName != username &&
			user.UserName != username1 {
			c.Error("Find Users List Error")
		}
	}
}

func (s *TestSuite) TestAccountCount(c *C) {
	username := "houkai2"
	s.AccountDao.Create(&model.Account{
		UserName: username,
		Password: "123456",
		Email:    "houkai@163.com",
	})

	count, _ := s.AccountDao.Count(&model.Account{
		UserName: username,
	})
	if count != 1 {
		c.Error("TestAccountCount Error")
	}
	fmt.Printf("%d\n", count)
}
