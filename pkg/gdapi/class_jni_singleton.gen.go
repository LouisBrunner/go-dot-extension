// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "log"
  "runtime"
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

// FIXME: avoid unused import warning
var _ log.Logger
var _ unsafe.Pointer
var _ runtime.Pinner

type ptrsForJNISingletonList struct {
}

var ptrsForJNISingleton ptrsForJNISingletonList

func initJNISingletonPtrs(iface gdc.Interface) {

  className := StringNameFromStr("JNISingleton")
  defer className.Destroy()
}

type JNISingleton struct {
  Object
}

func (me *JNISingleton) BaseClass() string {
  return "JNISingleton"
}

func NewJNISingleton() *JNISingleton {
  str := StringNameFromStr("JNISingleton") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &JNISingleton{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Enums

func (me *JNISingleton) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *JNISingleton) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *JNISingleton) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

// Signals
