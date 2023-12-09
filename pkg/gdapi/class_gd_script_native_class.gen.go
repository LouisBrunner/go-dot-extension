// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type GDScriptNativeClass struct {
  obj gdc.ObjectPtr
}

func (me *GDScriptNativeClass) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *GDScriptNativeClass) BaseClass() string {
  return "GDScriptNativeClass"
}



// Enums

func (me *GDScriptNativeClass) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *GDScriptNativeClass) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *GDScriptNativeClass) New()  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
