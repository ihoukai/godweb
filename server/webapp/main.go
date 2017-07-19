package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"reflect"
	"server/webapp/controller"
	"server/webapp/dao"
	daoImpl "server/webapp/dao/impl"
	"server/webapp/global/injector"
	"server/webapp/log"
	"server/webapp/model"
	srvImpl "server/webapp/service/impl"
	"time"
)

func main() {
	initDB()
	injectDao()
	injectSrv()
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

func injectDao() {
	// 反射创建所有的dao对象
	//log.Info("initFactory --- begin")
	values, err := injector.Invoke(daoImpl.NewFactory)
	if err != nil {
		panic(err)
	}
	factory := values[0].Interface().(*daoImpl.Factory)
	typ := reflect.TypeOf(factory)
	//遍历方法
	for i := 0; i < typ.NumMethod(); i++ {
		method := typ.Method(i)
		fmt.Printf("method => %s\n", method.Name)
		fn := reflect.ValueOf(factory).MethodByName(method.Name)
		if method.Type.NumIn() != 1 {
			continue
		}

		ret := fn.Call(nil)
		// 注入inject字段
		injector.Apply(ret[0].Interface())
		injector.Set(ret[0].Type(), ret[0])
	}
	//log.Info("initFactory --- end")
}

func injectSrv() {
	values, err := injector.Invoke(srvImpl.NewFactory)
	if err != nil {
		panic(err)
	}
	factory := values[0].Interface().(*srvImpl.Factory)
	typ := reflect.TypeOf(factory)
	//遍历方法
	for i := 0; i < typ.NumMethod(); i++ {
		method := typ.Method(i)
		fmt.Printf("method => %s\n", method.Name)
		fn := reflect.ValueOf(factory).MethodByName(method.Name)
		if method.Type.NumIn() != 1 {
			continue
		}

		ret := fn.Call(nil)
		// 注入inject字段
		injector.Apply(ret[0].Interface())
		injector.Set(ret[0].Type(), ret[0])
	}
}
