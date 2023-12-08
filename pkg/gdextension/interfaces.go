package gdextension

import "github.com/LouisBrunner/go-dot-extension/pkg/gdc"

type Class interface {
	BaseClass() string
}

type ClassConstructor func() Class

type LogLevel uint32

func (me LogLevel) String() string {
	switch me {
	case LogLevelDebug:
		return "debug"
	case LogLevelInfo:
		return "info"
	case LogLevelWarning:
		return "warning"
	case LogLevelError:
		return "error"
	default:
		return "unknown"
	}
}

const (
	LogLevelError LogLevel = iota
	LogLevelWarning
	LogLevelInfo
	LogLevelDebug
)

type Extension interface {
	Register(constructors ...ClassConstructor)
	SetInitializationLevel(level gdc.InitializationLevel)
	Initialize(rInitialization *gdc.InitializationRaw, init gdc.InitializationInitializeFn, fini gdc.InitializationDeinitializeFn) gdc.Bool
	Logf(level LogLevel, format string, args ...interface{})
	LogDetailedf(level LogLevel, description, function, file string, line int32, notifyEditor bool, format string, args ...interface{})
}
