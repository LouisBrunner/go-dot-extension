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

type ptrsForPolygonPathFinderList struct {
  fnSetup gdc.MethodBindPtr
  fnFindPath gdc.MethodBindPtr
  fnGetIntersections gdc.MethodBindPtr
  fnGetClosestPoint gdc.MethodBindPtr
  fnIsPointInside gdc.MethodBindPtr
  fnSetPointPenalty gdc.MethodBindPtr
  fnGetPointPenalty gdc.MethodBindPtr
  fnGetBounds gdc.MethodBindPtr
}

var ptrsForPolygonPathFinder ptrsForPolygonPathFinderList

func initPolygonPathFinderPtrs(iface gdc.Interface) {

  className := StringNameFromStr("PolygonPathFinder")
  defer className.Destroy()
  {
    methodName := StringNameFromStr("setup")
    defer methodName.Destroy()
    ptrsForPolygonPathFinder.fnSetup = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3251786936))
  }
  {
    methodName := StringNameFromStr("find_path")
    defer methodName.Destroy()
    ptrsForPolygonPathFinder.fnFindPath = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1562168077))
  }
  {
    methodName := StringNameFromStr("get_intersections")
    defer methodName.Destroy()
    ptrsForPolygonPathFinder.fnGetIntersections = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3932192302))
  }
  {
    methodName := StringNameFromStr("get_closest_point")
    defer methodName.Destroy()
    ptrsForPolygonPathFinder.fnGetClosestPoint = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2656412154))
  }
  {
    methodName := StringNameFromStr("is_point_inside")
    defer methodName.Destroy()
    ptrsForPolygonPathFinder.fnIsPointInside = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 556197845))
  }
  {
    methodName := StringNameFromStr("set_point_penalty")
    defer methodName.Destroy()
    ptrsForPolygonPathFinder.fnSetPointPenalty = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1602489585))
  }
  {
    methodName := StringNameFromStr("get_point_penalty")
    defer methodName.Destroy()
    ptrsForPolygonPathFinder.fnGetPointPenalty = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 2339986948))
  }
  {
    methodName := StringNameFromStr("get_bounds")
    defer methodName.Destroy()
    ptrsForPolygonPathFinder.fnGetBounds = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1639390495))
  }
}

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
  cargs := []gdc.ConstTypePtr{points.AsCTypePtr(), connections.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPolygonPathFinder.fnSetup), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *PolygonPathFinder) FindPath(from Vector2, to Vector2, ) PackedVector2Array {
  cargs := []gdc.ConstTypePtr{from.AsCTypePtr(), to.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewPackedVector2Array()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPolygonPathFinder.fnFindPath), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *PolygonPathFinder) GetIntersections(from Vector2, to Vector2, ) PackedVector2Array {
  cargs := []gdc.ConstTypePtr{from.AsCTypePtr(), to.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewPackedVector2Array()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPolygonPathFinder.fnGetIntersections), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *PolygonPathFinder) GetClosestPoint(point Vector2, ) Vector2 {
  cargs := []gdc.ConstTypePtr{point.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewVector2()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPolygonPathFinder.fnGetClosestPoint), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *PolygonPathFinder) IsPointInside(point Vector2, ) bool {
  cargs := []gdc.ConstTypePtr{point.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewBool()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPolygonPathFinder.fnIsPointInside), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *PolygonPathFinder) SetPointPenalty(idx int64, penalty float64, )  {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&idx) , gdc.ConstTypePtr(&penalty) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPolygonPathFinder.fnSetPointPenalty), me.obj, unsafe.SliceData(cargs), nil)

}

func  (me *PolygonPathFinder) GetPointPenalty(idx int64, ) float64 {
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&idx) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewFloat()
  pinner.Pin(&idx)

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPolygonPathFinder.fnGetPointPenalty), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return ret.Get()
}

func  (me *PolygonPathFinder) GetBounds() Rect2 {
  cargs := []gdc.ConstTypePtr{}
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewRect2()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPolygonPathFinder.fnGetBounds), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
