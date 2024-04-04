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

type Path2D struct {
  Node2D
}

func (me *Path2D) BaseClass() string {
  return "Path2D"
}

func NewPath2D() *Path2D {
  str := StringNameFromStr("Path2D") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &Path2D{}
  obj.SetBaseObject(objPtr)
  return obj
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
  cargs := []gdc.ConstTypePtr{curve.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Path2D) GetCurve() Curve2D {
  classNameV := StringNameFromStr("Path2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_curve")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 660369445) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewCurve2D()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
