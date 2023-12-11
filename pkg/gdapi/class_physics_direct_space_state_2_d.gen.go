// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type PhysicsDirectSpaceState2D struct {
  obj gdc.ObjectPtr
}

func (me *PhysicsDirectSpaceState2D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *PhysicsDirectSpaceState2D) BaseClass() string {
  return "PhysicsDirectSpaceState2D"
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

func  (me *PhysicsDirectSpaceState2D) IntersectPoint(parameters PhysicsPointQueryParameters2D, max_results int, ) Dictionary {
  classNameV := StringNameFromStr("PhysicsDirectSpaceState2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("intersect_point")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3278207904) // FIXME: should cache?
  var ret Dictionary
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(parameters.AsCTypePtr()), gdc.ConstTypePtr(&max_results), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *PhysicsDirectSpaceState2D) IntersectRay(parameters PhysicsRayQueryParameters2D, ) Dictionary {
  classNameV := StringNameFromStr("PhysicsDirectSpaceState2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("intersect_ray")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1590275562) // FIXME: should cache?
  var ret Dictionary
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(parameters.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *PhysicsDirectSpaceState2D) IntersectShape(parameters PhysicsShapeQueryParameters2D, max_results int, ) Dictionary {
  classNameV := StringNameFromStr("PhysicsDirectSpaceState2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("intersect_shape")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3803848594) // FIXME: should cache?
  var ret Dictionary
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(parameters.AsCTypePtr()), gdc.ConstTypePtr(&max_results), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *PhysicsDirectSpaceState2D) CastMotion(parameters PhysicsShapeQueryParameters2D, ) PackedFloat32Array {
  classNameV := StringNameFromStr("PhysicsDirectSpaceState2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("cast_motion")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 711275086) // FIXME: should cache?
  var ret PackedFloat32Array
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(parameters.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *PhysicsDirectSpaceState2D) CollideShape(parameters PhysicsShapeQueryParameters2D, max_results int, ) Vector2 {
  classNameV := StringNameFromStr("PhysicsDirectSpaceState2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("collide_shape")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3803848594) // FIXME: should cache?
  var ret Vector2
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(parameters.AsCTypePtr()), gdc.ConstTypePtr(&max_results), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *PhysicsDirectSpaceState2D) GetRestInfo(parameters PhysicsShapeQueryParameters2D, ) Dictionary {
  classNameV := StringNameFromStr("PhysicsDirectSpaceState2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_rest_info")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2803666496) // FIXME: should cache?
  var ret Dictionary
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(parameters.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

// Signals
