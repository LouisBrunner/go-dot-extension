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

type ConvexPolygonShape2D struct {
  Shape2D
}

func (me *ConvexPolygonShape2D) BaseClass() string {
  return "ConvexPolygonShape2D"
}

func NewConvexPolygonShape2D() *ConvexPolygonShape2D {
  str := StringNameFromStr("ConvexPolygonShape2D") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &ConvexPolygonShape2D{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Enums

func (me *ConvexPolygonShape2D) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *ConvexPolygonShape2D) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *ConvexPolygonShape2D) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *ConvexPolygonShape2D) SetPointCloud(point_cloud PackedVector2Array, )  {
  classNameV := StringNameFromStr("ConvexPolygonShape2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_point_cloud")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1509147220) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{point_cloud.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *ConvexPolygonShape2D) SetPoints(points PackedVector2Array, )  {
  classNameV := StringNameFromStr("ConvexPolygonShape2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_points")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1509147220) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{points.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *ConvexPolygonShape2D) GetPoints() PackedVector2Array {
  classNameV := StringNameFromStr("ConvexPolygonShape2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_points")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2961356807) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewPackedVector2Array()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
