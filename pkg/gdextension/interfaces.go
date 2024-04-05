package gdextension

import (
	"reflect"
	"unsafe"

	"github.com/LouisBrunner/go-dot-extension/pkg/gdapi"
	"github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type ConstantsDefinition struct {
	Constants map[string]int
	Enums     map[string]map[string]int
}

type Class interface {
	BaseClass() string
	SetBaseObject(obj gdc.ObjectPtr)
	Get(name gdapi.StringName) gdapi.Variant
	Set(name gdapi.StringName, value gdapi.Variant)
}

type ClassWithConstants interface {
	Class
	Constants() ConstantsDefinition
}

type Initializable interface {
	Class
	Init()
}

type Destroyable interface {
	Class
	Destroy()
}

type ClassConstructor func() Class

type Extension interface {
	Register(constructors ...ClassConstructor)
	SetInitializationLevel(level gdc.InitializationLevel)
	Initialize(rInitialization *gdc.InitializationRaw, init gdc.InitializationInitializeFn, fini gdc.InitializationDeinitializeFn) gdc.Bool
	Deinitialize(userdata unsafe.Pointer, level gdc.InitializationLevel)
	Logf(level LogLevel, format string, args ...interface{})
	LogDetailedf(level LogLevel, description, function, file string, line int32, notifyEditor bool, format string, args ...interface{})
	CreateClass(typ reflect.Type) (any, error)
}
