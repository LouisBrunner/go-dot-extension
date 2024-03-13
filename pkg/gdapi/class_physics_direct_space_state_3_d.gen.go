// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type PhysicsDirectSpaceState3D struct {
  obj gdc.ObjectPtr
}

func (me *PhysicsDirectSpaceState3D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *PhysicsDirectSpaceState3D) BaseClass() string {
  return "PhysicsDirectSpaceState3D"
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

func  (me *PhysicsDirectSpaceState3D) IntersectPoint(parameters PhysicsPointQueryParameters3D, max_results int, ) Dictionary {
  classNameV := StringNameFromStr("PhysicsDirectSpaceState3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("intersect_point")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 975173756) // FIXME: should cache?
  var ret Dictionary
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(parameters.AsCTypePtr()), gdc.ConstTypePtr(&max_results), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *PhysicsDirectSpaceState3D) IntersectRay(parameters PhysicsRayQueryParameters3D, ) Dictionary {
  classNameV := StringNameFromStr("PhysicsDirectSpaceState3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("intersect_ray")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3957970750) // FIXME: should cache?
  var ret Dictionary
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(parameters.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *PhysicsDirectSpaceState3D) IntersectShape(parameters PhysicsShapeQueryParameters3D, max_results int, ) Dictionary {
  classNameV := StringNameFromStr("PhysicsDirectSpaceState3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("intersect_shape")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3762137681) // FIXME: should cache?
  var ret Dictionary
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(parameters.AsCTypePtr()), gdc.ConstTypePtr(&max_results), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *PhysicsDirectSpaceState3D) CastMotion(parameters PhysicsShapeQueryParameters3D, ) PackedFloat32Array {
  classNameV := StringNameFromStr("PhysicsDirectSpaceState3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("cast_motion")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1778757334) // FIXME: should cache?
  var ret PackedFloat32Array
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(parameters.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *PhysicsDirectSpaceState3D) CollideShape(parameters PhysicsShapeQueryParameters3D, max_results int, ) Vector3 {
  classNameV := StringNameFromStr("PhysicsDirectSpaceState3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("collide_shape")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3762137681) // FIXME: should cache?
  var ret Vector3
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(parameters.AsCTypePtr()), gdc.ConstTypePtr(&max_results), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *PhysicsDirectSpaceState3D) GetRestInfo(parameters PhysicsShapeQueryParameters3D, ) Dictionary {
  classNameV := StringNameFromStr("PhysicsDirectSpaceState3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_rest_info")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1376751592) // FIXME: should cache?
  var ret Dictionary
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(parameters.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

// Signals
