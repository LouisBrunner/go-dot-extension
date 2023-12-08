// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type PhysicsServer3D struct {
  obj gdc.ObjectPtr
}

func createPhysicsServer3D(obj gdc.ObjectPtr) *PhysicsServer3D {
  return &PhysicsServer3D{
    obj: obj,
  }
}

func (me *PhysicsServer3D) BaseClass() string {
  return "PhysicsServer3D"
}
