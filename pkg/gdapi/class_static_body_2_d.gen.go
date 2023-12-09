// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

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
  panic("TODO: implement")
}

func  (me *StaticBody2D) SetConstantAngularVelocity(vel float32, )  {
  panic("TODO: implement")
}

func  (me *StaticBody2D) GetConstantLinearVelocity()  {
  panic("TODO: implement")
}

func  (me *StaticBody2D) GetConstantAngularVelocity()  {
  panic("TODO: implement")
}

func  (me *StaticBody2D) SetPhysicsMaterialOverride(physics_material_override PhysicsMaterial, )  {
  panic("TODO: implement")
}

func  (me *StaticBody2D) GetPhysicsMaterialOverride()  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
