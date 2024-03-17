package gdextension

import (
	"reflect"
	"unsafe"
)

// FIXME: better error checking everywhere here

func reflectSetProperty(instance any, prop *classProperty, value any) {
	reflectSetField(instance, prop.name, value)
}

func reflectGetProperty(instance any, prop *classProperty) reflect.Value {
	return reflectGetField(instance, prop.name)
}

func reflectCallMethod(instance any, method *classMethod, args []reflect.Value) []reflect.Value {
	return reflect.ValueOf(method.fn).Call(append([]reflect.Value{reflect.ValueOf(instance)}, args...))
}

func reflectGetField(instance any, name string) reflect.Value {
	return reflect.ValueOf(instance).Elem().FieldByName(name)
}

func reflectSetField(instance any, name string, value any) {
	reflect.ValueOf(instance).Elem().FieldByName(name).Set(reflect.ValueOf(value))
}

func reflectUnsafeField(instance any, name string) reflect.Value {
	strct := reflect.ValueOf(instance).Elem()
	fld := strct.FieldByName(name)
	return reflect.NewAt(fld.Type(), unsafe.Pointer(fld.UnsafeAddr())).Elem()
}

func reflectGetFieldUnsafe(instance any, name string) reflect.Value {
	return reflectUnsafeField(instance, name)
}

func reflectSetFieldUnsafe(instance any, name string, value any) {
	reflectUnsafeField(instance, name).Set(reflect.ValueOf(value))
}

func reflectIsA[T any](typ reflect.Type) bool {
	_, is := reflect.New(typ).Elem().Interface().(T)
	return is
}
