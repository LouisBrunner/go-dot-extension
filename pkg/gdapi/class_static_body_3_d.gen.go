// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type StaticBody3D struct {
  obj gdc.ObjectPtr
}

func (me *StaticBody3D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *StaticBody3D) BaseClass() string {
  return "StaticBody3D"
}



// Enums

func (me *StaticBody3D) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *StaticBody3D) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *StaticBody3D) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *StaticBody3D) SetConstantLinearVelocity(vel Vector3, )  {
  classNameV := StringNameFromStr("StaticBody3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_constant_linear_velocity")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3460891852) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(vel.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *StaticBody3D) SetConstantAngularVelocity(vel Vector3, )  {
  classNameV := StringNameFromStr("StaticBody3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_constant_angular_velocity")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3460891852) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(vel.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *StaticBody3D) GetConstantLinearVelocity() Vector3 {
  classNameV := StringNameFromStr("StaticBody3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_constant_linear_velocity")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3360562783) // FIXME: should cache?
  var ret Vector3
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *StaticBody3D) GetConstantAngularVelocity() Vector3 {
  classNameV := StringNameFromStr("StaticBody3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_constant_angular_velocity")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3360562783) // FIXME: should cache?
  var ret Vector3
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *StaticBody3D) SetPhysicsMaterialOverride(physics_material_override PhysicsMaterial, )  {
  classNameV := StringNameFromStr("StaticBody3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_physics_material_override")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1784508650) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(physics_material_override.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *StaticBody3D) GetPhysicsMaterialOverride() PhysicsMaterial {
  classNameV := StringNameFromStr("StaticBody3D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_physics_material_override")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 2521850424) // FIXME: should cache?
  var ret PhysicsMaterial
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}
// Properties
// FIXME: can't seem to be able to use those from this side of the API

// Signals
