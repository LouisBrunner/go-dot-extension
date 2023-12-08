// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type PhysicsServer3DRenderingServerHandler struct {
  obj gdc.ObjectPtr
}

func createPhysicsServer3DRenderingServerHandler(obj gdc.ObjectPtr) *PhysicsServer3DRenderingServerHandler {
  return &PhysicsServer3DRenderingServerHandler{
    obj: obj,
  }
}

func (me *PhysicsServer3DRenderingServerHandler) BaseClass() string {
  return "PhysicsServer3DRenderingServerHandler"
}
