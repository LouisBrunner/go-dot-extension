// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

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

func (me *JNISingleton) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *JNISingleton) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

// TODO: properties (class)

// TODO: signals (class)
