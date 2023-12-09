// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
// TODO: disgusting imports







  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type GodotPhysicsServer2D struct {
  obj gdc.ObjectPtr
}

func (me *GodotPhysicsServer2D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *GodotPhysicsServer2D) BaseClass() string {
  return "GodotPhysicsServer2D"
}



// Enums

func (me *GodotPhysicsServer2D) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *GodotPhysicsServer2D) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *GodotPhysicsServer2D) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}


// Methods

// TODO: properties (class)

// TODO: signals (class)
