// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

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

func (me *StaticBody3D) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *StaticBody3D) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *StaticBody3D) SetConstantLinearVelocity(vel Vector3, )  {
  panic("TODO: implement")
}

func  (me *StaticBody3D) SetConstantAngularVelocity(vel Vector3, )  {
  panic("TODO: implement")
}

func  (me *StaticBody3D) GetConstantLinearVelocity()  {
  panic("TODO: implement")
}

func  (me *StaticBody3D) GetConstantAngularVelocity()  {
  panic("TODO: implement")
}

func  (me *StaticBody3D) SetPhysicsMaterialOverride(physics_material_override PhysicsMaterial, )  {
  panic("TODO: implement")
}

func  (me *StaticBody3D) GetPhysicsMaterialOverride()  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
