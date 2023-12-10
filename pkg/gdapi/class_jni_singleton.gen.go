// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type JNISingleton struct {
  obj gdc.ObjectPtr
}

func (me *JNISingleton) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *JNISingleton) BaseClass() string {
  return "JNISingleton"
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

// Properties