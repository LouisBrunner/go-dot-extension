// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type SubViewportContainer struct {
  obj gdc.ObjectPtr
}

func (me *SubViewportContainer) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *SubViewportContainer) BaseClass() string {
  return "SubViewportContainer"
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
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enable), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *SubViewportContainer) IsStretchEnabled() bool {
  classNameV := StringNameFromStr("SubViewportContainer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_stretch_enabled")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *SubViewportContainer) SetStretchShrink(amount int, )  {
  classNameV := StringNameFromStr("SubViewportContainer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_stretch_shrink")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1286410249) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&amount), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *SubViewportContainer) GetStretchShrink() int {
  classNameV := StringNameFromStr("SubViewportContainer")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_stretch_shrink")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3905245786) // FIXME: should cache?
  var ret int
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

// Properties

func (me *SubViewportContainer) GetPropStretch() bool {
  panic("TODO: implement")
}

func (me *SubViewportContainer) SetPropStretch(value bool) {
  panic("TODO: implement")
}

func (me *SubViewportContainer) GetPropStretchShrink() int {
  panic("TODO: implement")
}

func (me *SubViewportContainer) SetPropStretchShrink(value int) {
  panic("TODO: implement")
}