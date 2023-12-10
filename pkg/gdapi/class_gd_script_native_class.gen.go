// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

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

func (me *GDScriptNativeClass) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *GDScriptNativeClass) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *GDScriptNativeClass) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *GDScriptNativeClass) New() Variant {
  classNameV := StringNameFromStr("GDScriptNativeClass")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("new")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1460262497) // FIXME: should cache?
  var ret Variant
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

// Properties