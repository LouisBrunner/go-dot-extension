// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type PhysicsServer3DExtension struct {
  obj gdc.ObjectPtr
}

func createPhysicsServer3DExtension(obj gdc.ObjectPtr) *PhysicsServer3DExtension {
  return &PhysicsServer3DExtension{
    obj: obj,
  }
}

func (me *PhysicsServer3DExtension) BaseClass() string {
  return "PhysicsServer3DExtension"
}
