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

type WorldBoundaryShape2D struct {
  Shape2D
}

func (me *WorldBoundaryShape2D) BaseClass() string {
  return "WorldBoundaryShape2D"
}

func NewWorldBoundaryShape2D() *WorldBoundaryShape2D {
  str := StringNameFromStr("WorldBoundaryShape2D") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &WorldBoundaryShape2D{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Enums

func (me *WorldBoundaryShape2D) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *WorldBoundaryShape2D) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *WorldBoundaryShape2D) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *WorldBoundaryShape2D) SetNormal(normal Vector2, )  {
  classNameV := StringNameFromStr("WorldBoundaryShape2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_normal")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 743155724) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{normal.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *WorldBoundaryShape2D) GetNormal() Vector2 {
  classNameV := StringNameFromStr("WorldBoundaryShape2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_normal")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3341600327) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewVector2()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *WorldBoundaryShape2D) SetDistance(distance float64, )  {
  classNameV := StringNameFromStr("WorldBoundaryShape2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_distance")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&distance) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *WorldBoundaryShape2D) GetDistance() float64 {
  classNameV := StringNameFromStr("WorldBoundaryShape2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_distance")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
