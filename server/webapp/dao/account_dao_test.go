package dao

// import (
// 	"fmt"
// 	_ "github.com/jinzhu/gorm/dialects/postgres"
// 	. "gopkg.in/check.v1"
// 	"server/webapp/global/errors"
// 	"server/webapp/model"
// )

// // TestAccount 测试用户创建及查询
// func (s *TestSuite) TestAccount(c *C) {
// 	username := "houkai"
// 	username1 := "houkai1"
// 	Account().Create(&model.Account{
// 		UserName: username,
// 		Password: "123456",
// 	})

// 	account, err := Account().Auth(username, "123456")
// 	if err != nil {
// 		c.Error(retCode.String())
// 	}
// 	if account.UserName != username {
// 		c.Errorf("Not Equal, Expect=%s Cur=%s", username, account.UserName)
// 	}

// 	Account().Create(&model.Account{
// 		UserName: username1,
// 		Password: "123456",
// 	})

// 	var users []model.Account
// 	err := _baseDao.Find(&users, NewQueryInfo().Where("id != 0"))
// 	if err != nil {
// 		c.Error(err.Error())
// 	}
// 	for _, user := range users {
// 		if user.UserName != username &&
// 			user.UserName != username1 {
// 			c.Error("Find Users List Error")
// 		}
// 	}
// }

// func (s *TestSuite) TestGetFriends(c *C) {
// 	userInfo := &model.UserInfoDetail{}
// 	userInfo.Diamond = 1000
// 	userInfo.Chips = 1000
// 	userInfo.UserName = "houkai"
// 	userInfo.AccountID = 1000
// 	err := _db.Create(userInfo).Error
// 	if err != nil {
// 		c.Error(err)
// 	}

// 	bui := &model.UserInfo{}
// 	err = _db.Where("user_name = ?", "houkai").Find(bui).Error
// 	if err != nil {
// 		c.Error(err)
// 	}
// }

// func (s *TestSuite) TestAccountCount(c *C) {
// 	username := "houkai2"
// 	Account().Create(&model.Account{
// 		UserName: username,
// 		Password: "123456",
// 		Email:    "houkai@163.com",
// 	})

// 	count, _ := Account().Count(&model.Account{
// 		UserName: username,
// 	})
// 	if count != 1 {
// 		c.Error("TestAccountCount Error")
// 	}
// 	fmt.Printf("%d\n", count)
// }
