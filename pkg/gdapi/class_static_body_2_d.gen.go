// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type StaticBody2D struct {
  obj gdc.ObjectPtr
}

func (me *StaticBody2D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *StaticBody2D) BaseClass() string {
  return "StaticBody2D"
}



// Enums

func (me *StaticBody2D) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *StaticBody2D) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *StaticBody2D) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

func  (me *StaticBody2D) SetConstantLinearVelocity(vel Vector2, )  {
  classNameV := StringNameFromStr("StaticBody2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_constant_linear_velocity")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 743155724) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(vel.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *StaticBody2D) SetConstantAngularVelocity(vel float32, )  {
  classNameV := StringNameFromStr("StaticBody2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_constant_angular_velocity")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 373806689) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(&vel), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *StaticBody2D) GetConstantLinearVelocity() Vector2 {
  classNameV := StringNameFromStr("StaticBody2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_constant_linear_velocity")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 3341600327) // FIXME: should cache?
  var ret Vector2
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *StaticBody2D) GetConstantAngularVelocity() float32 {
  classNameV := StringNameFromStr("StaticBody2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("get_constant_angular_velocity")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1740695150) // FIXME: should cache?
  var ret float32
  cargs := []gdc.ConstTypePtr{}
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), gdc.TypePtr(&ret))
  return ret
}

func  (me *StaticBody2D) SetPhysicsMaterialOverride(physics_material_override PhysicsMaterial, )  {
  classNameV := StringNameFromStr("StaticBody2D")
  defer classNameV.Destroy()
  methodNameV := StringNameFromStr("set_physics_material_override")
  defer methodNameV.Destroy()
  methodPtr := giface.ClassdbGetMethodBind(classNameV.AsCPtr(), methodNameV.AsCPtr(), 1784508650) // FIXME: should cache?
  cargs := []gdc.ConstTypePtr{gdc.ConstTypePtr(physics_material_override.AsCTypePtr()), }
  giface.ObjectMethodBindPtrcall(methodPtr, me.obj, unsafe.SliceData(cargs), nil)
}

func  (me *StaticBody2D) GetPhysicsMaterialOverride() PhysicsMaterial {
  classNameV := StringNameFromStr("StaticBody2D")
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

func (me *StaticBody2D) GetPropPhysicsMaterialOverride() PhysicsMaterial {
  panic("TODO: implement")
}

func (me *StaticBody2D) SetPropPhysicsMaterialOverride(value PhysicsMaterial) {
  panic("TODO: implement")
}

func (me *StaticBody2D) GetPropConstantLinearVelocity() Vector2 {
  panic("TODO: implement")
}

func (me *StaticBody2D) SetPropConstantLinearVelocity(value Vector2) {
  panic("TODO: implement")
}

func (me *StaticBody2D) GetPropConstantAngularVelocity() float32 {
  panic("TODO: implement")
}

func (me *StaticBody2D) SetPropConstantAngularVelocity(value float32) {
  panic("TODO: implement")
}