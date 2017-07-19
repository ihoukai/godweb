package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"server/webapp/controller/impl"
)

// Init 初始化控制层
func Init() http.Handler {
	router := gin.Default()
	router.Use(gin.Logger())
	//group
	v1 := router.Group("api/v1")
	{
		registerRouter(v1)
	}
	return router
}

func registerRouter(router *gin.RouterGroup) {
	impl.RouterAccount(router)
}
