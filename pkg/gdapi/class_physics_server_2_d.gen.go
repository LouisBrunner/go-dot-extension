// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type PhysicsServer2D struct {
  obj gdc.ObjectPtr
}

func createPhysicsServer2D(obj gdc.ObjectPtr) *PhysicsServer2D {
  return &PhysicsServer2D{
    obj: obj,
  }
}

func (me *PhysicsServer2D) BaseClass() string {
  return "PhysicsServer2D"
}
