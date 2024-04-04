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

type PhysicsDirectSpaceState2D struct {
  Object
}

func (me *PhysicsDirectSpaceState2D) BaseClass() string {
  return "PhysicsDirectSpaceState2D"
}

func NewPhysicsDirectSpaceState2D() *PhysicsDirectSpaceState2D {
  str := StringNameFromStr("PhysicsDirectSpaceState2D") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &PhysicsDirectSpaceState2D{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Enums

func (me *PhysicsDirectSpaceState2D) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *PhysicsDirectSpaceState2D) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *PhysicsDirectSpaceState2D) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *PhysicsDirectSpaceState2D) IntersectPoint(parameters PhysicsPointQueryParameters2D, max_results int64, ) []Dictionary {
  classNameV := StringNameFromStr("PhysicsDirectSpaceState2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("intersect_point")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2118456068) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{parameters.AsCTypePtr(), gdc.ConstTypePtr(&max_results) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewArray()
  defer ret.Destroy()
  pinner.Pin(&max_results)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  sliceRet, err := ConvertArrayToSlice[Dictionary](ret)
  if err != nil {
    log.Printf("Error converting return value to slice: %v", err) // FIXME: bad logging
    return nil
  }
return sliceRet
}

func  (me *PhysicsDirectSpaceState2D) IntersectRay(parameters PhysicsRayQueryParameters2D, ) Dictionary {
  classNameV := StringNameFromStr("PhysicsDirectSpaceState2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("intersect_ray")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1590275562) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{parameters.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewDictionary()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *PhysicsDirectSpaceState2D) IntersectShape(parameters PhysicsShapeQueryParameters2D, max_results int64, ) []Dictionary {
  classNameV := StringNameFromStr("PhysicsDirectSpaceState2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("intersect_shape")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2488867228) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{parameters.AsCTypePtr(), gdc.ConstTypePtr(&max_results) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewArray()
  defer ret.Destroy()
  pinner.Pin(&max_results)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  sliceRet, err := ConvertArrayToSlice[Dictionary](ret)
  if err != nil {
    log.Printf("Error converting return value to slice: %v", err) // FIXME: bad logging
    return nil
  }
return sliceRet
}

func  (me *PhysicsDirectSpaceState2D) CastMotion(parameters PhysicsShapeQueryParameters2D, ) PackedFloat32Array {
  classNameV := StringNameFromStr("PhysicsDirectSpaceState2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("cast_motion")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 711275086) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{parameters.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewPackedFloat32Array()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *PhysicsDirectSpaceState2D) CollideShape(parameters PhysicsShapeQueryParameters2D, max_results int64, ) []Vector2 {
  classNameV := StringNameFromStr("PhysicsDirectSpaceState2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("collide_shape")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2488867228) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{parameters.AsCTypePtr(), gdc.ConstTypePtr(&max_results) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewArray()
  defer ret.Destroy()
  pinner.Pin(&max_results)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  sliceRet, err := ConvertArrayToSlice[Vector2](ret)
  if err != nil {
    log.Printf("Error converting return value to slice: %v", err) // FIXME: bad logging
    return nil
  }
return sliceRet
}

func  (me *PhysicsDirectSpaceState2D) GetRestInfo(parameters PhysicsShapeQueryParameters2D, ) Dictionary {
  classNameV := StringNameFromStr("PhysicsDirectSpaceState2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_rest_info")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2803666496) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{parameters.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewDictionary()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

// Signals
