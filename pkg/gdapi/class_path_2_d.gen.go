// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type Path2D struct {
  obj gdc.ObjectPtr
}

func (me *Path2D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *Path2D) BaseClass() string {
  return "Path2D"
}



// Enums

func (me *Path2D) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *Path2D) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *Path2D) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *Path2D) SetCurve(curve Curve2D, )  {
  classNameV := StringNameFromStr("Path2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_curve")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 659985499) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(curve.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *Path2D) GetCurve() Curve2D {
  classNameV := StringNameFromStr("Path2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_curve")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 660369445) // FIXME: should cache?
  var ret Curve2D
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

// Properties

func (me *Path2D) GetPropCurve() Curve2D {
  panic("TODO: implement")
}

func (me *Path2D) SetPropCurve(value Curve2D) {
  panic("TODO: implement")
}