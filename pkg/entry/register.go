package entry

import "github.com/LouisBrunner/go-dot-extension/pkg/gdextension"

var globalRegisteredClasses []gdextension.ClassConstructor

func Register(constructors ...gdextension.ClassConstructor) {
	globalRegisteredClasses = append(globalRegisteredClasses, constructors...)
}
