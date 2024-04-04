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

type PolygonPathFinder struct {
  Resource
}

func (me *PolygonPathFinder) BaseClass() string {
  return "PolygonPathFinder"
}

func NewPolygonPathFinder() *PolygonPathFinder {
  str := StringNameFromStr("PolygonPathFinder") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &PolygonPathFinder{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Enums

func (me *PolygonPathFinder) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *PolygonPathFinder) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *PolygonPathFinder) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *PolygonPathFinder) Setup(points PackedVector2Array, connections PackedInt32Array, )  {
  classNameV := StringNameFromStr("PolygonPathFinder")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("setup")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3251786936) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{points.AsCTypePtr(), connections.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *PolygonPathFinder) FindPath(from Vector2, to Vector2, ) PackedVector2Array {
  classNameV := StringNameFromStr("PolygonPathFinder")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("find_path")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1562168077) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{from.AsCTypePtr(), to.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewPackedVector2Array()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *PolygonPathFinder) GetIntersections(from Vector2, to Vector2, ) PackedVector2Array {
  classNameV := StringNameFromStr("PolygonPathFinder")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_intersections")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3932192302) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{from.AsCTypePtr(), to.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewPackedVector2Array()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *PolygonPathFinder) GetClosestPoint(point Vector2, ) Vector2 {
  classNameV := StringNameFromStr("PolygonPathFinder")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_closest_point")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2656412154) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{point.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewVector2()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *PolygonPathFinder) IsPointInside(point Vector2, ) bool {
  classNameV := StringNameFromStr("PolygonPathFinder")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("is_point_inside")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 556197845) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{point.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *PolygonPathFinder) SetPointPenalty(idx int64, penalty float64, )  {
  classNameV := StringNameFromStr("PolygonPathFinder")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_point_penalty")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1602489585) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&idx) , gdc.ConstTypePtr(&penalty) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *PolygonPathFinder) GetPointPenalty(idx int64, ) float64 {
  classNameV := StringNameFromStr("PolygonPathFinder")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_point_penalty")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2339986948) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&idx) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()
  pinner.Pin(&idx)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *PolygonPathFinder) GetBounds() Rect2 {
  classNameV := StringNameFromStr("PolygonPathFinder")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_bounds")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1639390495) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewRect2()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
