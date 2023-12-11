// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "unsafe"

  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

var _ unsafe.Pointer // FIXME: avoid unused import warning

type PhysicsServer3DRenderingServerHandler struct {
  obj gdc.ObjectPtr
}

func (me *PhysicsServer3DRenderingServerHandler) SetBaseObject(obj gdc.ObjectPtr) {
  me.obj = obj
}

func (me *PhysicsServer3DRenderingServerHandler) BaseClass() string {
  return "PhysicsServer3DRenderingServerHandler"
}



// Enums

func (me *PhysicsServer3DRenderingServerHandler) Type() gdc.VariantType {
  return gdc.VariantTypeObject
}

func (me *PhysicsServer3DRenderingServerHandler) AsTypePtr() gdc.TypePtr {
  return gdc.TypePtr(me.obj)
}

func (me *PhysicsServer3DRenderingServerHandler) AsCTypePtr() gdc.ConstTypePtr {
  return gdc.ConstTypePtr(me.obj)
}

// Methods

// Signals
