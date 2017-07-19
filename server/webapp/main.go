package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"server/webapp/controller"
	"server/webapp/dao"
	daoImpl "server/webapp/dao/impl"
	"server/webapp/global/injector"
	"server/webapp/global/log"
	"server/webapp/model"
	srvImpl "server/webapp/service/impl"
	"time"
)

func main() {
	initDB()
	injector.SetFactory(daoImpl.NewFactory)
	injector.SetFactory(srvImpl.NewFactory)
	run(controller.Init())
}

func run(handler http.Handler) {
	// todo 从配置文件中加载配置数据
	srv := &http.Server{
		Addr:    ":8080",
		Handler: handler,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Error("listen: %s\n", err)
		}
	}()

	// 等待中断信号 5秒后安全的关闭服务
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Info("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Error("Server Shutdown:%s", err.Error())
	}
	log.Info("Server exist")
}

func initDB() {
	// todo 从配置文件中加载配置数据
	tables := []interface{}{
		&model.Account{},
		&model.Session{}}
	dao.Init(&dao.Config{
		DB:       "postgres",
		Host:     "localhost",
		User:     "gorm",
		DBName:   "gorm-test",
		PassWord: "123456",
		IDZones: []*dao.IDZone{
			&dao.IDZone{
				StartID: 100000,
				EndID:   999999,
			},
			&dao.IDZone{
				StartID: 200000,
				EndID:   299999,
				CurID:   299999,
			},
		},
	}, tables...)
	db, err := dao.Connect()
	if err != nil {
		panic(fmt.Sprintf("err:%s\n", err.Error()))
	}
	db.LogMode(true)
	// 注入数据库
	injector.Map(db)
}
