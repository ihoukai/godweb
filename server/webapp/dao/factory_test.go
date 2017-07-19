package dao

import (
	//"fmt"
	//"github.com/codegangsta/inject"
	. "gopkg.in/check.v1"
	//"reflect"
	//"server/webapp/dao/interfaces"
	//"server/webapp/model"
	//"server/webapp/dao/internal"
	//"strconv"
	"fmt"
	"server/webapp/dao/interfaces"
	"server/webapp/global/injector"
)

func NewTT(a string) {
	fmt.Printf(a)
}

type Hah struct {
	Name string "inject"
}

type My struct {
	Acc  interfaces.IAccountDao `inject`
	Sess interfaces.ISessionDao `inject`
}

func (s *TestSuite) TestGet(c *C) {
	injector.Map("hello world indayed\n")
	injector.Invoke(NewTT)

	injector.Apply(s)
	fmt.Printf("s=>%q\n", s)

	//injector.Map(factory.NewAccountDao())
	//injector.MapTo(factory.NewAccountDao(), (*interfaces.IAccountDao)(nil))

	var m My
	injector.Apply(&m)
	fmt.Printf("my=>%q\n", m)

	// a := Get((*interfaces.IAccountDao)(nil))
	// acc := a.(interfaces.IAccountDao)

	//username := "houkai"
	//username1 := "houkai1"
	// m.Acc.Create(&model.Account{
	// 	UserName: username,
	// 	Password: "123456",
	// })

	// account, err := acc.Auth(username, "123456")
	// if err != nil {
	// }
	// if account.UserName != username {
	// 	c.Errorf("Not Equal, Expect=%s Cur=%s", username, account.UserName)
	// }

	// acc.Create(&model.Account{
	// 	UserName: username1,
	// 	Password: "123456",
	// })

	// var users []model.Account
	// err := _baseDao.Find(&users, NewQueryInfo().Where("id != 0"))
	// if err != nil {
	// 	c.Error(err.Error())
	// }
	// for _, user := range users {
	// 	if user.UserName != username &&
	// 		user.UserName != username1 {
	// 		c.Error("Find Users List Error")
	// 	}
	// }
}

// func myTest(a, b interface{}) {
// 	typ := reflect.TypeOf(a)
// 	fmt.Printf("%s\n", typ)
// 	fmt.Printf("%s\n", b)

// 	t := reflect.TypeOf(a)
// 	for t.Kind() == reflect.Ptr {
// 		t = t.Elem()
// 	}
// 	fmt.Printf("%s\n", t)

// 	if typ == b {
// 		fmt.Printf("Equal\n")
// 	} else {
// 		fmt.Printf("Not Equal\n")
// 	}
// }
