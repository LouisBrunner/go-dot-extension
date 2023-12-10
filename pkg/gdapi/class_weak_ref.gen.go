// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type WeakRef struct {
  obj gdc.ObjectPtr
}

func (me *WeakRef) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *WeakRef) BaseClass() string {
  return "WeakRef"
}



// Enums

func (me *WeakRef) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *WeakRef) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *WeakRef) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *WeakRef) GetRef() Variant {
  classNameV := StringNameFromStr("WeakRef")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_ref")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1214101251) // FIXME: should cache?
  var ret Variant
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

// Properties