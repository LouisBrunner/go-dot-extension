// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type RefCounted struct {
  obj gdc.ObjectPtr
}

func (me *RefCounted) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *RefCounted) BaseClass() string {
  return "RefCounted"
}



// Enums

func (me *RefCounted) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *RefCounted) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *RefCounted) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *RefCounted) InitRef() bool {
  classNameV := StringNameFromStr("RefCounted")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("init_ref")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2240911060) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *RefCounted) Reference() bool {
  classNameV := StringNameFromStr("RefCounted")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("reference")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2240911060) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *RefCounted) Unreference() bool {
  classNameV := StringNameFromStr("RefCounted")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("unreference")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2240911060) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *RefCounted) GetReferenceCount() int {
  classNameV := StringNameFromStr("RefCounted")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_reference_count")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

// Signals
