// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type PhysicsDirectBodyState3DExtension struct {
  obj gdc.ObjectPtr
}

func (me *PhysicsDirectBodyState3DExtension) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *PhysicsDirectBodyState3DExtension) BaseClass() string {
  return "PhysicsDirectBodyState3DExtension"
}



// Enums

func (me *PhysicsDirectBodyState3DExtension) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *PhysicsDirectBodyState3DExtension) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *PhysicsDirectBodyState3DExtension) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

// Properties