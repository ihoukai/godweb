package impl

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"server/webapp/global/errors"
	"server/webapp/global/injector"
	"server/webapp/service/interfaces"
)

// RouterAccount 路由
func RouterAccount(router *gin.RouterGroup) {
	router.POST("user/login", getAccount().login)
	router.GET("user/register", getAccount().register)
}

func getAccount() *accountCtl {
	if account == nil {
		account = &accountCtl{}
		injector.Apply(account)
		fmt.Printf("=>%q\n", account)
	}
	return account
}

var account *accountCtl

type accountCtl struct {
	AccountSrv interfaces.IAccountSrv `inject:"t"`
}

// 用户参数
type accountParam struct {
	UserName string `form:"username" json:"username" binding:"required" `
	Password string `form:"password" json:"password" binding:"required" validate:"number,len=6"`
	Email    string `form:"email" json:"email" validate:"email"`
}

func (u *accountCtl) login(c *gin.Context) {
	// todo 数据校验
	var json accountParam
	err := c.BindJSON(&json)
	if err != nil {
		httpResponse(c, nil, errors.New(errors.ParamsError, err))
		return
	}
	data, e := u.AccountSrv.Login(json.UserName, json.Password)
	httpResponse(c, data, e)
}

func (u *accountCtl) register(c *gin.Context) {

}
