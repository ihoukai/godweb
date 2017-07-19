package dao

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"reflect"
	"server/webapp/dao/impl"
	"server/webapp/global/injector"
	//"server/webapp/log"
)

var factory *impl.Factory

func initFactory(_db *gorm.DB) {
	// 反射创建所有的dao对象
	//log.Info("initFactory --- begin")
	factory = impl.NewFactory(_db)
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
		injector.Set(ret[0].Type(), ret[0])
	}
	//log.Info("initFactory --- end")
}

func Get(inter interface{}) interface{} {
	t := reflect.TypeOf(inter)
	for t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	fmt.Printf("%s\n", t)
	return injector.Get(t).Interface()
}
