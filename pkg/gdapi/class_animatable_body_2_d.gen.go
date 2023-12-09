// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type AnimatableBody2D struct {
  obj gdc.ObjectPtr
}

func (me *AnimatableBody2D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *AnimatableBody2D) BaseClass() string {
  return "AnimatableBody2D"
}



// Enums

func (me *AnimatableBody2D) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *AnimatableBody2D) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

func  (me *AnimatableBody2D) SetSyncToPhysics(enable bool, )  {
  panic("TODO: implement")
}

func  (me *AnimatableBody2D) IsSyncToPhysicsEnabled()  {
  panic("TODO: implement")
}

// TODO: properties (class)

// TODO: signals (class)
