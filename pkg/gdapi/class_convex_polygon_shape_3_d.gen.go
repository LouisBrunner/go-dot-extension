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

type ConvexPolygonShape3D struct {
  Shape3D
}

func (me *ConvexPolygonShape3D) BaseClass() string {
  return "ConvexPolygonShape3D"
}

func NewConvexPolygonShape3D() *ConvexPolygonShape3D {
  str := StringNameFromStr("ConvexPolygonShape3D") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &ConvexPolygonShape3D{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Enums

func (me *ConvexPolygonShape3D) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *ConvexPolygonShape3D) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *ConvexPolygonShape3D) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *ConvexPolygonShape3D) SetPoints(points PackedVector3Array, )  {
  classNameV := StringNameFromStr("ConvexPolygonShape3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_points")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 334873810) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{points.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *ConvexPolygonShape3D) GetPoints() PackedVector3Array {
  classNameV := StringNameFromStr("ConvexPolygonShape3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_points")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 497664490) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewPackedVector3Array()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
