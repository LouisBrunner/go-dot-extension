// Code generated automatically by `go-dot-extension/pkg/gdapi/apigenerator`. DO NOT EDIT.
package gdapi

import (
  "github.com/LouisBrunner/go-dot-extension/pkg/gdc"
)

type PhysicsShapeQueryParameters3D struct {
  obj gdc.ObjectPtr
}

func createPhysicsShapeQueryParameters3D(obj gdc.ObjectPtr) *PhysicsShapeQueryParameters3D {
  return &PhysicsShapeQueryParameters3D{
    obj: obj,
  }
}

func (me *PhysicsShapeQueryParameters3D) BaseClass() string {
  return "PhysicsShapeQueryParameters3D"
}
