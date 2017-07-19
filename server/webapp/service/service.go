package service

func init() {
	// 反射创建所有的dao对象
	//log.Info("initFactory --- begin")
	// factory = internal.NewFactory(_db)
	// typ := reflect.TypeOf(factory)
	// //遍历方法
	// for i := 0; i < typ.NumMethod(); i++ {
	// 	method := typ.Method(i)
	// 	fmt.Printf("method => %s\n", method.Name)
	// 	fn := reflect.ValueOf(factory).MethodByName(method.Name)
	// 	if method.Type.NumIn() != 1 {
	// 		continue
	// 	}

	// 	ret := fn.Call(nil)
	// 	injector.Set(ret[0].Type(), ret[0])
	// 	fmt.Printf("injector.Set(%s, %s)\n", ret[0].Type(), ret[0])
	// }
	//log.Info("initFactory --- end")
}
