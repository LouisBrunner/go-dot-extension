// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type AnimatableBody3D struct {
  obj gdc.ObjectPtr
}

func (me *AnimatableBody3D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *AnimatableBody3D) BaseClass() string {
  return "AnimatableBody3D"
}



// Enums

func (me *AnimatableBody3D) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *AnimatableBody3D) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *AnimatableBody3D) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *AnimatableBody3D) SetSyncToPhysics(enable bool, )  {
  panic("TODO: implement")
}

func  (me *AnimatableBody3D) IsSyncToPhysicsEnabled()  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
