package gdextension

import (
	"github.com/LouisBrunner/go-dot-extension/pkg/gdapi"
	"github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type Class interface {
	BaseClass() string
	SetBaseObject(obj gdc.ObjectPtr)
	Get(name gdapi.StringName) gdapi.Variant
	Set(name gdapi.StringName, value gdapi.Variant)
}

type Destroyable interface {
	Destroy()
}

type ClassConstructor func() Class

type Extension interface {
	Register(constructors ...ClassConstructor)
	SetInitializationLevel(level gdc.InitializationLevel)
	Initialize(rInitialization *gdc.InitializationRaw, init gdc.InitializationInitializeFn, fini gdc.InitializationDeinitializeFn) gdc.Bool
	Logf(level LogLevel, format string, args ...interface{})
	LogDetailedf(level LogLevel, description, function, file string, line int32, notifyEditor bool, format string, args ...interface{})
}
