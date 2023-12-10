// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type GodotPhysicsServer3D struct {
  obj gdc.ObjectPtr
}

func (me *GodotPhysicsServer3D) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *GodotPhysicsServer3D) BaseClass() string {
  return "GodotPhysicsServer3D"
}



// Enums

func (me *GodotPhysicsServer3D) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *GodotPhysicsServer3D) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *GodotPhysicsServer3D) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

// Properties