// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"
  "runtime"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

// FIXME: avoid unused import warning
var _ unsafe.Pointer
var _ runtime.Pinner

type PhysicsDirectBodyState3DExtension struct {
  PhysicsDirectBodyState3D
}

func (me *PhysicsDirectBodyState3DExtension) BaseClass() string {
  return "PhysicsDirectBodyState3DExtension"
}

func NewPhysicsDirectBodyState3DExtension() *PhysicsDirectBodyState3DExtension {
  str := StringNameFromStr("PhysicsDirectBodyState3DExtension") // FIXME: should cache?
  defer str.Destroy()

	objPtr := giface.ClassdbConstructObject(str.AsCPtr())
  obj := &PhysicsDirectBodyState3DExtension{}
  obj.SetBaseObject(objPtr)
  return obj
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

// Signals
