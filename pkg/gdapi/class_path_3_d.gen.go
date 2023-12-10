// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type Path3D struct {
  obj gdc.ObjectPtr
}

func (me *Path3D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *Path3D) BaseClass() string {
  return "Path3D"
}



// Enums

func (me *Path3D) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *Path3D) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *Path3D) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *Path3D) SetCurve(curve Curve3D, )  {
  classNameV := StringNameFromStr("Path3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_curve")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 408955118) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(curve.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Path3D) GetCurve() Curve3D {
  classNameV := StringNameFromStr("Path3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_curve")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 4244715212) // FIXME: should cache?
  var ret Curve3D
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

// Properties

func (me *Path3D) GetPropCurve() Curve3D {
  panic("TODO: implement")
}

func (me *Path3D) SetPropCurve(value Curve3D) {
  panic("TODO: implement")
}
// Signals
// FIXME: can't seem to be able to connect them from this side of the API