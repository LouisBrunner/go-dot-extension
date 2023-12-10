// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type ConvexPolygonShape2D struct {
  obj gdc.ObjectPtr
}

func (me *ConvexPolygonShape2D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *ConvexPolygonShape2D) BaseClass() string {
  return "ConvexPolygonShape2D"
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
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(point_cloud.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *ConvexPolygonShape2D) SetPoints(points PackedVector2Array, )  {
  classNameV := StringNameFromStr("ConvexPolygonShape2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_points")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1509147220) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(points.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *ConvexPolygonShape2D) GetPoints() PackedVector2Array {
  classNameV := StringNameFromStr("ConvexPolygonShape2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_points")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2961356807) // FIXME: should cache?
  var ret PackedVector2Array
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

// Properties

func (me *ConvexPolygonShape2D) GetPropPoints() PackedVector2Array {
  panic("TODO: implement")
}

func (me *ConvexPolygonShape2D) SetPropPoints(value PackedVector2Array) {
  panic("TODO: implement")
}