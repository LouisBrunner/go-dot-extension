// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"
  "runtime"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

// FIXME: avoid unused import warning
var _ unsafe.Pointer
var _ runtime.Pinner

type SubViewportContainer struct {
  Container
}

func (me *SubViewportContainer) BaseClass() string {
  return "SubViewportContainer"
}

func NewSubViewportContainer() *SubViewportContainer {
  str := StringNameFromStr("SubViewportContainer") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &SubViewportContainer{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Enums

func (me *SubViewportContainer) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *SubViewportContainer) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *SubViewportContainer) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *SubViewportContainer) SetStretch(enable bool, )  {
  classNameV := StringNameFromStr("SubViewportContainer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_stretch")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *SubViewportContainer) IsStretchEnabled() bool {
  classNameV := StringNameFromStr("SubViewportContainer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_stretch_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *SubViewportContainer) SetStretchShrink(amount int64, )  {
  classNameV := StringNameFromStr("SubViewportContainer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_stretch_shrink")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&amount) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *SubViewportContainer) GetStretchShrink() int64 {
  classNameV := StringNameFromStr("SubViewportContainer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_stretch_shrink")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewInt()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
