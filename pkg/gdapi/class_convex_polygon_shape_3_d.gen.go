// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type ConvexPolygonShape3D struct {
  obj gdc.ObjectPtr
}

func (me *ConvexPolygonShape3D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *ConvexPolygonShape3D) BaseClass() string {
  return "ConvexPolygonShape3D"
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
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(points.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *ConvexPolygonShape3D) GetPoints() PackedVector3Array {
  classNameV := StringNameFromStr("ConvexPolygonShape3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_points")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 497664490) // FIXME: should cache?
  var ret PackedVector3Array
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
