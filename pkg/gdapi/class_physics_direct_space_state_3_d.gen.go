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
  classNameV := StringNameFromStr("PhysicsDirectSpaceState3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("intersect_point")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 975173756) // FIXME: should cache?
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

func  (me *PhysicsDirectSpaceState3D) IntersectRay(parameters PhysicsRayQueryParameters3D, ) Dictionary {
  classNameV := StringNameFromStr("PhysicsDirectSpaceState3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("intersect_ray")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3957970750) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{parameters.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewDictionary()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *PhysicsDirectSpaceState3D) IntersectShape(parameters PhysicsShapeQueryParameters3D, max_results int64, ) []Dictionary {
  classNameV := StringNameFromStr("PhysicsDirectSpaceState3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("intersect_shape")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3762137681) // FIXME: should cache?
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

func  (me *PhysicsDirectSpaceState3D) CastMotion(parameters PhysicsShapeQueryParameters3D, ) PackedFloat32Array {
  classNameV := StringNameFromStr("PhysicsDirectSpaceState3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("cast_motion")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1778757334) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{parameters.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewPackedFloat32Array()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

func  (me *PhysicsDirectSpaceState3D) CollideShape(parameters PhysicsShapeQueryParameters3D, max_results int64, ) []Vector3 {
  classNameV := StringNameFromStr("PhysicsDirectSpaceState3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("collide_shape")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3762137681) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{parameters.AsCTypePtr(), gdc.ConstTypePtr(&max_results) , }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewArray()
  defer ret.Destroy()
  pinner.Pin(&max_results)

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  sliceRet, err := ConvertArrayToSlice[Vector3](ret)
  if err != nil {
    log.Printf("Error converting return value to slice: %v", err) // FIXME: bad logging
    return nil
  }
return sliceRet
}

func  (me *PhysicsDirectSpaceState3D) GetRestInfo(parameters PhysicsShapeQueryParameters3D, ) Dictionary {
  classNameV := StringNameFromStr("PhysicsDirectSpaceState3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_rest_info")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1376751592) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{parameters.AsCTypePtr(), }
  pinner := runtime.Pinner{}
  defer pinner.Unpin()
  ret := NewDictionary()

  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), ret.AsTypePtr())
  return *ret
}

// Signals
