// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "log"
  "runtime"
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

// FIXME: avoid unused import warning
var _ log.Logger
var _ unsafe.Pointer
var _ runtime.Pinner

type ptrsForPath2DList struct {
  fnSetCurve gdc.MethodBindPtr
  fnGetCurve gdc.MethodBindPtr
}

var ptrsForPath2D ptrsForPath2DList

func initPath2DPtrs(iface gdc.Interface) {

  className := StringNameFromStr("Path2D")
  defer className.Destroy()
  {
    methodName := StringNameFromStr("set_curve")
    defer methodName.Destroy()
    ptrsForPath2D.fnSetCurve = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 659985499))
  }
  {
    methodName := StringNameFromStr("get_curve")
    defer methodName.Destroy()
    ptrsForPath2D.fnGetCurve = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 660369445))
  }
}

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
  cargs := []gdc.ConstTypePtr{curve.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPath2D.fnSetCurve), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *Path2D) GetCurve() Curve2D {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewCurve2D()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPath2D.fnGetCurve), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
