// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type XROrigin3D struct {
  obj gdc.ObjectPtr
}

func (me *XROrigin3D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *XROrigin3D) BaseClass() string {
  return "XROrigin3D"
}



// Enums

func (me *XROrigin3D) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *XROrigin3D) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *XROrigin3D) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *XROrigin3D) SetWorldScale(world_scale float32, )  {
  classNameV := StringNameFromStr("XROrigin3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_world_scale")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&world_scale), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *XROrigin3D) GetWorldScale() float32 {
  classNameV := StringNameFromStr("XROrigin3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_world_scale")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *XROrigin3D) SetCurrent(enabled bool, )  {
  classNameV := StringNameFromStr("XROrigin3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_current")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2586408642) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&enabled), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *XROrigin3D) IsCurrent() bool {
  classNameV := StringNameFromStr("XROrigin3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_current")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 36873697) // FIXME: should cache?
  var ret bool
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

// Properties

func (me *XROrigin3D) GetPropWorldScale() float32 {
  panic("TODO: implement")
}

func (me *XROrigin3D) SetPropWorldScale(value float32) {
  panic("TODO: implement")
}

func (me *XROrigin3D) GetPropCurrent() bool {
  panic("TODO: implement")
}

func (me *XROrigin3D) SetPropCurrent(value bool) {
  panic("TODO: implement")
}