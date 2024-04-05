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

type ptrsForPhysicsDirectSpaceState3DList struct {
  fnIntersectPoint gdc.MethodBindPtr
  fnIntersectRay gdc.MethodBindPtr
  fnIntersectShape gdc.MethodBindPtr
  fnCastMotion gdc.MethodBindPtr
  fnCollideShape gdc.MethodBindPtr
  fnGetRestInfo gdc.MethodBindPtr
}

var ptrsForPhysicsDirectSpaceState3D ptrsForPhysicsDirectSpaceState3DList

func initPhysicsDirectSpaceState3DPtrs(iface gdc.Interface) {

  className := StringNameFromStr("PhysicsDirectSpaceState3D")
  defer className.Destroy()
  {
    methodName := StringNameFromStr("intersect_point")
    defer methodName.Destroy()
    ptrsForPhysicsDirectSpaceState3D.fnIntersectPoint = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 975173756))
  }
  {
    methodName := StringNameFromStr("intersect_ray")
    defer methodName.Destroy()
    ptrsForPhysicsDirectSpaceState3D.fnIntersectRay = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3957970750))
  }
  {
    methodName := StringNameFromStr("intersect_shape")
    defer methodName.Destroy()
    ptrsForPhysicsDirectSpaceState3D.fnIntersectShape = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3762137681))
  }
  {
    methodName := StringNameFromStr("cast_motion")
    defer methodName.Destroy()
    ptrsForPhysicsDirectSpaceState3D.fnCastMotion = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1778757334))
  }
  {
    methodName := StringNameFromStr("collide_shape")
    defer methodName.Destroy()
    ptrsForPhysicsDirectSpaceState3D.fnCollideShape = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 3762137681))
  }
  {
    methodName := StringNameFromStr("get_rest_info")
    defer methodName.Destroy()
    ptrsForPhysicsDirectSpaceState3D.fnGetRestInfo = ensurePtr(iface.ClassdbGetMethodBind(className.AsCPtr(), methodName.AsCPtr(), 1376751592))
  }
}

type PhysicsDirectSpaceState3D struct {
  Object
}

func (me *PhysicsDirectSpaceState3D) BaseClass() string {
  return "PhysicsDirectSpaceState3D"
}

func NewPhysicsDirectSpaceState3D() *PhysicsDirectSpaceState3D {
  str := StringNameFromStr("PhysicsDirectSpaceState3D") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &PhysicsDirectSpaceState3D{}
  obj.SetBaseObject(objPtr)
  return obj
}



// Enums

func (me *PhysicsDirectSpaceState3D) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *PhysicsDirectSpaceState3D) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *PhysicsDirectSpaceState3D) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *PhysicsDirectSpaceState3D) IntersectPoint(parameters PhysicsPointQueryParameters3D, max_results int64, ) []Dictionary {
  cargs := []gdc.ConstTypePtr{parameters.AsCTypePtr(), gdc.ConstTypePtr(&max_results) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewArray()
  defer ret.Destroy()
  pinner.Pin(&max_results)

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsDirectSpaceState3D.fnIntersectPoint), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  sliceRet, err := ConvertArrayToSlice[Dictionary](ret)
  if err != nil {
    log.Printf("Error converting return value to slice: %v", err) // FIXME: bad logging
    return nil
  }
return sliceRet
}

func  (me *PhysicsDirectSpaceState3D) IntersectRay(parameters PhysicsRayQueryParameters3D, ) Dictionary {
  cargs := []gdc.ConstTypePtr{parameters.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewDictionary()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsDirectSpaceState3D.fnIntersectRay), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *PhysicsDirectSpaceState3D) IntersectShape(parameters PhysicsShapeQueryParameters3D, max_results int64, ) []Dictionary {
  cargs := []gdc.ConstTypePtr{parameters.AsCTypePtr(), gdc.ConstTypePtr(&max_results) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewArray()
  defer ret.Destroy()
  pinner.Pin(&max_results)

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsDirectSpaceState3D.fnIntersectShape), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  sliceRet, err := ConvertArrayToSlice[Dictionary](ret)
  if err != nil {
    log.Printf("Error converting return value to slice: %v", err) // FIXME: bad logging
    return nil
  }
return sliceRet
}

func  (me *PhysicsDirectSpaceState3D) CastMotion(parameters PhysicsShapeQueryParameters3D, ) PackedFloat32Array {
  cargs := []gdc.ConstTypePtr{parameters.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewPackedFloat32Array()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsDirectSpaceState3D.fnCastMotion), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *PhysicsDirectSpaceState3D) CollideShape(parameters PhysicsShapeQueryParameters3D, max_results int64, ) []Vector3 {
  cargs := []gdc.ConstTypePtr{parameters.AsCTypePtr(), gdc.ConstTypePtr(&max_results) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewArray()
  defer ret.Destroy()
  pinner.Pin(&max_results)

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsDirectSpaceState3D.fnCollideShape), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  sliceRet, err := ConvertArrayToSlice[Vector3](ret)
  if err != nil {
    log.Printf("Error converting return value to slice: %v", err) // FIXME: bad logging
    return nil
  }
return sliceRet
}

func  (me *PhysicsDirectSpaceState3D) GetRestInfo(parameters PhysicsShapeQueryParameters3D, ) Dictionary {
  cargs := []gdc.ConstTypePtr{parameters.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewDictionary()

  giface.ObjectMethodBindPtrcall(ensurePtr(ptrsForPhysicsDirectSpaceState3D.fnGetRestInfo), me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

// Signals
