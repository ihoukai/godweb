package injector

import (
	"github.com/codegangsta/inject"
	"reflect"
	"server/webapp/global/log"
)

var injector = inject.New()

// SetFactory 设置创建对象的工厂方法，根据工厂提供的方法注入对象
func SetFactory(ifactory interface{}) {
	values, err := injector.Invoke(ifactory)
	if err != nil {
		panic(err)
	}
	factory := values[0].Interface()
	typ := reflect.TypeOf(factory)
	//遍历方法
	for i := 0; i < typ.NumMethod(); i++ {
		method := typ.Method(i)
		log.Info("method => %s", method.Name)
		fn := reflect.ValueOf(factory).MethodByName(method.Name)
		if method.Type.NumIn() != 1 {
			continue
		}

		ret := fn.Call(nil)
		// 注入inject字段
		Apply(ret[0].Interface())
		Set(ret[0].Type(), ret[0])
	}
}

// Apply ：Maps dependencies in the Type map to each field in the struct
// that is tagged with 'inject'. Returns an error if the injection
// fails.
func Apply(value interface{}) error {
	return injector.Apply(value)
}

// Invoke : attempts to call the interface{} provided as a function,
// providing dependencies for function arguments based on Type. Returns
// a slice of reflect.Value representing the returned values of the function.
// Returns an error if the injection fails.
func Invoke(value interface{}) ([]reflect.Value, error) {
	return injector.Invoke(value)
}

// Map : Maps the interface{} value based on its immediate type from reflect.TypeOf.
func Map(val interface{}) {
	vt := reflect.TypeOf(val)
	if vt.Kind() == reflect.Ptr {
		vt = vt.Elem()
	}
	if vt.Kind() == reflect.Interface {
		vt = vt.Elem()
	}
	log.Info("inject => %s %s", reflect.TypeOf(val).String(), vt.String())
	injector.Map(val)
}

// MapTo : Maps the interface{} value based on the pointer of an Interface provided.
// This is really only useful for mapping a value as an interface, as interfaces
// cannot at this time be referenced directly without a pointer.
func MapTo(val interface{}, ifacePtr interface{}) {
	vt := reflect.TypeOf(val)
	if vt.Kind() == reflect.Ptr {
		vt = vt.Elem()
	}
	if vt.Kind() == reflect.Interface {
		vt = vt.Elem()
	}
	log.Info("inject => %s %s", reflect.TypeOf(ifacePtr).String(), vt.String())
	injector.MapTo(val, ifacePtr)
}

// Set : Provides a possibility to directly insert a mapping based on type and value.
// This makes it possible to directly map type arguments not possible to instantiate
// with reflect like unidirectional channels.
func Set(typ reflect.Type, value reflect.Value) {
	vt := value
	if vt.Kind() == reflect.Ptr {
		vt = vt.Elem()
	}
	if vt.Kind() == reflect.Interface {
		vt = vt.Elem()
	}
	log.Info("inject => %s %s", typ.String(), vt.String())
	injector.Set(typ, value)
}

// Get : Returns the Value that is mapped to the current type. Returns a zeroed Value if
// the Type has not been mapped.
func Get(typ reflect.Type) reflect.Value {
	return injector.Get(typ)
}
